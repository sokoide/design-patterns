package infrastructure

import (
	"fmt"
	"sync"
	"time"
)

// --- 1. Pseudo AD Client (Infrastructure) ---

// PseudoADClient simulates a connection to Active Directory or LDAP.
type PseudoADClient struct {
	// Simulating network latency or connection state
	host string
}

func NewPseudoADClient(host string) *PseudoADClient {
	return &PseudoADClient{host: host}
}

// CheckGroupMembership simulates the network call to AD.
func (c *PseudoADClient) CheckGroupMembership(group, user string) (bool, error) {
	fmt.Printf("  [Infra:AD] Connecting to %s checking if '%s' is in '%s'...\n", c.host, user, group)
	time.Sleep(200 * time.Millisecond) // Simulate latency

	// Pseudo logic: "alice" is in "admin", everyone is in "users"
	if group == "users" {
		return true, nil
	}
	if group == "admin" && user == "alice" {
		return true, nil
	}
	return false, nil
}

// --- 2. Memory Store (Infrastructure) ---

// MemoryStore is a wrapper around sync.Map to act as a raw in-memory KV store.
type MemoryStore struct {
	data sync.Map
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

func (m *MemoryStore) Load(key string) (interface{}, bool) {
	return m.data.Load(key)
}

func (m *MemoryStore) Store(key string, value interface{}) {
	m.data.Store(key, value)
}
