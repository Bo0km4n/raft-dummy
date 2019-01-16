package raft

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
			s.validateAppendEntry(in)
			return &proto.AppendEntriesResult{Term: s.getCurTerm(), Success: true}, nil
		}
	}

	if s.isFollower() && len(in.Entries) != 0 {
		s.Info("receive log replicate")
		if err := s.appendLog(in); err != nil {
			return &proto.AppendEntriesResult{Success: false}, err
		}
	}

	return &proto.AppendEntriesResult{Success: true}, nil
}

func (s *state) RequestVoteRPC(c context.Context, in *proto.RequestVote) (*proto.RequestVoteResult, error) {
	if in.Term < s.getCurTerm() {
		return &proto.RequestVoteResult{
			Term:        s.getCurTerm(),
			VoteGranted: false,
		}, nil
	}
	if s.getMode() == CANDIDATE {
		return &proto.RequestVoteResult{
			Term:        in.Term,
			VoteGranted: false,
		}, nil
	}
	if s.getVotedFor() == 0 || s.getVotedFor() == in.CandidateId {
		// update log
		s.setVotedFor(in.CandidateId)
		s.Info(fmt.Sprintf("accpet vote from %d", s.getVotedFor()))

		return &proto.RequestVoteResult{
			Term:        in.Term,
			VoteGranted: true,
		}, nil
	}

	s.Info(fmt.Sprintf("voted=%d from %d", s.getVotedFor(), in.CandidateId))

	return &proto.RequestVoteResult{
		VoteGranted: false,
	}, errors.New("Not matched vote")
}

func (s *state) LogCommitRequestRPC(c context.Context, in *proto.LogCommitRequest) (*proto.LogCommitResponse, error) {
	if !s.isLeader() {
		return &proto.LogCommitResponse{Success: false}, nil
	}
	s.Info("receive new log")

	req := &proto.AppendEntries{
		Term:         s.getCurTerm(),
		LeaderId:     s.getMachineID(),
		PrevLogIndex: s.GetLastLogIndex(),
		PrevLogTerm:  s.GetLastLogTerm(),
		Entries:      []*proto.Entry{},
		LeaderCommit: s.getCommitIndex(),
	}

	for i := range in.Requests {
		req.Entries = append(req.Entries, &proto.Entry{
			Term:  req.Term,
			Index: req.PrevLogIndex + int64(i) + 1,
			Data:  in.Requests[i],
		})
	}

	if err := s.appendLog(req); err != nil {
		return &proto.LogCommitResponse{Success: false}, err
	}

	if err := s.replicateLog(req); err != nil {
		return &proto.LogCommitResponse{Success: false}, err
	}

	return &proto.LogCommitResponse{Success: true}, nil
}

func (s *state) IsLeaderRPC(c context.Context, in *proto.IsReaderRequest) (*proto.IsReaderResponse, error) {
	if s.getMode() == LEADER {
		return &proto.IsReaderResponse{IsReader: true}, nil
	}
	return &proto.IsReaderResponse{IsReader: false}, nil

}
