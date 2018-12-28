package node

import "fmt"

func (s *state) handleCandidate() {
	for range s.candidateChan {
		if s.broadcastVoteRPC() {
			s.setMode(LEADER)
			s.Info(s.mode, "change to LEADER")
			s.broadcastHeartBeat()
			// os.Exit(0)
		} else {
			s.setMode(FOLLOWER)
			s.Info(s.mode, "change to FOLLOWER")
			s.ResetElectionTimeout()
			go s.handleCandidate()
		}
		return
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
