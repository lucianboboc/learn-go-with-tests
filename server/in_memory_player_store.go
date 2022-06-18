package main

import "sync"

type InMemoryPlayerStore struct {
	store map[string]int
	mutex sync.RWMutex
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}, sync.RWMutex{}}
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	return i.store[name]
}
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	i.store[name]++
}
