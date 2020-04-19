package types

/************************* Basic Set *************************/

type Set struct {
	elems map[interface{}]bool
}

func NewSet() *Set {
	return &Set{
		elems: make(map[interface{}]bool),
	}
}

func NewSetWithInitSize(size int) *Set {
	return &Set{
		elems: make(map[interface{}]bool, size),
	}
}

func (set *Set) Len() int {
	return len(set.elems)
}

func (set *Set) Contains(elem interface{}) bool {
	_, exist := set.elems[elem]
	return exist
}

func (set *Set) Insert(elem interface{}) {
	set.elems[elem] = true
}

func (set *Set) Delete(elem interface{}) {
	delete(set.elems, elem)
}

func (set *Set) Foreach(visit func (interface{})) {
	for elem, _ := range set.elems {
		visit(elem)
	}
}

func (set *Set) Filter(predict func (interface{}) bool) []interface{} {
	filtered := make([]interface{}, 0, len(set.elems))
	for elem, _ := range set.elems {
		if predict(elem) {
			filtered = append(filtered, elem)
		}
	}
	return filtered
}

func (set *Set) DeleteIf(predict func (interface{}) bool) {
	for elem, _ := range set.elems {
		if predict(elem) {
			delete(set.elems, elem)
		}
	}
}
