package types

import (
	"sync"
)

/************************* Concurrent Map *************************/

type ConcurrentMap struct {
	elems map[Any]Any
	
	locker sync.RWMutex
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		elems: make(map[Any]Any),
	}
}

func (m *ConcurrentMap) Len() int {
	m.locker.RLock()
	defer m.locker.RUnlock()

	return len(m.elems)
}

func (m *ConcurrentMap) Contains(key Any) bool {
	m.locker.RLock()
	defer m.locker.RUnlock()

	_, exist := m.elems[key]
	return exist
}

func (m *ConcurrentMap) Get(key Any) (Any, bool) {
	m.locker.RLock()
	defer m.locker.RUnlock()

	return m.elems[key]
}

func (m *ConcurrentMap) Put(key, value Any) {
	m.locker.Lock()
	defer m.locker.Unlock()

	m.elems[key] = value
}

func (m *ConcurrentMap) Delete(key Any) {
	m.locker.Lock()
	defer m.locker.Unlock()

	delete(m.elems, key)
}

// Take care NOT to run Long-Time routine
func (m *ConcurrentMap) ForEach(visit func (Any, Any)) {
	m.locker.RLock()
	defer m.locker.RUnlock()

	for key, val := range m.elems {
		visit(key, val)
	}
}
