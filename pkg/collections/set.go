package collections

type Set[T comparable] struct {
	mp       map[T]struct{}
	capacity int
}

func (set *Set[T]) Add(value T) bool {
	if _, ok := set.mp[value]; !ok {
		set.mp[value] = struct{}{}
		return true
	}
	return false
}
func (set *Set[T]) AddAll(values ...T) bool {
	var changed bool
	for _, value := range values {
		if _, ok := set.mp[value]; !ok {
			set.mp[value] = struct{}{}
			changed = true
		}
	}
	return changed
}
func (set *Set[T]) Remove(value T) bool {
	if _, ok := set.mp[value]; ok {
		delete(set.mp, value)
		return true
	}
	return false
}

func (set *Set[T]) Size() int {
	return len(set.mp)
}
func (set *Set[T]) IsEmpty() bool {
	return len(set.mp) == 0
}
func (set *Set[T]) TrimToSize() {
	set.mp = CopyMap(set.mp)
}
func (set *Set[T]) Clear() {
	if set.capacity > 0 {
		set.mp = make(map[T]struct{}, set.capacity)
	} else {
		set.mp = make(map[T]struct{})
	}
}
func (set *Set[T]) Capacity() int {
	return set.capacity
}
func (set *Set[T]) ToArray() []T {
	result := make([]T, 0, len(set.mp))
	for k := range set.mp {
		result = append(result, k)
	}
	return result
}
func NewSet[T comparable]() Set[T] {
	return NewSetCapacity[T](0)
}
func NewSetCapacity[T comparable](capacity int) Set[T] {
	result := Set[T]{capacity: capacity}
	result.Clear()
	return result
}
func NewSetItems[T comparable](values ...T) Set[T] {
	result := NewSetCapacity[T](len(values))
	result.AddAll(values...)
	return result
}
