package kvs

import "sync"

type Storage interface {
	Set(key string, value interface{}) (interface{}, error)
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

func (k *kvs) Set(key string, value interface{}) (interface{}, error) {
	return nil, nil
}

func (k *kvs) Get(key string) (interface{}, error) {
	return nil, nil
}
