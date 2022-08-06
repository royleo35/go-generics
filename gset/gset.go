package gset

type Set[T comparable] map[T]struct{}

func New[T comparable](data []T) Set[T] {
	res := make(Set[T])
	for _, v := range data {
		res[v] = struct{}{}
	}
	return res
}

func (s Set[T]) Contain(val T) bool {
	_, ok := s[val]
	return ok
}

func (s Set[T]) Add(val T) {
	s[val] = struct{}{}
}

func (s Set[T]) Delete(val T) {
	delete(s, val)
}

func (s Set[T]) Copy() Set[T] {
	res := make(Set[T], len(s))
	for k := range s {
		res[k] = struct{}{}
	}
	return s
}

// 并集
func (s Set[T]) Merge(other Set[T]) Set[T] {
	res := s.Copy()
	for elem := range other {
		res[elem] = struct{}{}
	}
	return res
}

// 交集
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	if len(s) > len(other) {
		return other.Intersection(s)
	}
	res := New(([]T)(nil))
	for elem := range s {
		if _, ok := other[elem]; ok {
			res[elem] = struct{}{}
		}
	}
	return res
}

// 差集 s1 - s2
func (s1 Set[T]) Difference(s2 Set[T]) Set[T] {
	inter := s1.Intersection(s2)
	res := New(([]T)(nil))
	for elem := range s1 {
		// 不属于交集
		if _, ok := inter[elem]; !ok {
			res[elem] = struct{}{}
		}
	}
	return res
}

// 对s1有副作用，但是性能好
func (s1 Set[T]) UnsafeMerge(s2 Set[T]) Set[T] {
	for k := range s2 {
		s1[k] = struct{}{}
	}
	return s1
}

func (s Set[T]) ToSlice() []T {
	res := make([]T, len(s))
	idx := 0
	for k := range s {
		res[idx] = k
		idx++
	}
	return res
}

func (s1 Set[T]) Equal(s2 Set[T]) bool {
	if len(s1) != len(s2) {
		return false
	}
	m := s1.Merge(s2)
	return len(m) == len(s1)
}

func (s1 Set[T]) UnsafeEqual(s2 Set[T]) bool {
	if len(s1) != len(s2) {
		return false
	}
	m := s1.UnsafeMerge(s2)
	return len(m) == len(s1)
}
