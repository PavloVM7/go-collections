// Copyright Ⓒ 2023 Pavlo Moisieienko. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package lists contains ordered collections and their manipulation
package lists

import "errors"

var (
	ErrIndexOutOfRange = errors.New("index out of range")
)

type LinkedList[T any] struct {
	first *listItem[T]
	last  *listItem[T]
	size  int
}

func (list *LinkedList[T]) AddLast(value T) {
	item := &listItem[T]{value: value}
	if list.last != nil {
		list.last.append(item)
		list.last = item
	} else {
		list.first = item
		list.last = item
	}
	list.size++
}
func (list *LinkedList[T]) AddFirst(value T) {
	item := &listItem[T]{value: value}
	if list.first != nil {
		list.first.insert(item)
		list.first = item
	} else {
		list.first = item
		list.last = item
	}
	list.size++
}
func (list *LinkedList[T]) GetFirst() (T, bool) {
	if list.first != nil {
		return list.first.value, true
	}
	var res T
	return res, false
}
func (list *LinkedList[T]) GetLast() (T, bool) {
	if list.last != nil {
		return list.last.value, true
	}
	var res T
	return res, false
}
func (list *LinkedList[T]) Get(index int) (T, error) {
	item, err := list.get(index)
	if err == nil {
		return item.value, nil
	}
	var res T
	return res, err
}
func (list *LinkedList[T]) RemoveFirst() (T, bool) {
	var res T
	if list.first != nil {
		res = list.first.value
		if list.first.next != nil {
			list.first.remove()
			list.first = list.first.next
		} else {
			list.first = nil
			list.last = nil
		}
		list.size--
		return res, true
	}
	return res, false
}
func (list *LinkedList[T]) RemoveLast() (T, bool) {
	var res T
	if list.last != nil {
		res = list.last.value
		if list.last.prev != nil {
			list.last.remove()
			list.last = list.last.prev
		} else {
			list.first = nil
			list.last = nil
		}
		list.size--
		return res, true
	}
	return res, false
}
func (list *LinkedList[T]) Remove(index int) (T, error) {
	item, err := list.get(index)
	var res T
	if err == nil {
		res = item.value
		item.remove()
		if list.first == item {
			list.first = item.next
		}
		if list.last == item {
			list.last = item.prev
		}
		list.size--
	}
	return res, err
}
func (list *LinkedList[T]) get(index int) (*listItem[T], error) {
	if index >= 0 && index < list.size {
		for i, item := 0, list.first; item != nil; i, item = i+1, item.next {
			if i == index {
				return item, nil
			}
		}
	}
	return nil, ErrIndexOutOfRange
}
func (list *LinkedList[T]) ToArray() []T {
	result := make([]T, list.size)
	for i, item := 0, list.first; item != nil; i, item = i+1, item.next {
		result[i] = item.value
	}
	return result
}
func (list *LinkedList[T]) Clear() {
	list.first = nil
	list.last = nil
	list.size = 0
}
func (list *LinkedList[T]) Size() int {
	return list.size
}
func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}
