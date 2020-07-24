package cache

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("not found")

type Memory struct {
	m  map[string]string
	mu sync.Mutex
}

func NewInMemoryCache() *Memory {
	return &Memory{
		m: make(map[string]string),
	}
}

func (c *Memory) Set(key, value string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.m[key] = value

	return nil
}

func (c *Memory) Get(key string) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	val, ok := c.m[key]
	if !ok {
		return "", ErrNotFound
	}

	return val, nil
}
