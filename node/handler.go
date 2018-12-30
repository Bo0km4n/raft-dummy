package node

import (
	"context"
	"fmt"

	"github.com/Bo0km4n/raft-dummy/proto"
)

func (s *state) handleCandidate() {
	for range s.candidateChan {
		if s.broadcastVoteRPC() {
			s.setMode(LEADER)
			s.Info(s.mode, "change to LEADER")
			s.tickHeartBeat()
			// os.Exit(0)
		} else {
			s.setMode(FOLLOWER)
			s.Info(s.mode, "change to FOLLOWER")
			s.ResetElectionTimeout()
		}
	}
}

func (s *state) startHeartBeat() {
	for range s.timer.C {
		s.Warn(s.mode, fmt.Sprintf("timouted %d", s.electionTimeout))
		s.setMode(CANDIDATE)
		s.incrementTerm()
		s.candidateChan <- struct{}{}
	}
}

func (s *state) stringMode() string {
	switch s.mode {
	case LEADER:
		return "Leader"
	case CANDIDATE:
		return "Candidate"
	case FOLLOWER:
		return "Follower"
	default:
		return "nil"
	}
}

func (s *state) heartBeat() {
	s.ResetElectionTimeout()
	// s.Info(s.getMode(), "received heart beats")
}

func (s *state) broadcastVoteRPC() bool {
	voted := 1
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
		if res.VoteGranted {
			// s.Info(s.getMode(), fmt.Sprintf("voted from: %s", n.Addr))
			voted++
		}
		s.mu.Unlock()
	}
	return voted > len(s.nodes)/2
}

func (s *state) tickHeartBeat() {
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

func (s *state) isLeader() bool {
	return s.getMode() == LEADER
}
