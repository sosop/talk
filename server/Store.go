package server

import (
	"net"
)

type Store interface {
	Keep(key string, conn net.Conn)
	Get(key string) net.Conn
}

type MemoryStore struct {
	Conns map[string]net.Conn
}

func NewMemoryStore() *MemoryStore {
	conns := make(map[string]net.Conn)
	return &MemoryStore{conns}
}

func (ms *MemoryStore) Keep(key string, conn net.Conn) {

}

func (ms *MemoryStore) Get(key string) net.Conn {
	return nil
}
