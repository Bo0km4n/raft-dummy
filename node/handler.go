package node

import "os"

func (s *state) handleMode() {
	go s.handleFollower()
	go s.handleCandidate()
	go s.handleLeader()
}

func (s *state) handleFollower() {
	for range s.followerChan {
		s.mu.Lock()
		s.Info(s.mode, "changed to Follower")
		s.ResetElectionTimeout()
		s.mode = FOLLOWER
		s.mu.Unlock()
	}
}

func (s *state) handleCandidate() {
	for range s.candidateChan {
		s.mu.Lock()
		s.mode = CANDIDATE
		s.Info(s.mode, "changed to Candidate")
		s.currentTerm++
		s.mu.Unlock()
		if !s.broadcastVoteRPC() {
			s.Info(s.mode, "Not accepted leader vote")
			s.followerChan <- struct{}{}
		} else {
			s.leaderChan <- struct{}{}
		}
	}
}

func (s *state) handleLeader() {
	for range s.leaderChan {
		s.mu.Lock()
		s.mode = LEADER
		s.Info(s.mode, "changed to Leader")
		s.mu.Unlock()
		os.Exit(0)
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
