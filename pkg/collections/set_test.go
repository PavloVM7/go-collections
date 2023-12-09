package collections

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

func TestNewSet(t *testing.T) {
	set := NewSet[int]()
	if set.Size() != 0 {
		t.Fatalf("invalid size, expected: %d, actual: %d", 0, set.Size())
	}
	if !set.IsEmpty() {
		t.Fatal("the set isn't empty")
	}
	if set.Capacity() != 0 {
		t.Fatalf("invalid capacity, expected: %d, actual: %d", 0, set.Capacity())
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
	set := NewSet[int]()
	values := []int{1, 2, 3}
	for _, v := range values {
		added := set.Add(v)
		if !added {
			t.Fatalf("value %v was not added to the set", v)
		}
	}
	if set.Size() != len(values) {
		t.Fatalf("invalid set size, expected: %v, actual: %v", len(values), set.Size())
	}
	for _, v := range values {
		added := set.Add(v)
		if added {
			t.Fatalf("dublicate value %v was added to the set", v)
		}
	}
	if set.Size() != len(values) {
		t.Fatalf("invalid set size, expected: %v, actual: %v", len(values), set.Size())
	}
}

func TestSet_AddAll(t *testing.T) {
	set := NewSet[string]()
	values := []string{"string 1", "string 2", "string 3"}
	values2 := []string{"string 4", "string 5"}

	changed := set.AddAll(values...)
	if !changed {
		t.Fatalf("the set was not changed")
	}
	if set.Size() != len(values) {
		t.Fatalf("invalid size, expected: %d, actual: %d", len(values), set.Size())
	}
	changed = set.AddAll(values...)
	if changed {
		t.Fatalf("the set was changed when trying to add duplicate values")
	}
	changed = set.AddAll(values2...)
	if !changed {
		t.Fatalf("the set was not changed")
	}
	expectedSize := len(values) + len(values2)
	if set.Size() != expectedSize {
		t.Fatalf("invalid size, expected: %d, actual: %d", expectedSize, set.Size())
	}
	changed = set.AddAll(values2...)
	if changed {
		t.Fatalf("the set was changed when trying to add duplicate values")
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
	set := NewSet[int]()
	values := []int{1, 2, 3}
	set.AddAll(values...)
	if set.Size() != len(values) {
		t.Fatalf("invalid set size, expected: %d, actual: %d", len(values), set.Size())
	}
	for _, value := range values {
		removed := set.Remove(value)
		if !removed {
			t.Fatalf("known value %v was not removed", value)
		}
	}
	if !set.IsEmpty() {
		t.Fatal("set is not empty")
	}
	removed := set.Remove(111)
	if removed {
		t.Fatal("unknown value was removed")
	}
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
		{"one", NewSetItems[int](1), 1},
		{"three", NewSetItems[int](1, 2, 3), 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.set.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Contains(t *testing.T) {
	values := []string{"string 1", "string 2", "string 3"}
	set := NewSetCapacity[string](len(values))
	if !set.AddAll(values...) {
		t.Fatalf("values was not added to the set")
	}
	for _, value := range values {
		if !set.Contains(value) {
			t.Fatalf("the set does not contain value %s", value)
		}
	}
	unknown := "unknown string value"
	if set.Contains(unknown) {
		t.Fatal("the set contains an unknown value")
	}
}

func TestSet_TrimToSize(t *testing.T) {
	const amount = 1_000
	const rest = 20
	set := NewSetCapacity[string](amount)
	value := func(i int) string {
		return fmt.Sprintf("this is a set long value %d", i)
	}
	for i := 1; i <= amount; i++ {
		v := value(i)
		if !set.Add(v) {
			t.Fatalf("value %v was not added to the set", v)
		}
	}
	if set.Size() != amount {
		t.Fatalf("invalid set size, expected: %d, actual: %d", amount, set.Size())
	}
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)

	for i := rest + 1; i <= amount; i++ {
		v := value(i)
		if !set.Remove(v) {
			t.Fatalf("value %s was not removed from the set", v)
		}
	}
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)

	runtime.GC()

	var m3 runtime.MemStats
	runtime.ReadMemStats(&m3)

	set.TrimToSize()

	var m4 runtime.MemStats
	runtime.ReadMemStats(&m4)

	runtime.GC()

	var m5 runtime.MemStats
	runtime.ReadMemStats(&m5)

	memToString := func(ms *runtime.MemStats) string {
		return fmt.Sprintf("%d Kb", ms.Alloc/1024)
	}

	t.Logf("Memory after fill: %s; after remove: %s (GC: %s); after trim: %s (GC: %s)",
		memToString(&m1), memToString(&m2), memToString(&m3), memToString(&m4), memToString(&m5))

	if set.Size() != rest {
		t.Fatalf("invalid set size, expected: %d, actual: %d", rest, set.Size())
	}
	for i := 1; i <= rest; i++ {
		v := value(i)
		if !set.Contains(v) {
			t.Fatalf("the set does not contain value %s", v)
		}
	}
}
