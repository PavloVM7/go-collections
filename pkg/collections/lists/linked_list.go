// Copyright Ⓒ 2023 Pavlo Moisieienko. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package lists contains ordered collections and their manipulation
package lists

import "errors"

var (
	// ErrIndexOutOfRange error: 'index is out of range'
	ErrIndexOutOfRange = errors.New("index is out of range")
)

// LinkedList is an implementation of a doubly-linked list.
type LinkedList[T any] struct {
	first *listItem[T]
	last  *listItem[T]
	size  int
}

// AddLast appends specified element to the end of this list.
//   - value - the value to be appended
func (list *LinkedList[T]) AddLast(value T) {
	item := &listItem[T]{value: value}
	if list.last != nil {
		list.last.append(item)
	} else {
		list.first = item
	}
	list.last = item
	list.size++
}

// AddFirst inserts specified element to the beginning this list.
//   - value - the value to be inserted
func (list *LinkedList[T]) AddFirst(value T) {
	item := &listItem[T]{value: value}
	if list.first != nil {
		list.first.insert(item)
	} else {
		list.last = item
	}
	list.first = item
	list.size++
}

// GetFirst returns the first element of this list and true if it exists.
// If the list is empty, this method returns a default value of type T and false
func (list *LinkedList[T]) GetFirst() (T, bool) {
	if list.first != nil {
		return list.first.value, true
	}
	var res T
	return res, false
}

// GetLast returns the last element of this list and true if it exists.
// If the lis is empty, this method returns a default value of type T and false.
func (list *LinkedList[T]) GetLast() (T, bool) {
	if list.last != nil {
		return list.last.value, true
	}
	var res T
	return res, false
}

// Get returns an item at the specified position in this list
// or a default value of type T and an error if the index is out of range.
func (list *LinkedList[T]) Get(index int) (T, error) {
	item, err := list.getByIndex(index)
	if err == nil {
		return item.value, nil
	}
	var res T
	return res, err
}

// RemoveFirst removes the first item from this list and returns its value and true if it exists.
// If the list is empty, a default value of type T and false is returned.
func (list *LinkedList[T]) RemoveFirst() (T, bool) {
	var res T
	if list.first != nil {
		res = list.removeItem(list.first)
		return res, true
	}
	return res, false
}

// RemoveLast removes the last item from this list and returns its value and true if it exists.
// If the list is empty, a default value of type T and false is returned.
func (list *LinkedList[T]) RemoveLast() (T, bool) {
	var res T
	if list.last != nil {
		res = list.removeItem(list.last)
		return res, true
	}
	return res, false
}

// Remove removes the element at the specified position in this list and returns its value
// or a default value of type T and an error if the index is out of range.
func (list *LinkedList[T]) Remove(index int) (T, error) {
	item, err := list.getByIndex(index)
	var res T
	if err == nil {
		res = list.removeItem(item)
	}
	return res, err
}
func (list *LinkedList[T]) removeItem(item *listItem[T]) T {
	res := item.value
	item.removeYourself()
	if list.first == item {
		list.first = item.next
	}
	if list.last == item {
		list.last = item.prev
	}
	list.size--
	return res
}

// RemoveFirstOccurrence removes from the list the first occurrence of an element that satisfies the condition
// specified by the function (when traversing the list from head to tail).
// Returns the value and index of the removed element, or a default value of type T and -1 if no element was removed.
//   - needToRemove - a function that is applied to each element to determine if it should be deleted
func (list *LinkedList[T]) RemoveFirstOccurrence(needToRemove func(value T) bool) (T, int) {
	var index = -1
	item := list.first
	for item != nil {
		index++
		if needToRemove(item.value) {
			return list.removeItem(item), index
		}
		item = item.next
	}
	var res T
	return res, -1
}

// RemoveLastOccurrence removes from the list the last occurrence of an element that satisfies the condition
// specified by the needToRemove function (when traversing the list from tail to head).
// Returns the value and index of the removed element, or a default value of type T and -1 if no element was removed.
//   - needToRemove - a function that is applied to each element to determine if it should be deleted
func (list *LinkedList[T]) RemoveLastOccurrence(needToRemove func(value T) bool) (T, int) {
	var index = list.size
	item := list.last
	for item != nil {
		index--
		if needToRemove(item.value) {
			return list.removeItem(item), index
		}
		item = item.prev
	}
	var res T
	return res, -1
}

// RemoveAll removes from the list all elements that satisfy the condition specified by the needToRemove function.
// Returns the number of elements removed
//   - needToRemove - a function that is applied to each element to determine if it should be deleted
func (list *LinkedList[T]) RemoveAll(needRemove func(value T) bool) int {
	count := 0
	item := list.first
	for item != nil {
		if needRemove(item.value) {
			list.removeItem(item)
			count++
		}
		item = item.next
	}
	return count
}
func (list *LinkedList[T]) getByIndex(index int) (*listItem[T], error) {
	if index >= 0 && index < list.size {
		for i, item := 0, list.first; item != nil; i, item = i+1, item.next {
			if i == index {
				return item, nil
			}
		}
	}
	return nil, ErrIndexOutOfRange
}

// ToArray returns an array containing all elements of this list in the proper sequence
// (from the first to the last element).
func (list *LinkedList[T]) ToArray() []T {
	result := make([]T, list.size)
	for i, item := 0, list.first; item != nil; i, item = i+1, item.next {
		result[i] = item.value
	}
	return result
}

// Clear clears this list
func (list *LinkedList[T]) Clear() {
	list.first = nil
	list.last = nil
	list.size = 0
}

// Size returns the number of elements in this list
func (list *LinkedList[T]) Size() int {
	return list.size
}

// NewLinkedList constructs an empty list
func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// NewLinkedListItems constructs a list containing the specified elements
func NewLinkedListItems[T any](values ...T) *LinkedList[T] {
	result := NewLinkedList[T]()
	for _, value := range values {
		result.AddLast(value)
	}
	return result
}
