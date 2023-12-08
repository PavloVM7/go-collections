package collections

import (
	"reflect"
	"testing"
)

func TestNewSet(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		want Set[T]
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSet[int](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSetCapacity(t *testing.T) {
	type args struct {
		capacity int
	}
	type testCase[T comparable] struct {
		name string
		args args
		want Set[T]
	}
	tests := []testCase[any]{
		{"zero", args{capacity: 0}, Set[any]{mp: make(map[any]struct{}), capacity: 0}},
		{"less then zero", args{capacity: -1}, Set[any]{mp: make(map[any]struct{}), capacity: -1}},
		{"more then zero", args{capacity: 12}, Set[any]{mp: make(map[any]struct{}, 12), capacity: 12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSetCapacity[any](tt.args.capacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSetCapacity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Add(t *testing.T) {
	type args[T comparable] struct {
		value T
	}
	type testCase[T comparable] struct {
		name string
		set  Set[T]
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.set.Add(tt.args.value); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_AddAll(t *testing.T) {
	type args[T comparable] struct {
		values []T
	}
	type testCase[T comparable] struct {
		name string
		set  Set[T]
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.set.AddAll(tt.args.values...); got != tt.want {
				t.Errorf("AddAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Capacity(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		set  Set[T]
		want int
	}
	tests := []testCase[int]{
		{"empty", NewSet[int](), 0},
		{"capacity", NewSetCapacity[int](123), 123},
		{"capacity less then zero", NewSetCapacity[int](-1), -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.set.Capacity(); got != tt.want {
				t.Errorf("Capacity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Clear(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		set  Set[T]
		want int
	}
	tests := []testCase[string]{
		{"empty", NewSet[string](), 0},
		{"one", NewSetItems[string]("one"), 1},
		{"three", NewSetItems[string]("one", "two", "three"), 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.set.Size() != tt.want {
				t.Fatalf("expected: %d, actual: %d", tt.want, tt.set.Size())
			}
			tt.set.Clear()
			if tt.set.Size() != 0 {
				t.Fatalf("the set was not cleared, size: %d", tt.set.Size())
			}
		})
	}
}

func TestSet_IsEmpty(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		set  Set[T]
		want bool
	}
	tests := []testCase[string]{
		{"empty", NewSet[string](), true},
		{"empty with capacity", NewSetCapacity[string](17), true},
		{"one", NewSetItems[string]("string 1"), false},
		{"three", NewSetItems[string]("string 1", "two string", "three"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.set.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Remove(t *testing.T) {
	t.Fail()
}

func TestSet_Size(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		set  Set[T]
		want int
	}
	tests := []testCase[int]{
		{"empty", NewSet[int](), 0},
		{"empty with capacity", NewSetCapacity[int](123), 0},
		{"one", NewSetItems[int](1), 0},
		{"three", NewSetItems[int](1, 2, 3), 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.set.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_TrimToSize(t *testing.T) {
	t.Fail()
}
