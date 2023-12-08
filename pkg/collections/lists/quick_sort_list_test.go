package lists

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_sortItems(t *testing.T) {
	list := NewLinkedListItems[int](3, 5, 1)
	sortItems(list.first, list.last, func(item1, item2 int) bool { return item1 < item2 })
	actual := list.ToArray()
	expected := []int{1, 3, 5}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("sortItems(), got: %v, want: %v", actual, expected)
	}
}

func Test_sortItems1(t *testing.T) {
	type args[T any] struct {
		list *LinkedList[T]
		less func(item1, item2 T) bool
		want []int
	}
	type testCase[T any] struct {
		name string
		args args[T]
	}
	less := func(val1, val2 int) bool { return val1 < val2 }
	tests := []testCase[int]{
		{"empty", args[int]{NewLinkedList[int](), less, []int{}}},
		{"", args[int]{NewLinkedListItems[int](1), less, []int{1}}},
		{"", args[int]{NewLinkedListItems[int](3, 1), less, []int{1, 3}}},
		{"", args[int]{NewLinkedListItems[int](1, 5), less, []int{1, 5}}},
		{"", args[int]{NewLinkedListItems[int](1, 2, 3), less, []int{1, 2, 3}}},
		{"", args[int]{NewLinkedListItems[int](1, 3, 2), less, []int{1, 2, 3}}},
		{"", args[int]{NewLinkedListItems[int](2, 1, 3), less, []int{1, 2, 3}}},
		{"", args[int]{NewLinkedListItems[int](2, 3, 1), less, []int{1, 2, 3}}},
		{"", args[int]{NewLinkedListItems[int](3, 1, 2), less, []int{1, 2, 3}}},
		{"", args[int]{NewLinkedListItems[int](3, 2, 1), less, []int{1, 2, 3}}},
		{"", args[int]{NewLinkedListItems[int](4, 1, 2, 3, 5), less, []int{1, 2, 3, 4, 5}}},
		{"", args[int]{NewLinkedListItems[int](5, 1, 4, 2, 3), less, []int{1, 2, 3, 4, 5}}},
		{"", args[int]{NewLinkedListItems[int](2, 5, 4, 1, 3), less, []int{1, 2, 3, 4, 5}}},
		{"", args[int]{NewLinkedListItems[int](3, 5, 4, 2, 1), less, []int{1, 2, 3, 4, 5}}},
	}
	for _, tt := range tests {
		if tt.name == "" {
			tt.name = fmt.Sprint(tt.args.list.ToArray())
		}
		t.Run(tt.name, func(t *testing.T) {
			sortItems(tt.args.list.first, tt.args.list.last, tt.args.less)
			actual := tt.args.list.ToArray()
			if !reflect.DeepEqual(actual, tt.args.want) {
				t.Errorf("sortItems() got: %v, want: %v", actual, tt.args.want)
			}
		})
	}
}

func Test_sortItems_string_big(t *testing.T) {
	expected := []string{"str01", "str02", "str03", "str04", "str05", "str06", "str07", "str08", "str09"}
	shuffle := circleLeftShiftIterator(expected)
	less := func(val1, val2 string) bool { return val1 < val2 }
	count := 2 * 3 * 4 * 5 * 6 * 7 * 8 * 9
	t.Log("count:", count)
	for i := 0; i < count; i++ {
		source := shuffle()
		list := NewLinkedListItems[string](source...)
		SortList[string](list, less)
		actual := list.ToArray()
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("sortItems() got: %v, want: %v", actual, expected)
		}
	}
}
func Test_sortItems_int_big(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	shuffle := circleLeftShiftIterator(expected)
	less := func(val1, val2 int) bool { return val1 < val2 }
	count := 2 * 3 * 4 * 5 * 6 * 7 * 8 * 9 * 10
	t.Log("count:", count)
	for i := 0; i < count; i++ {
		source := shuffle()
		list := NewLinkedListItems[int](source...)
		SortList(list, less)
		actual := list.ToArray()
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("sortItems() got: %v, want: %v", actual, expected)
		}
	}

}
func Test_sortItems_int_big_revers(t *testing.T) {
	expected := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	shuffle := circleLeftShiftIterator(expected)
	less := func(val1, val2 int) bool { return val1 > val2 }
	count := 2 * 3 * 4 * 5 * 6 * 7 * 8 * 9 * 10
	t.Log("count:", count)
	for i := 0; i < count; i++ {
		source := shuffle()
		list := NewLinkedListItems[int](source...)
		sortItems(list.first, list.last, less)
		actual := list.ToArray()
		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("sortItems() got: %v, want: %v", actual, expected)
		}
	}

}

