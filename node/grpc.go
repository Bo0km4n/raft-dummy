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
	if s.getMode() == CANDIDATE {
		return &proto.RequestVoteResult{
			Term:        in.Term,
			VoteGranted: false,
		}, nil
	}
	if s.getVotedFor() == 0 || s.getVotedFor() == in.CandidateId {
		// update log
		s.setVotedFor(in.CandidateId)
		s.Info(s.getMode(), fmt.Sprintf("accpet vote from %d", s.getVotedFor()))

		return &proto.RequestVoteResult{
			Term:        in.Term,
			VoteGranted: true,
		}, nil
	}

	s.Info(s.getMode(), fmt.Sprintf("voted=%d from %d", s.getVotedFor(), in.CandidateId))

	return &proto.RequestVoteResult{
		VoteGranted: false,
	}, errors.New("Not matched vote")
}
