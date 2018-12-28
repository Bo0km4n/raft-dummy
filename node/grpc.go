package node

import (
	"context"
	"errors"
	"fmt"

	"github.com/Bo0km4n/raft-dummy/proto"
)

func (s *state) AppendEntriesRPC(c context.Context, in *proto.AppendEntries) (*proto.AppendEntriesResult, error) {
	if len(in.Entries) == 0 {
		switch s.getMode() {
		case CANDIDATE:
			if in.Term < s.getCurTerm() {
				s.candidateChan <- struct{}{}
				return &proto.AppendEntriesResult{Term: s.getCurTerm(), Success: false}, nil
			}
			s.setMode(FOLLOWER)
		case FOLLOWER:
			// update
			if in.Term < s.getCurTerm() {
				return &proto.AppendEntriesResult{Term: s.getCurTerm(), Success: false}, nil
			}
			s.setTerm(in.Term)
			s.heartBeat()
			return &proto.AppendEntriesResult{Term: s.getCurTerm(), Success: true}, nil
		}
	}
	return &proto.AppendEntriesResult{}, nil
}

func (s *state) RequestVoteRPC(c context.Context, in *proto.RequestVote) (*proto.RequestVoteResult, error) {
	if in.Term < s.getCurTerm() {
		return &proto.RequestVoteResult{
			Term:        s.getCurTerm(),
			VoteGranted: false,
		}, nil
	}
	if s.getVotedFor() == 0 || s.getVotedFor() == in.CandidateId {
		// update log
		s.setVotedFor(in.CandidateId)
		return &proto.RequestVoteResult{
			Term:        in.Term,
			VoteGranted: true,
		}, nil
	}

	s.Info(s.getMode(), fmt.Sprintf("voted=%d from %d", s.getVotedFor(), in.CandidateId))

	return &proto.RequestVoteResult{}, errors.New("Not matched vote")
}

func (s *state) heartBeat() {
	s.ResetElectionTimeout()
	s.Info(s.getMode(), "received heart beats")
}

func (s *state) broadcastVoteRPC() bool {
	completed := 1
	for _, n := range s.nodes {
		s.mu.Lock()
		client := proto.NewRaftClient(n.Conn)
		res, err := client.RequestVoteRPC(context.Background(), &proto.RequestVote{
			Term:         s.currentTerm,
			CandidateId:  s.candidateID,
			LastLogIndex: s.GetLastLogIndex(),
			LastLogTerm:  s.GetLastLogTerm(),
		})
		if err != nil {
			// s.Warn(s.mode, err.Error())
			continue
		}
		if res.Term > s.currentTerm {
			s.currentTerm = res.Term
			return false
		}
		// pp.Println(res, err)
		if res.VoteGranted {
			//s.Info(s.getMode(), fmt.Sprintf("voted from: %s", n.Addr))
			completed++
		}
		s.mu.Unlock()
	}
	return completed > len(s.nodes)/2
}

func (s *state) broadcastHeartBeat() {
	go func() {
		for {
			for _, n := range s.nodes {
				client := proto.NewRaftClient(n.Conn)
				res, _ := client.AppendEntriesRPC(context.Background(), &proto.AppendEntries{
					Term:    s.currentTerm,
					Entries: []*proto.Entry{},
				})
				if !res.Success && res.Term > s.getCurTerm() {
					s.setMode(FOLLOWER)
					s.ResetElectionTimeout()
					s.Info(s.getMode(), "LEADER to be FOLLOWER")
					return
				}
			}
		}
	}()
}
