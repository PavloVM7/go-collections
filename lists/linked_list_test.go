// Copyright Ⓒ 2023 Pavlo Moisieienko. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lists

import (
	"testing"
)

func TestLinkedList_AddFirst(t *testing.T) {
	list := NewLinkedList[int]()
	list.AddFirst(1)
	list.AddFirst(2)
	list.AddFirst(3)
	if list.Size() != 3 {
		t.Fatal("wrong size, expected:", 3, "actual:", list.Size())
	}
	actual, ok := list.GetFirst()
	if !ok {
		t.Fatal("a value does not exists")
	}
	if actual != 3 {
		t.Fatal("wrong value, expected:", 3, "actual:", actual)
	}

	last, lok := list.GetLast()
	if !lok {
		t.Fatal("the last value doesn't exists")
	}
	if last != 1 {
		t.Fatal("wrong the last value, expected:", 1, "actual:", last)
	}
}
func TestLinkedList_AddFirst_first(t *testing.T) {
	list := NewLinkedList[int]()
	list.AddFirst(1)
	if list.Size() != 1 {
		t.Fatal("wrong size, expected:", 1, "actual:", list.Size())
	}
	actual, ok := list.GetFirst()
	if !ok {
		t.Fatal("The value has not been added")
	}
	if actual != 1 {
		t.Fatal("wrong value, expected:", 1, "actual:", actual)
	}
	last, lok := list.GetLast()
	if !lok {
		t.Fatal("last value does not exists")
	}
	if last != 1 {
		t.Fatal("wrong last value, expected:", 1, "actual:", last)
	}
	if list.first != list.last {
		t.Fatal("the last and first values are not the same")
	}
}

func TestLinkedList_GetLast_empty_list(t *testing.T) {
	list := NewLinkedList[*listTestStruct]()
	actual, ok := list.GetLast()
	if ok {
		t.Fatal("the item exists")
	}
	if actual != nil {
		t.Fatal("nil value is expected")
	}
}

func TestLinkedList_GetLast_empty_list_not_nil(t *testing.T) {
	list := NewLinkedList[int]()
	actual, ok := list.GetLast()
	if ok {
		t.Fatal("the item exists")
	}
	if actual != 0 {
		t.Fatal("0 value is expected")
	}
}

func TestLinkedList_GetFirst_empty_list(t *testing.T) {
	list := NewLinkedList[*listTestStruct]()
	actual, ok := list.GetFirst()
	if ok {
		t.Fatal("the item exists")
	}
	if actual != nil {
		t.Fatal("nil value is expected")
	}
}
func TestLinkedList_GetFirst_empty_list_not_nil(t *testing.T) {

	list := NewLinkedList[int]()
	actual, ok := list.GetFirst()
	if ok {
		t.Fatal("the item exists")
	}
	if actual != 0 {
		t.Fatal("0 value is expected")
	}
}

func TestNewLinkedList(t *testing.T) {
	list := NewLinkedList[string]()
	if list.first != nil {
		t.Fatal("first does not equal nil")
	}
	if list.last != nil {
		t.Fatal("last does not equal nil")
	}
	if list.size != 0 {
		t.Fatal("the size does not equal 0")
	}
}

type listTestStruct struct {
	name  string
	value int
}
