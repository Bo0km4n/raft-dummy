package kvs

import (
	"errors"
	"strings"
	"sync"
)

// Storage raft storage interface
type Storage interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, error)
	Del(key string)
	Put(key string, value interface{}) error
	Eval(statement string) (*Result, error)
}

type KVS struct {
	Bucket *sync.Map
}

type Result struct {
	Success bool
	Value   interface{}
}

// NewStorage generate Storage interface
func NewStorage() Storage {
	return &KVS{
		Bucket: &sync.Map{},
	}
}

func (k *KVS) Set(key string, value interface{}) {
	k.Bucket.Store(key, value)
}

func (k *KVS) Get(key string) (interface{}, error) {
	v, ok := k.Bucket.Load(key)
	if !ok {
		return nil, errors.New("Failed load value")
	}
	return v, nil
}

func (k *KVS) Del(key string) {
	k.Bucket.Delete(key)
}

func (k *KVS) Put(key string, value interface{}) error {
	if _, ok := k.Bucket.Load(key); !ok {
		return errors.New("Not found key")
	}
	k.Bucket.Store(key, value)
	return nil
}

func (k *KVS) Eval(statement string) (*Result, error) {
	tokens := strings.Split(statement, " ")
	switch tokens[0] {
	case "SET":
		k.Set(tokens[1], tokens[2])
		return &Result{
			Success: true,
			Value:   tokens[2],
		}, nil
	case "GET":
		res, err := k.Get(tokens[1])
		if err != nil {
			return nil, err
		}
		return &Result{
			Success: true,
			Value:   res,
		}, nil
	case "PUT":
		if err := k.Put(tokens[1], tokens[2]); err != nil {
			return nil, err
		}
		return &Result{
			Success: true,
			Value:   tokens[2],
		}, nil
	case "DEL":
		k.Del(tokens[1])
		return &Result{
			Success: true,
			Value:   nil,
		}, nil
	}
	return nil, nil
}
