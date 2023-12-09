// Copyright Ⓒ 2023 Pavlo Moisieienko. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collections

// Set is a collection that does not contain duplicate elements.
// Set is not thread safe and not intended for concurrent usage.
//   - T - value type
type Set[T comparable] struct {
	mp       map[T]struct{}
	capacity int
}

// Add adds a specified value to the set.
// Returns true if the value did not exist and was added to the set, otherwise returns false.
func (set *Set[T]) Add(value T) bool {
	if _, ok := set.mp[value]; !ok {
		set.mp[value] = struct{}{}
		return true
	}
	return false
}

// AddAll adds all the specified values to the Set.
// Returns true if this Set changed as result of the call.
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

// Contains returns true if the set contains the value
func (set *Set[T]) Contains(value T) bool {
	_, ok := set.mp[value]
	return ok
}

// Remove removes a value from the set.
// Returns true if this Set changed as result of the call.
func (set *Set[T]) Remove(value T) bool {
	if _, ok := set.mp[value]; ok {
		delete(set.mp, value)
		return true
	}
	return false
}

// Size returns the current size of the set.
func (set *Set[T]) Size() int {
	return len(set.mp)
}

// IsEmpty returns true if the set does not contain any values.
func (set *Set[T]) IsEmpty() bool {
	return len(set.mp) == 0
}

// TrimToSize trims the capacity of this Set instance to be set's current size.
// An application chan use this operation to minimize the storage of a Set instance.
func (set *Set[T]) TrimToSize() {
	set.mp = CopyMap(set.mp)
}

// Clear clears the Set.
func (set *Set[T]) Clear() {
	if set.capacity > 0 {
		set.mp = make(map[T]struct{}, set.capacity)
	} else {
		set.mp = make(map[T]struct{})
	}
}

// Capacity returns the capacity value that was set when the Set was created.
func (set *Set[T]) Capacity() int {
	return set.capacity
}

// ToSlice return a slice of the set elements.
func (set *Set[T]) ToSlice() []T {
	result := make([]T, 0, len(set.mp))
	for k := range set.mp {
		result = append(result, k)
	}
	return result
}

// NewSet returns a new empty Set instance with capacity equal 0.
//   - T - value type
func NewSet[T comparable]() Set[T] {
	return NewSetCapacity[T](0)
}

// NewSetCapacity returns a new empty Set instance with an initial space size (capacity)
//   - T - value type
//   - capacity - initial space size
func NewSetCapacity[T comparable](capacity int) Set[T] {
	result := Set[T]{capacity: capacity}
	result.Clear()
	return result
}

// NewSetItems returns a new instance of Set containing specified values.
// The Set capacity is equal to the number of values.
//   - values ...T - values that the Set will contain
func NewSetItems[T comparable](values ...T) Set[T] {
	result := NewSetCapacity[T](len(values))
	result.AddAll(values...)
	return result
}
