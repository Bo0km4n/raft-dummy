package node

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var servers []State

func TestMain(m *testing.M) {
	servers = launchServers()
	code := m.Run()
	os.Exit(code)
}

func launchServers() []State {
	fmt.Println("start nodes")
	var nodes []State
	for i := 0; i < 5; i++ {
		logger, _ := NewLog(fmt.Sprintf("./testdata/log_%d.log", i+1))
		n := NewNode(int64(i+1), int64(i+1), logger)
		port := fmt.Sprintf(":5005%d", i+1)
		go func() {
			fmt.Printf("bind port %s\n", port)
			n.Start(port)
		}()
		nodes = append(nodes, n)
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

func TestCommitLog(t *testing.T) {
	time.Sleep(2 * time.Second)
	defer stopServers()
}
