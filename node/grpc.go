package node

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/Bo0km4n/raft-dummy/proto"
)

func (s *state) AppendEntriesRPC(c context.Context, in *proto.AppendEntries) (*proto.AppendEntriesResult, error) {
	if len(in.Entries) == 0 {
		s.heartBeat()
	}
	return nil, nil
}

func (s *state) RequestVoteRPC(c context.Context, in *proto.RequestVote) (*proto.RequestVoteResult, error) {
	defer s.mu.Unlock()
	s.mu.Lock()
	if s.currentTerm < in.Term {
		s.currentTerm = in.Term
		s.resetVoted()
	}
	if s.mode != FOLLOWER || s.votedFor != 0 {
		s.Info(s.mode, fmt.Sprintf("can't accept vote from %d, %v", in.CandidateId, s.votedFor))
		return &proto.RequestVoteResult{
			Term:        s.currentTerm,
			VoteGranted: false,
		}, nil
	}
	s.timer.Stop()

	s.Info(s.mode, fmt.Sprintf("received request vote from %d", in.CandidateId))
	s.votedFor = in.CandidateId
	return &proto.RequestVoteResult{
		Term:        in.Term,
		VoteGranted: true,
	}, nil
}

func (s *state) heartBeat() {
	s.mu.Lock()
	s.ResetElectionTimeout()
	s.mu.Unlock()
}

func (s *state) broadcastVoteRPC() bool {
	completed := 1
	for _, n := range s.nodes {
		conn, err := grpc.Dial(n.Addr, grpc.WithInsecure())
		if err != nil {
			continue
		}
		client := proto.NewRaftClient(conn)
		res, err := client.RequestVoteRPC(context.Background(), &proto.RequestVote{
			Term:         s.currentTerm,
			CandidateId:  s.candidateID,
			LastLogIndex: s.GetLastLogIndex(),
			LastLogTerm:  s.GetLastLogTerm(),
		})
		// pp.Println(res, err)
		if res.VoteGranted {
			completed++
		}
	}
	return completed > len(s.nodes)/2
}
