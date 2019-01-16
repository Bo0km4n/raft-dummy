package kvs

import (
	"errors"
	"sync"
)

type Storage interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, error)
}

type kvs struct {
	Bucket *sync.Map
}

func NewStorage() Storage {
	return &kvs{
		Bucket: &sync.Map{},
	}
}

func (k *kvs) Set(key string, value interface{}) {
	k.Bucket.Store(key, value)
}

func (k *kvs) Get(key string) (interface{}, error) {
	v, ok := k.Bucket.Load(key)
	if !ok {
		return nil, errors.New("Failed load value")
	}
	return v, nil
}
