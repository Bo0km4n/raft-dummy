package cmd

import (
	"fmt"
	"time"

	"github.com/Bo0km4n/raft-dummy/raft"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newNodeCommand())
}

var (
	candidateID int64
	machineID   int64
)

func newNodeCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "sim",
		Short: "start a raft simulation",
		Run:   startNode,
	}

	cmd.Flags().Int64Var(&machineID, "machine-id", 1, "machine id")
	cmd.Flags().Int64Var(&candidateID, "candidate-id", 1, "candidate id")

	return &cmd
}

func startNode(cnd *cobra.Command, args []string) {
	fmt.Println("start nodes")
	var nodes []raft.State
	for i := 0; i < 5; i++ {
		logger, _ := raft.NewLog(fmt.Sprintf("./testdata/log_%d.log", i+1))
		n := raft.NewNode(int64(i+1), int64(i+1), logger)
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

	// time.Sleep(2 * time.Second)

	// for _, v := range nodes {
	// 	v.Stop()
	// }
	time.Sleep(1000 * time.Second)
}
