package raft

import (
	"context"
	"errors"
	"fmt"

	"github.com/Bo0km4n/raft-dummy/kvs"

	"github.com/Bo0km4n/raft-dummy/proto"
)

func (s *state) handleCandidate() {
	for range s.candidateChan {
		if s.broadcastVoteRPC() {
			s.toLeader()
		} else {
			s.toFollower()
		}
	}
}

func (s *state) startHeartBeat() {
	for range s.timer.C {
		s.Warn(fmt.Sprintf("timouted %d", s.electionTimeout))
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
	// s.Info( "received heart beats")
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
			// s.Info( fmt.Sprintf("voted from: %s", n.Addr))
			voted++
		}
		s.mu.Unlock()
	}
	return voted > len(s.nodes)/2
}

func (s *state) tickHeartBeat() {
	go func() {
		for {
			s.bcastHeartBeat()
		}
	}()
}

func (s *state) bcastHeartBeat() {
	for _, n := range s.nodes {
		client := proto.NewRaftClient(n.Conn)
		res, _ := client.AppendEntriesRPC(context.Background(), &proto.AppendEntries{
			Term:         s.currentTerm,
			LeaderCommit: s.getCommitIndex(),
			PrevLogIndex: s.GetLastLogIndex(),
			PrevLogTerm:  s.GetLastLogTerm(),
			Entries:      []*proto.Entry{},
		})
		if !res.Success && res.Term > s.getCurTerm() {
			s.setMode(FOLLOWER)
			s.ResetElectionTimeout()
			s.Info("LEADER to be FOLLOWER")
			return
		}
	}
}

func (s *state) replicateLog(e *proto.AppendEntries) ([]*proto.Result, error) {
	agreements := 0
	for _, n := range s.nodes {
		client := proto.NewRaftClient(n.Conn)
		res, err := client.AppendEntriesRPC(context.Background(), e)
		if err != nil {
			s.Warn(err.Error())
			continue
		}
		if !res.Success {
			s.Warn("Not agreeded request")
			continue
		}
		agreements++
	}
	if agreements > (len(s.nodes)+1)/2 {
		s.Info("Agreeded log")
		return s.commitLog(s.getCommitIndex()+1, s.getCommitIndex()+1+int64(len(e.Entries)))
	}
	return nil, errors.New("Not agreeded majority")
}

func (s *state) maybeLogReplication(e *proto.Entry) error {
	return nil
}

func (s *state) commitLog(start, end int64) ([]*proto.Result, error) {
	logs := s.logs[start:end]
	results := []*proto.Result{}
	for i := range logs {
		result := s.eval(logs[i])
		results = append(results, &proto.Result{
			Success: result.Success,
			Value:   result.Value.(string),
		})
	}
	s.setCommitIndex(end - 1)
	return results, nil
}

func (s *state) appendLog(e *proto.AppendEntries) error {
	for i := range e.Entries {
		s.logger.Append(fmt.Sprintf(`%d,%d,%s`, e.Entries[i].Index, e.Entries[i].Term, string(e.Entries[i].Data)))
		s.logs = append(s.logs, e.Entries[i])
	}
	return nil
}

func (s *state) eval(entry *proto.Entry) *kvs.Result {
	result, err := s.Storage.Eval(string(entry.Data))
	if err != nil {
		return nil
	}
	return result
}

func (s *state) validateAppendEntry(e *proto.AppendEntries) {
	if s.getCommitIndex() < e.LeaderCommit {
		// s.Info(fmt.Sprintf("callled validateAppendEntry and call commitLog getCommitIndex=%d, e.LeaderCommit=%d", s.getCommitIndex(), e.LeaderCommit))
		s.commitLog(s.getCommitIndex()+1, e.LeaderCommit+1)
	}
}
