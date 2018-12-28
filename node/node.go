package node

import (
	"fmt"
	"math/rand"
	"net"
	"sync"
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

type mode int64

const (
	LEADER = iota
	CANDIDATE
	FOLLOWER
)

var electionTimeoutMin = int64(150)
var electionTimeoutMax = int64(300)

type state struct {
	mode            mode
	followerChan    chan struct{}
	candidateChan   chan struct{}
	leaderChan      chan struct{}
	machineID       int64
	candidateID     int64
	currentTerm     int64
	votedFor        int64
	logs            []*log
	commitIndex     int64
	lastApplied     int64
	electionTimeout time.Duration
	nodes           []*node
	mu              sync.Mutex
	timer           *time.Timer
}

type node struct {
	Addr string `json:"addr"`
	Conn *grpc.ClientConn
}

type log struct {
	Term int64
	Data []byte
}

func (s *state) ResetElectionTimeout() {
	s.electionTimeout = time.Duration(rand.Int63n(electionTimeoutMax-electionTimeoutMin) + electionTimeoutMin)
	s.timer.Reset(s.electionTimeout * time.Millisecond)
}

func NewNode(machineID, candidateID int64) State {
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
	s.handleMode()
	go s.startHeartBeat()
}

func (s *state) AddNode(addrs ...string) error {
	for _, addr := range addrs {
		s.nodes = append(s.nodes, &node{Addr: addr})
	}
	return nil
}

func (s *state) startHeartBeat() {
	s.timer = time.NewTimer(s.electionTimeout * time.Millisecond)
	for range s.timer.C {
		s.candidateChan <- struct{}{}
	}
}

func (s *state) Info(mode mode, msg string) {
	logrus.Infof("Machined-id: %d, %s, %d, msg: %s", s.machineID, s.stringMode(), s.currentTerm, msg)
}

func (s *state) GetLastLogIndex() int64 {
	if len(s.logs) == 0 {
		return 0
	}
	return int64(len(s.logs) - 1)
}

func (s *state) GetLastLogTerm() int64 {
	if len(s.logs) == 0 {
		return 0
	}
	return int64(s.logs[len(s.logs)-1].Term)
}
