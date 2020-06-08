package types

import (
	"sync"
)

/************************* Basic Set *************************/

type Set struct {
	elems map[Any]Void
}

func NewSet() *Set {
	return &Set{
		elems: make(map[Any]Void),
	}
}

func (set *Set) Len() int {
	return len(set.elems)
}

func (set *Set) Contains(elem Any) bool {
	_, exist := set.elems[elem]
	return exist
}

func (set *Set) Insert(elem Any) {
	set.elems[elem] = Void{}
}

func (set *Set) Delete(elem Any) {
	delete(set.elems, elem)
}

func (set *Set) GetAll() []Any {
	elems := make([]Any, 0, len(set.elems))
	for elem, _ := range set.elems {
		elems = append(elems, elem)
	}
	return elems
}

func (set *Set) Clear() {
	set.elems = make(map[Any]Void)
}

func (set *Set) ForEach(visit func(Any)) {
	for elem, _ := range set.elems {
		visit(elem)
	}
}

func (set *Set) Filter(predict func(Any) bool) []Any {
	filtered := make([]Any, 0, len(set.elems))
	for elem, _ := range set.elems {
		if predict(elem) {
			filtered = append(filtered, elem)
		}
	}
	return filtered
}

func (set *Set) DeleteIf(predict func(Any) bool) {
	for elem, _ := range set.elems {
		if predict(elem) {
			delete(set.elems, elem)
		}
	}
}

/************************* Concurrent Set *************************/

type ConcurrentSet struct {
	elems map[Any]Void

	locker sync.RWMutex
}

func NewConcurrentSet() *ConcurrentSet {
	return &ConcurrentSet{
		elems: make(map[Any]Void),
	}
}

func (set *ConcurrentSet) Len() int {
	set.locker.RLock()
	defer set.locker.RUnlock()

	return len(set.elems)
}

func (set *ConcurrentSet) Contains(elem Any) bool {
	set.locker.RLock()
	defer set.locker.RUnlock()

	_, exist := set.elems[elem]
	return exist
}

func (set *ConcurrentSet) Insert(elem Any) {
	set.locker.Lock()
	defer set.locker.Unlock()

	set.elems[elem] = Void{}
}

func (set *ConcurrentSet) Delete(elem Any) {
	set.locker.Lock()
	defer set.locker.Unlock()

	delete(set.elems, elem)
}

func (set *ConcurrentSet) GetAll() []Any {
	set.locker.RLock()
	defer set.locker.RUnlock()

	elems := make([]Any, 0, len(set.elems))
	for elem, _ := range set.elems {
		elems = append(elems, elem)
	}

	return elems
}

// Take care NOT to run Long-Time routine
func (set *ConcurrentSet) ForEach(visit func(Any)) {
	set.locker.RLock()
	defer set.locker.RUnlock()

	for elem, _ := range set.elems {
		visit(elem)
	}
}
