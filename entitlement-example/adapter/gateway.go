package adapter

import (
	"entitlement-example/domain"
	"entitlement-example/infrastructure"
	"fmt"
)

// Ensure InfraGateway implements the domain interface
var _ domain.EntitlementRepository = (*InfraGateway)(nil)

// --- Sub-Adapters (Component Adapters) ---

// CacheRepo is an Adapter that wraps the raw MemoryStore (Infra).
type CacheRepo struct {
	store *infrastructure.MemoryStore
}

func NewCacheRepo(store *infrastructure.MemoryStore) *CacheRepo {
	return &CacheRepo{store: store}
}

func (c *CacheRepo) Get(group, user string) (bool, bool) {
	key := fmt.Sprintf("%s:%s", group, user)
	val, found := c.store.Load(key)
	if !found {
		return false, false
	}
	return val.(bool), true
}

func (c *CacheRepo) Set(group, user string, isMember bool) {
	key := fmt.Sprintf("%s:%s", group, user)
	c.store.Store(key, isMember)
}

// ADRepo is an Adapter that wraps the PseudoADClient (Infra).
type ADRepo struct {
	client *infrastructure.PseudoADClient
}

func NewADRepo(client *infrastructure.PseudoADClient) *ADRepo {
	return &ADRepo{client: client}
}

func (a *ADRepo) FetchMembership(group, user string) (bool, error) {
	return a.client.CheckGroupMembership(group, user)
}

// --- Gateway Adapter (The Orchestrator) ---

// InfraGateway implements the Domain Interface.
// It orchestrates the CacheRepo and ADRepo.
type InfraGateway struct {
	cache *CacheRepo
	ad    *ADRepo
}

func NewInfraGateway(cache *CacheRepo, ad *ADRepo) *InfraGateway {
	return &InfraGateway{
		cache: cache,
		ad:    ad,
	}
}

// IsMember implements domain.EntitlementRepository.
// Strategy: Cache Look-aside.
func (g *InfraGateway) IsMember(groupName, userName string) (bool, error) {
	// 1. Check Cache
	if isMember, found := g.cache.Get(groupName, userName); found {
		fmt.Println("  [Adapter:Gateway] Cache HIT")
		return isMember, nil
	}
	fmt.Println("  [Adapter:Gateway] Cache MISS")

	// 2. Fetch from AD
	isMember, err := g.ad.FetchMembership(groupName, userName)
	if err != nil {
		return false, err
	}

	// 3. Update Cache
	g.cache.Set(groupName, userName, isMember)

	return isMember, nil
}
