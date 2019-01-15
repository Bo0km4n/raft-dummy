package node

import (
	"fmt"
	"math/rand"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/Bo0km4n/raft-dummy/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type State interface {
	Start(port string)
	AddNode(addr ...string) error
	LaunchSideHandlers()
}

const (
	LEADER = iota
	CANDIDATE
	FOLLOWER
)

var electionTimeoutMin = int64(150)
var electionTimeoutMax = int64(300)

type state struct {
	mode            int64
	followerChan    chan struct{}
	candidateChan   chan struct{}
	leaderChan      chan struct{}
	machineID       int64
	candidateID     int64
	currentTerm     int64
	votedFor        int64
	logs            []*proto.Entry
	commitIndex     int64
	lastApplied     int64
	electionTimeout time.Duration
	nodes           []*node
	mu              sync.Mutex
	timer           *time.Timer
	logger          *Log
}

type node struct {
	Addr string `json:"addr"`
	Conn *grpc.ClientConn
}

func (s *state) ResetElectionTimeout() {
	s.electionTimeout = time.Duration(rand.Int63n(electionTimeoutMax-electionTimeoutMin) + electionTimeoutMin)
	s.timer.Reset(s.electionTimeout * time.Millisecond)
}

func NewNode(machineID, candidateID int64, logger *Log) State {
	rand.Seed(time.Now().Unix())

	s := &state{
		mode:            FOLLOWER,
		machineID:       machineID,
		followerChan:    make(chan struct{}),
		candidateChan:   make(chan struct{}),
		leaderChan:      make(chan struct{}),
		candidateID:     candidateID,
		currentTerm:     0,
		votedFor:        0,
		commitIndex:     0,
		lastApplied:     0,
		electionTimeout: time.Duration(rand.Int63n(electionTimeoutMax-electionTimeoutMin) + electionTimeoutMin),
		mu:              sync.Mutex{},
		logger:          logger,
	}

	return s
}

func (s *state) Start(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}

	srv := grpc.NewServer()
	proto.RegisterRaftServer(srv, s)

	reflection.Register(srv)
	if err := srv.Serve(lis); err != nil {
		fmt.Println(err)
		return
	}
}

func (s *state) LaunchSideHandlers() {
	s.timer = time.NewTimer(s.electionTimeout * time.Millisecond)

	go s.handleCandidate()
	go s.startHeartBeat()
}

func (s *state) AddNode(addrs ...string) error {
	for _, addr := range addrs {
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			s.Warn(err.Error())
			return err
		}
		s.nodes = append(s.nodes, &node{Addr: addr, Conn: conn})
	}
	return nil
}

func (s *state) Info(msg string) {
	logrus.Infof("Machined-id: %d, Mode: %s, Term: %d, msg: %s", s.machineID, s.stringMode(), s.currentTerm, msg)
}

func (s *state) Warn(msg string) {
	logrus.Warnf("Machined-id: %d, Mode: %s, Term: %d, msg: %s", s.machineID, s.stringMode(), s.currentTerm, msg)
}

func (s *state) GetLastLogIndex() int64 {
	if len(s.logs) == 0 {
		return -1
	}
	return int64(len(s.logs) - 1)
}

func (s *state) GetLastLogTerm() int64 {
	if len(s.logs) == 0 {
		return 0
	}
	return int64(s.logs[len(s.logs)-1].Term)
}

func (s *state) resetVoted() {
	s.votedFor = 0
}

func (s *state) getMode() int64 {
	return atomic.LoadInt64(&s.mode)
}

func (s *state) getCurTerm() int64 {
	return atomic.LoadInt64(&s.currentTerm)
}

func (s *state) incrementTerm() int64 {
	atomic.AddInt64(&s.currentTerm, 1)
	return atomic.LoadInt64(&s.currentTerm)
}

func (s *state) setTerm(v int64) int64 {
	atomic.StoreInt64(&s.currentTerm, v)
	return atomic.LoadInt64(&s.currentTerm)
}

func (s *state) setMode(v int64) int64 {
	atomic.StoreInt64(&s.mode, v)
	return atomic.LoadInt64(&s.mode)
}

func (s *state) setVotedFor(v int64) int64 {
	atomic.StoreInt64(&s.votedFor, v)
	return atomic.LoadInt64(&s.votedFor)
}

func (s *state) getVotedFor() int64 {
	return atomic.LoadInt64(&s.votedFor)
}

func (s *state) getCommitIndex() int64 {
	return atomic.LoadInt64(&s.commitIndex)
}

func (s *state) setCommitIndex(idx int64) {
	atomic.StoreInt64(&s.commitIndex, idx)
}

func (s *state) toLeader() {
	s.Info("change to LEADER")
	s.setMode(LEADER)
	s.setVotedFor(0)
	s.tickHeartBeat()
}

func (s *state) toFollower() {
	s.Info("change to FOLLOWER")
	s.setMode(FOLLOWER)
	s.setVotedFor(0)
	s.ResetElectionTimeout()
}

func (s *state) isLeader() bool {
	return s.getMode() == LEADER
}

func (s *state) isFollower() bool {
	return s.getMode() == FOLLOWER
}

func (s *state) getMachineID() int64 {
	return s.machineID
}
