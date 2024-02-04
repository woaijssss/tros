package structure

import "sort"

type OrderedSetElemType interface {
	byte | int8 | int16 | int32 | int64 | int | string
}

// OrderedSet 有序集合
type OrderedSet[T OrderedSetElemType] interface {
	Exist(target T) bool
	AddOne(elem T)
	Add(elems ...T)
	Len() int
	Cap() int
	AllElements() []T
}

func NewOrderedSet[T OrderedSetElemType]() OrderedSet[T] {
	return &orderedSet[T]{
		mset: make(map[T]int),
	}
}

func NewOrderedSetWithCap[T OrderedSetElemType](cap int) OrderedSet[T] {
	return &orderedSet[T]{
		mset: make(map[T]int, cap),
		list: make([]T, 0, cap),
	}
}

type orderedSet[T OrderedSetElemType] struct {
	mset map[T]int
	list []T
}

func (set *orderedSet[T]) Exist(target T) bool {
	if _, ok := set.mset[target]; ok {
		return true
	}
	return false
}

func (set *orderedSet[T]) AddOne(elem T) {
	if _, ok := set.mset[elem]; ok {
		return
	}
	//set.list = append(set.list, elem)
	idx := set.findBestPos(elem)
	if idx < set.Len() {
		set.list = append(set.list[:idx], append([]T{elem}, set.list[idx:]...)...)

	} else {
		set.list = append(set.list, elem)
	}
	set.mset[elem] = 1
}

func (set *orderedSet[T]) Add(elems ...T) {
	for _, elem := range elems {
		set.AddOne(elem)
	}
}

func (set *orderedSet[T]) Len() int {
	return len(set.list)
}

func (set *orderedSet[T]) Cap() int {
	return cap(set.list)
}

func (set *orderedSet[T]) AllElements() []T {
	return set.list
}

func (set *orderedSet[T]) findBestPos(target T) int {
	return sort.Search(len(set.list), func(i int) bool { return set.list[i] > target })
}
