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

func TestLinkedList_RemoveAll(t *testing.T) {
	type testCase[T any] struct {
		name       string
		list       *LinkedList[T]
		needRemove func(value T) bool
		want       T
		wantArray  []T
	}
	tests := []testCase[int]{
		{
			name:       "empty",
			list:       NewLinkedList[int](),
			needRemove: func(value int) bool { return value > 0 },
			want:       0,
			wantArray:  []int{},
		},
		{
			name:       "not found",
			list:       NewLinkedListItems[int](1, 2, 3),
			needRemove: func(value int) bool { return value == 5 },
			want:       0,
			wantArray:  []int{1, 2, 3},
		},
		{
			name:       "single value",
			list:       NewLinkedListItems[int](1),
			needRemove: func(value int) bool { return value == 1 },
			want:       1,
			wantArray:  []int{},
		},
		{
			name:       "first value",
			list:       NewLinkedListItems[int](1, 2),
			needRemove: func(value int) bool { return value == 1 },
			want:       1,
			wantArray:  []int{2},
		},
		{
			name:       "last value",
			list:       NewLinkedListItems[int](1, 2),
			needRemove: func(value int) bool { return value == 2 },
			want:       1,
			wantArray:  []int{1},
		},
		{
			name:       "middle value",
			list:       NewLinkedListItems[int](2, 1, 2, 3, 2, 5, 2),
			needRemove: func(value int) bool { return value == 2 },
			want:       4,
			wantArray:  []int{1, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.list.RemoveAll(tt.needRemove)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveAll() got = %v, want %v", got, tt.want)
			}
			gotArray := tt.list.ToArray()
			if !reflect.DeepEqual(gotArray, tt.wantArray) {
				t.Errorf("RemoveAll() gotArray = %v, wantArray %v", gotArray, tt.wantArray)
			}
			if tt.list.Size() != len(gotArray) {
				t.Errorf("RemoveAll() gotSize = %v, wantSize %v", len(gotArray), tt.list.Size())
			}
		})
	}
}
func TestLinkedList_RemoveLastOccurrence(t *testing.T) {
	type testCase[T any] struct {
		name       string
		list       *LinkedList[T]
		needRemove func(value T) bool
		want       T
		wantIndex  int
		wantArray  []T
	}
	tests := []testCase[int]{
		{
			name:       "empty",
			list:       NewLinkedList[int](),
			needRemove: func(value int) bool { return value > 0 },
			want:       0,
			wantIndex:  -1,
			wantArray:  []int{},
		},
		{
			name:       "not found",
			list:       NewLinkedListItems[int](1, 2, 3),
			needRemove: func(value int) bool { return value == 5 },
			want:       0,
			wantIndex:  -1,
			wantArray:  []int{1, 2, 3},
		},
		{
			name:       "single value",
			list:       NewLinkedListItems[int](1),
			needRemove: func(value int) bool { return value == 1 },
			want:       1,
			wantIndex:  0,
			wantArray:  []int{},
		},
		{
			name:       "first value",
			list:       NewLinkedListItems[int](1, 2),
			needRemove: func(value int) bool { return value == 1 },
			want:       1,
			wantIndex:  0,
			wantArray:  []int{2},
		},
		{
			name:       "last value",
			list:       NewLinkedListItems[int](1, 2),
			needRemove: func(value int) bool { return value == 2 },
			want:       2,
			wantIndex:  1,
			wantArray:  []int{1},
		},
		{
			name:       "middle value",
			list:       NewLinkedListItems[int](1, 2, 3, 2, 5),
			needRemove: func(value int) bool { return value == 2 },
			want:       2,
			wantIndex:  3,
			wantArray:  []int{1, 2, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotIndex := tt.list.RemoveLastOccurrence(tt.needRemove)
			if !reflect.DeepEqual(gotValue, tt.want) {
				t.Errorf("RemoveLastOccurrence() got = %v, want %v", gotValue, tt.want)
			}
			if gotIndex != tt.wantIndex {
				t.Errorf("RemoveLastOccurrence() gotIndex = %v, wantIndex %v", gotIndex, tt.wantIndex)
			}
			gotArray := tt.list.ToArray()
			if !reflect.DeepEqual(gotArray, tt.wantArray) {
				t.Errorf("RemoveLastOccurrence() gotArray = %v, wantArray %v", gotArray, tt.wantArray)
			}
			if tt.list.Size() != len(gotArray) {
				t.Errorf("RemoveLastOccurrence() gotSize = %v, wantSize %v", len(gotArray), tt.list.Size())
			}
		})
	}
}
func TestLinkedList_RemoveFirstOccurrence(t *testing.T) {
	type testCase[T any] struct {
		name       string
		list       *LinkedList[T]
		needRemove func(value T) bool
		want       T
		wantIndex  int
		wantArray  []T
	}
	tests := []testCase[int]{
		{
			name:       "empty",
			list:       NewLinkedList[int](),
			needRemove: func(value int) bool { return value > 0 },
			want:       0,
			wantIndex:  -1,
			wantArray:  []int{},
		},
		{
			name:       "not found",
			list:       NewLinkedListItems[int](1, 2, 3),
			needRemove: func(value int) bool { return value == 5 },
			want:       0,
			wantIndex:  -1,
			wantArray:  []int{1, 2, 3},
		},
		{
			name:       "single value",
			list:       NewLinkedListItems[int](1),
			needRemove: func(value int) bool { return value == 1 },
			want:       1,
			wantIndex:  0,
			wantArray:  []int{},
		},
		{
			name:       "first value",
			list:       NewLinkedListItems[int](1, 2),
			needRemove: func(value int) bool { return value == 1 },
			want:       1,
			wantIndex:  0,
			wantArray:  []int{2},
		},
		{
			name:       "last value",
			list:       NewLinkedListItems[int](1, 2),
			needRemove: func(value int) bool { return value == 2 },
			want:       2,
			wantIndex:  1,
			wantArray:  []int{1},
		},
		{
			name:       "middle value",
			list:       NewLinkedListItems[int](1, 2, 3, 2),
			needRemove: func(value int) bool { return value == 2 },
			want:       2,
			wantIndex:  1,
			wantArray:  []int{1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotIndex := tt.list.RemoveFirstOccurrence(tt.needRemove)
			if !reflect.DeepEqual(gotValue, tt.want) {
				t.Errorf("RemoveFirstOccurrence() got = %v, want %v", gotValue, tt.want)
			}
			if gotIndex != tt.wantIndex {
				t.Errorf("RemoveFirstOccurrence() gotIndex = %v, wantIndex %v", gotIndex, tt.wantIndex)
			}
			gotArray := tt.list.ToArray()
			if !reflect.DeepEqual(gotArray, tt.wantArray) {
				t.Errorf("RemoveFirstOccurrence() gotArray = %v, wantArray %v", gotArray, tt.wantArray)
			}
			if tt.list.Size() != len(gotArray) {
				t.Errorf("RemoveFirstOccurrence() gotSize = %v, wantSize %v", len(gotArray), tt.list.Size())
			}
		})
	}
}

