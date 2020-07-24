package cache

import "sync"

type Memory struct {
	m  map[string][]string
	mu sync.Mutex
}

func (c *Memory) Set(key, value string) error {
	return nil
}

func (c *Memory) Get(key string) (string, error) {
	return "", nil
}