func Test_circleLeftShift_string(t *testing.T) {
	tests := []struct {
		source []string
	}{
		{[]string{"string 1"}},
		{[]string{"string 1", "string 2"}},
		{[]string{"string 1", "string 2", "string 3"}},
		{[]string{"string 1", "string 2", "string 3", "string 4"}},
		{[]string{"string 1", "string 2", "string 3", "string 4", "string 5"}},
	}
	factorial := func(n int) int {
		res := 1
		for i := 2; i <= n; i++ {
			res *= i
		}
		return res
	}
	for _, tt := range tests {
		count := factorial(len(tt.source))
		t.Run(fmt.Sprintf("%v %v", count, tt.source), func(t *testing.T) {
			arIter := circleLeftShiftIterator(tt.source)
			exists := make([][]string, 0, count)
			contains := func(array []string) bool {
				for _, a := range exists {
					if reflect.DeepEqual(a, array) {
						return true
					}
				}
				return false
			}
			for i := 0; i < count; i++ {
				got := arIter()
				if contains(got) {
					t.Fatal("duplicate array:", got)
				}
				exists = append(exists, got)
				t.Log(got)
			}
			if len(exists) != count {
				t.Fatalf("incorrect array length: %d, want: %d", len(exists), count)
			}
		})
	}

}
func Test_circleLeftShift_int(t *testing.T) {
	factorial := func(n int) int {
		res := 1
		for i := 2; i <= n; i++ {
			res *= i
		}
		return res
	}
	tests := []struct {
		source []int
	}{
		{[]int{1}},
		{[]int{1, 2}},
		{[]int{1, 2, 3}},
		{[]int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}},
	}

	for _, tt := range tests {
		count := factorial(len(tt.source))
		t.Run(fmt.Sprintf("%v %v", count, tt.source), func(t *testing.T) {
			arIter := circleLeftShiftIterator(tt.source)
			exists := make([][]int, 0, count)
			contains := func(array []int) bool {
				for _, a := range exists {
					if reflect.DeepEqual(a, array) {
						return true
					}
				}
				return false
			}
			for i := 0; i < count; i++ {
				got := arIter()
				if contains(got) {
					t.Fatal("duplicate array:", got)
				}
				exists = append(exists, got)
				t.Log(got)
			}
			if len(exists) != count {
				t.Fatalf("incorrect array length: %d, want: %d", len(exists), count)
			}
		})
	}
}
func Test_circleLeftShift(t *testing.T) {
	ar := []int{1, 2, 3}
	arIter := circleLeftShiftIterator(ar)
	want1 := []int{1, 2, 3}
	want2 := []int{2, 3, 1}
	want3 := []int{3, 1, 2}
	want4 := []int{2, 1, 3}
	want5 := []int{1, 3, 2}
	want6 := []int{3, 2, 1}
	got1 := arIter()
	got2 := arIter()
	got3 := arIter()
	got4 := arIter()
	got5 := arIter()
	got6 := arIter()

	if !reflect.DeepEqual(got1, want1) {
		t.Fatal("want:", want1, "got:", got1)
	}
	if !reflect.DeepEqual(got2, want2) {
		t.Fatal("want:", want2, "got:", got2)
	}
	if !reflect.DeepEqual(got3, want3) {
		t.Fatal("want:", want3, "got:", got3)
	}
	if !reflect.DeepEqual(got4, want4) {
		t.Fatal("want:", want4, "got:", got4)
	}
	if !reflect.DeepEqual(got5, want5) {
		t.Fatal("want:", want5, "got:", got5)
	}
	if !reflect.DeepEqual(got6, want6) {
		t.Fatal("want:", want6, "got:", got6)
	}
	got7 := arIter()
	if !reflect.DeepEqual(got7, want1) {
		t.Fatal("want:", want1, "got:", got7)
	}
	got8 := arIter()
	if !reflect.DeepEqual(got8, want2) {
		t.Fatal("want:", want2, "got:", got8)
	}
}