func TestLinkedList_Remove(t *testing.T) {
	const (
		want1 = 1
		want2 = 2
		want3 = 3
		want4 = 4
	)
	list := NewLinkedList[int]()
	list.AddLast(want1)
	list.AddLast(want2)
	list.AddLast(want3)
	list.AddLast(want4)
	got3, _ := list.Remove(2)
	if got3 != want3 {
		t.Fatalf("unexpected value: %v, want: %v", got3, want3)
	}
	gotAr1 := list.ToArray()
	wantAr1 := []int{1, 2, 4}
	if !reflect.DeepEqual(gotAr1, wantAr1) {
		t.Fatalf("Remove() got: %v, want: %v", gotAr1, wantAr1)
	}

	got1, _ := list.Remove(0)
	if got1 != want1 {
		t.Fatalf("unexpected value: %v, want: %v", got1, want1)
	}
	gotAr2 := list.ToArray()
	wantAr2 := []int{2, 4}
	if !reflect.DeepEqual(gotAr2, wantAr2) {
		t.Fatalf("Remove() got: %v, want: %v", gotAr2, wantAr2)
	}

	got4, _ := list.Remove(1)
	if got4 != want4 {
		t.Fatalf("unexpected value: %v, want: %v", got4, want4)
	}
	gotAr3 := list.ToArray()
	wantAr3 := []int{2}
	if !reflect.DeepEqual(gotAr3, wantAr3) {
		t.Fatalf("Remove() got: %v, want: %v", gotAr3, wantAr3)
	}

	got2, _ := list.Remove(0)
	if got2 != want2 {
		t.Fatalf("unexpected value: %v, want: %v", got2, want2)
	}
	gotAr4 := list.ToArray()
	if len(gotAr4) != 0 {
		t.Fatal("empty array expected, got:", gotAr4)
	}
}
func TestLinkedList_Remove_last(t *testing.T) {
	const expected1 = "value 1"
	const expected2 = "value 2"
	list := NewLinkedList[string]()
	list.AddLast(expected1)
	list.AddLast(expected2)
	actual, err := list.Remove(list.Size() - 1)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	if actual != expected2 {
		t.Fatalf("unexpected value: '%s', want: '%s'", actual, expected2)
	}
	if list.Size() != 1 {
		t.Fatalf("unexpected list size: %d, want: %d", list.Size(), 1)
	}
	first, _ := list.GetFirst()
	if first != expected1 {
		t.Fatalf("unexpected first value: '%v'; want: '%v'", first, expected1)
	}
	last, _ := list.GetLast()
	if last != expected1 {
		t.Fatalf("unexpected first value: '%v'; want: '%v'", last, expected1)
	}
	if list.first != list.last {
		t.Fatalf("values 'first' and 'last' must be the same; actual values: first: %v, last: %v",
			list.first, list.last)
	}
}
func TestLinkedList_Remove_single(t *testing.T) {
	const expected = "single value"
	list := NewLinkedList[string]()
	list.AddLast(expected)
	actual, err := list.Remove(0)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	if actual != expected {
		t.Fatalf("unexpected value: '%s', expected: '%s'", actual, expected)
	}
	if list.Size() != 0 {
		t.Fatalf("unexpected list size: %d, want: %d", list.Size(), 0)
	}
	if list.first != nil {
		t.Fatal("the first item should be nil, actual:", list.first)
	}
	if list.last != nil {
		t.Fatal("the last value should be nil, actual:", list.last)
	}
}
func TestLinkedList_Remove_fail(t *testing.T) {
	list := NewLinkedList[string]()
	actual, err := list.Remove(0)
	if !errors.Is(err, ErrIndexOutOfRange) {
		t.Fatalf("expected error: '%v', got: '%v'", ErrIndexOutOfRange, err)
	}
	if actual != "" {
		t.Fatalf("expected: '', actual: '%s'", actual)
	}
	list.AddLast("value")
	actual, err = list.Remove(1)
	if !errors.Is(err, ErrIndexOutOfRange) {
		t.Fatalf("expected error: '%v', got: '%v'", ErrIndexOutOfRange, err)
	}
	if actual != "" {
		t.Fatalf("expected: '', actual: '%s'", actual)
	}
}

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
			t.Fatalf("incorrect value: %v, want: %v", actual, expectedValue)
		}
		expectedSize--
		if list.Size() != expectedSize {
			t.Fatalf("incorrect list size: %v, want: %v", list.Size(), expectedSize)
		}
	}
	actual, ok := list.RemoveFirst()
	if ok {
		t.Fatal("the list must be empty")
	}
	if actual != 0 {
		t.Fatalf("incorrect value: %v, want: %v", actual, 0)
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
		t.Fatalf("expected: %v, actual: %v", 1, actual)
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
		t.Fatalf("incorrect last value: %v, want: %v", last, 2)
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

func TestLinkedList_RemoveLast(t *testing.T) {
	list := NewLinkedList[int]()
	list.AddFirst(1)
	list.AddFirst(2)
	list.AddFirst(3)
	expectedSize := list.Size()
	for i := 0; i < 3; i++ {
		actual, ok := list.RemoveLast()
		if !ok {
			t.Fatal("the last element must exist")
		}
		expectedValue := i + 1
		if actual != expectedValue {
			t.Fatalf("incorrect value: %v, want: %v", actual, expectedValue)
		}
		expectedSize--
		if list.Size() != expectedSize {
			t.Fatalf("incorrect list size: %v, want: %v", list.Size(), expectedSize)
		}
	}
	actual, ok := list.RemoveLast()
	if ok {
		t.Fatal("the list must be empty")
	}
	if actual != 0 {
		t.Fatalf("incorrect value: %v, want: %v", actual, 0)
	}
}
func TestLinkedList_RemoveLast_before_last(t *testing.T) {
	list := NewLinkedList[int]()
	list.AddLast(2)
	list.AddFirst(1)
	if list.Size() != 2 {
		t.Fatal("unexpected list size:", list.Size(), "expected:", 2)
	}
	actual, ok := list.RemoveLast()
	if !ok {
		t.Fatalf("unexpected value: %v, expected: true", ok)
	}
	if actual != 2 {
		t.Fatalf("expected: %v, actual: %v", 2, actual)
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
	first, _ := list.GetFirst()
	if first != 1 {
		t.Fatalf("incorrect last value: %v, want: %v", first, 1)
	}
}
func TestLinkedList_RemoveLast_single(t *testing.T) {
	list := NewLinkedList[int]()
	list.AddFirst(1)
	if list.Size() != 1 {
		t.Fatal("unexpected list size:", list.Size(), "expected:", 1)
	}
	actual, ok := list.RemoveLast()
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
func TestLinkedList_RemoveLast_empty(t *testing.T) {
	list := NewLinkedList[int]()
	actual, ok := list.RemoveLast()
	if list.Size() != 0 {
		t.Fatal("incorrect list size, expected:", 0, "actual:", list.Size())
	}
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
		t.Fatal("incorrect list size, expected:", 5, "actual:", list.Size())
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
		t.Fatal("incorrect size, expected:", 3, "actual:", list.Size())
	}
	actual := list.ToArray()
	if len(actual) != list.Size() {
		t.Fatal("incorrect array size, expected:", list.Size(), "actual:", len(actual))
	}
	expected := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("incorrect array, expected:", expected, "actual:", actual)
	}
}

func TestLinkedList_AddLast(t *testing.T) {
	list := NewLinkedList[int]()
	list.AddLast(1)
	list.AddLast(2)
	list.AddLast(3)
	if list.Size() != 3 {
		t.Fatal("incorrect size, expected:", 3, "actual:", list.Size())
	}
	first, fok := list.GetFirst()
	if !fok {
		t.Fatal("the value has not been added")
	}
	if first != 1 {
		t.Fatal("incorrect first value, expected:", 1, "actual:", first)
	}
	last, lok := list.GetLast()
	if !lok {
		t.Fatal("last value has not been added")
	}
	if last != 3 {
		t.Fatal("incorrect last value, expected:", 3, "actual:", last)
	}
}
func TestLinkedList_AddLast_first(t *testing.T) {
	list := NewLinkedList[int]()
	list.AddLast(1)
	if list.Size() != 1 {
		t.Fatal("incorrect size, expected:", 1, "actual:", list.Size())
	}
	actual, ok := list.GetFirst()
	if !ok {
		t.Fatal("The value has not been added")
	}
	if actual != 1 {
		t.Fatal("incorrect value, expected:", 1, "actual:", actual)
	}
	last, lok := list.GetLast()
	if !lok {
		t.Fatal("last value does not exists")
	}
	if last != 1 {
		t.Fatal("incorrect last value, expected:", 1, "actual:", last)
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
		t.Fatal("incorrect size, expected:", 3, "actual:", list.Size())
	}
	actual, ok := list.GetFirst()
	if !ok {
		t.Fatal("a value does not exists")
	}
	if actual != 3 {
		t.Fatal("incorrect value, expected:", 3, "actual:", actual)
	}

	last, lok := list.GetLast()
	if !lok {
		t.Fatal("the last value doesn't exists")
	}
	if last != 1 {
		t.Fatal("incorrect last value, expected:", 1, "actual:", last)
	}
}
func TestLinkedList_AddFirst_first(t *testing.T) {
	list := NewLinkedList[int]()
	list.AddFirst(1)
	if list.Size() != 1 {
		t.Fatal("incorrect size, expected:", 1, "actual:", list.Size())
	}
	actual, ok := list.GetFirst()
	if !ok {
		t.Fatal("The value has not been added")
	}
	if actual != 1 {
		t.Fatal("incorrect value, expected:", 1, "actual:", actual)
	}
	last, lok := list.GetLast()
	if !lok {
		t.Fatal("last value does not exists")
	}
	if last != 1 {
		t.Fatal("incorrect last value, expected:", 1, "actual:", last)
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
