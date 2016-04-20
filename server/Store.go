package server

import (
	"net"
)

type Store interface {
	Keep(key string, conn net.Conn)
	Get(key string) net.Conn
	Del(key string)
}

type MemoryStore struct {
	Conns map[string]net.Conn
}

func NewMemoryStore() *MemoryStore {
	conns := make(map[string]net.Conn)
	return &MemoryStore{conns}
}

func (ms *MemoryStore) Keep(key string, conn net.Conn) {
	ms.Conns[key] = conn
}

func (ms *MemoryStore) Get(key string) net.Conn {
	if c, ok := ms.Conns[key]; ok {
		return c
	}
	return nil
}

func (ms *MemoryStore) Del(key string) {
	if _, ok := ms.Conns[key]; ok {
		delete(ms.Conns, key)
	}
}
