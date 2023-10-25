package lists

import (
	"fmt"
	"reflect"
	"testing"
)

func BenchmarkSortList_int(b *testing.B) {
	benchmarks := []struct {
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}
	list := NewLinkedList[int]()
	fillList := func(array []int) {
		for _, val := range array {
			list.AddLast(val)
		}
	}
	less := func(v1, v2 int) bool { return v1 < v2 }
	for _, bm := range benchmarks {
		b.Run(fmt.Sprint(bm.want), func(b *testing.B) {
			iterator := circleLeftShiftIterator(bm.want)
			b.ResetTimer()
			b.StopTimer()
			for i := 0; i < b.N; i++ {
				list.Clear()
				fillList(iterator())
				b.StartTimer()

				SortList(list, less)

				b.StopTimer()
				actual := list.ToArray()
				if !reflect.DeepEqual(actual, bm.want) {
					b.Fatalf("SortList() %d. got: %v; want: %v", i, actual, bm.want)
				}
			}
		})
	}
}
