package raft

import (
	"context"
	"flag"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/Bo0km4n/raft-dummy/kvs"
	"github.com/Bo0km4n/raft-dummy/proto"
	"github.com/k0kubun/pp"
)

var servers []*state

var (
	cwd = flag.String("cwd", "", "set cwd")
)

func init() {
	flag.Parse()
	if *cwd != "" {
		if err := os.Chdir(*cwd); err != nil {
			panic(err)
		}
	}
}

func TestMain(m *testing.M) {
	servers = launchServers()
	code := m.Run()
	os.Exit(code)
}

func launchServers() []*state {
	fmt.Println("start nodes")
	var nodes []*state
	for i := 0; i < 5; i++ {
		path := fmt.Sprintf("%s/testdata/log_%d.log", *cwd, i+1)
		cleanLog(path)
		logger, err := NewLog(path)
		if err != nil {
			panic(err)
		}
		n := NewNode(int64(i+1), int64(i+1), logger)
		port := fmt.Sprintf(":5005%d", i+1)
		go func() {
			fmt.Printf("bind port %s\n", port)
			n.Start(port)
		}()
		nodes = append(nodes, n.(*state))
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Waiting launched nodes...")

	nodes[0].AddNode("127.0.0.1:50052", "127.0.0.1:50053", "127.0.0.1:50054", "127.0.0.1:50055")
	nodes[1].AddNode("127.0.0.1:50051", "127.0.0.1:50053", "127.0.0.1:50054", "127.0.0.1:50055")
	nodes[2].AddNode("127.0.0.1:50051", "127.0.0.1:50052", "127.0.0.1:50054", "127.0.0.1:50055")
	nodes[3].AddNode("127.0.0.1:50051", "127.0.0.1:50052", "127.0.0.1:50053", "127.0.0.1:50055")
	nodes[4].AddNode("127.0.0.1:50051", "127.0.0.1:50052", "127.0.0.1:50053", "127.0.0.1:50054")

	fmt.Println("Add 4 nodes to each nodes")

	for _, n := range nodes {
		n.LaunchSideHandlers()
	}

	return nodes
}

func stopServers() {
	for _, s := range servers {
		s.Stop()
	}
}

func cleanLog(path string) {
	os.Remove(path)
}

func TestCommitLog(t *testing.T) {
	defer stopServers()

	time.Sleep(2 * time.Second)
	var leader *state
	for _, s := range servers {
		res, _ := s.IsLeaderRPC(context.Background(), &proto.IsReaderRequest{})
		if res.IsReader {
			leader = s
		}
	}
	newLog := &proto.LogCommitRequest{
		Requests: [][]byte{
			[]byte("SET x 10"),
			[]byte("SET y 9"),
		},
	}
	res, err := leader.LogCommitRequestRPC(context.Background(), newLog)
	if err != nil {
		t.Fatal(err)
	}

	pp.Println(res)

	storage, _ := leader.Storage.(*kvs.KVS)
	x, _ := storage.Get("x")
	if x.(string) != "10" {
		t.Errorf("Not matched x")
	}
	y, _ := storage.Get("y")
	if y.(string) != "9" {
		t.Errorf("Not matched y")
	}
}
