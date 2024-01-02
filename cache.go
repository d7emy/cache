package cache

import "sync"

type cache struct {
	db map[string]bool
	sync.Mutex
}

func NewCache() *cache {
	return &cache{db: map[string]bool{}}
}

func (m *cache) Set(t string) {
	m.Lock()
	m.db[t] = true
	m.Unlock()
}

func (m *cache) Check(t string) bool {
	m.Lock()
	defer m.Unlock()
	return m.db[t]
}

func (m *cache) SetCheck(t string) bool {
	if !m.Check(t) {
		m.Set(t)
		return false
	}
	return true
}

func (m *cache) Len() int {
	m.Lock()
	defer m.Unlock()
	return len(m.db)
}
