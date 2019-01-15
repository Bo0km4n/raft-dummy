package raft

import "os"

type Log struct {
	File *os.File
}

func NewLog(path string) (*Log, error) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return &Log{File: f}, nil
}
