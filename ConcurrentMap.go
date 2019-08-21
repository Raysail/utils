package utils

import (
	"sync"
)

type ConcurrentMap struct {
	lock *sync.RWMutex
	cmap map[interface{}]interface{}
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		lock: new(sync.RWMutex),
		cmap: make(map[interface{}]interface{}),
	}
}

func (m *ConcurrentMap) Get(k interface{}) (interface{}, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	value, ok := m.cmap[k]
	return value, ok
}

func (m *ConcurrentMap) Set(k interface{}, v interface{}) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	if val, ok := m.cmap[k]; !ok {
		m.cmap[k] = v
	} else if val != v {
		m.cmap[k] = v
	} else {
		return false
	}
	return true
}

func (m *ConcurrentMap) Contains(k interface{}) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if _, ok := m.cmap[k]; !ok {
		return false
	}
	return true
}

func (m *ConcurrentMap) Remove(k interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.cmap, k)
}

func (m *ConcurrentMap) Items() map[interface{}]interface{} {
	m.lock.RLock()
	defer m.lock.RUnlock()
	r := make(map[interface{}]interface{})
	for k, v := range m.cmap {
		r[k] = v
	}
	return r
}
