// Copyright Ⓒ 2023 Pavlo Moisieienko. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lists

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestLinkedList_RemoveFirst(t *testing.T) {
	list := NewLinkedList[int]()
	list.AddLast(1)
	list.AddLast(2)
	list.AddLast(3)
	expectedSize := list.Size()
	for i := 0; i < 3; i++ {
		actual, ok := list.RemoveFirst()
		if !ok {
			t.Fatal("the first element must exist")
		}
		expectedValue := i + 1
		if actual != expectedValue {
			t.Fatalf("wrong value: %v, want: %v", actual, expectedValue)
		}
		expectedSize--
		if list.Size() != expectedSize {
			t.Fatalf("wrong list size: %v, want: %v", list.Size(), expectedSize)
		}
	}
	actual, ok := list.RemoveFirst()
	if ok {
		t.Fatal("the list must be empty")
	}
	if actual != 0 {
		t.Fatalf("wrong value: %v, want: %v", actual, 0)
	}
}
func TestLinkedList_RemoveFirst_before_last(t *testing.T) {
	list := NewLinkedList[int]()
	list.AddFirst(1)
	list.AddLast(2)
	if list.Size() != 2 {
		t.Fatal("unexpected list size:", list.Size(), "expected:", 2)
	}
	actual, ok := list.RemoveFirst()
	if !ok {
		t.Fatalf("unexpected value: %v, expected: true", ok)
	}
	if actual != 1 {
		t.Fatalf("expected: %v, actual: %v", 0, actual)
	}
	if list.Size() != 1 {
		t.Fatalf("unexpeted list size: %v, want: %v", list.Size(), 1)
	}
	if list.first == nil {
		t.Fatal("the first element must exist")
	}
	if list.last == nil {
		t.Fatal("the last element must exist")
	}
	if list.first.prev != nil {
		t.Fatal("the 'prev' value of the first element must be nil")
	}
	if list.first.next != nil {
		t.Fatal("the 'next' value of the first element must be nil")
	}
	if list.last.prev != nil {
		t.Fatal("the 'prev' value of the last element must be nil")
	}
	if list.last.next != nil {
		t.Fatal("the 'next' value of the last element must be nil")
	}
	if list.last != list.first {
		t.Fatal("values 'first' and 'last' must be the same")
	}
	last, _ := list.GetLast()
	if last != 2 {
		t.Fatalf("wrong last value: %v, want: %v", list, 2)
	}
}
func TestLinkedList_RemoveFirst_single(t *testing.T) {
	list := NewLinkedList[int]()
	list.AddFirst(1)
	if list.Size() != 1 {
		t.Fatal("unexpected list size:", list.Size(), "expected:", 1)
	}
	actual, ok := list.RemoveFirst()
	if !ok {
		t.Fatalf("unexpected value: %v, expected: true", ok)
	}
	if actual != 1 {
		t.Fatalf("expected: %v, actual: %v", 0, actual)
	}
	if list.Size() != 0 {
		t.Fatalf("unexpeted list size: %v, want: %v", list.Size(), 0)
	}
	if list.first != nil {
		t.Fatal("the first item should be nil, actual:", list.first)
	}
	if list.last != nil {
		t.Fatal("the last value should be nil, actual:", list.last)
	}
}
func TestLinkedList_RemoveFirst_empty(t *testing.T) {
	list := NewLinkedList[int]()
	actual, ok := list.RemoveFirst()
	if ok {
		t.Fatalf("unexpected value: %v, expected: false", ok)
	}
	if actual != 0 {
		t.Fatalf("unexpected value: %v, expected: %v", actual, 0)
	}
}

func TestLinkedList_Get(t *testing.T) {
	crt := func(num int) string {
		return fmt.Sprint("list item ", num)
	}
	list := NewLinkedList[string]()
	list.AddFirst(crt(3))
	list.AddFirst(crt(2))
	list.AddFirst(crt(1))
	list.AddLast(crt(4))
	list.AddLast(crt(5))
	if list.Size() != 5 {
		t.Fatal("wrong list size, expected:", 5, "actual:", list.Size())
	}
	for i := 0; i < list.Size(); i++ {
		actual, err := list.Get(i)
		if err != nil {
			t.Fatal("unexpected error:", err, "i:", i)
		}
		expected := crt(i + 1)
		if actual != expected {
			t.Fatalf("i: %d; unexpected value: '%s', expect: '%s'", i, actual, expected)
		}
	}
}
func TestLinkedList_Get_fail(t *testing.T) {
	list := NewLinkedList[string]()
	val, err := list.Get(-1)
	if !errors.Is(err, ErrIndexOutOfRange) {
		t.Fatalf("expected error: '%v', got: '%v'", ErrIndexOutOfRange, err)
	}
	if val != "" {
		t.Fatalf("expected: '', actual: '%s'", val)
	}
}

func TestLinkedList_ToArray_empty(t *testing.T) {
	list := NewLinkedList[int]()
	actual := list.ToArray()
	if len(actual) != 0 {
		t.Fatal("an empty array is expected")
	}
}
func TestLinkedList_ToArray(t *testing.T) {
	list := NewLinkedList[int]()
	list.AddFirst(3)
	list.AddFirst(2)
	list.AddFirst(1)
	list.AddLast(4)
	list.AddLast(5)
	if list.Size() != 5 {
		t.Fatal("wrong size, expected:", 3, "actual:", list.Size())
	}
	actual := list.ToArray()
	if len(actual) != list.Size() {
		t.Fatal("wrong array size, expected:", list.Size(), "actual:", len(actual))
	}
	expected := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("wrong array, expected:", expected, "actual:", actual)
	}
}

func TestLinkedList_AddLast(t *testing.T) {
	list := NewLinkedList[int]()
	list.AddLast(1)
	list.AddLast(2)
	list.AddLast(3)
	if list.Size() != 3 {
		t.Fatal("wrong size, expected:", 3, "actual:", list.Size())
	}
	first, fok := list.GetFirst()
	if !fok {
		t.Fatal("the value has not been added")
	}
	if first != 1 {
		t.Fatal("wrong first value, expected:", 1, "actual:", first)
	}
	last, lok := list.GetLast()
	if !lok {
		t.Fatal("last value has not been added")
	}
	if last != 3 {
		t.Fatal("wrong last value, expected:", 3, "actual:", last)
	}
}
func TestLinkedList_AddLast_first(t *testing.T) {
	list := NewLinkedList[int]()
	list.AddLast(1)
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
		t.Fatal("wrong last value, expected:", 1, "actual:", last)
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
