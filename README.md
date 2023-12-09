[![Go](https://github.com/PavloVM7/go-collections/actions/workflows/go.yml/badge.svg)](https://github.com/PavloVM7/go-collections/actions/workflows/go.yml)
# go-collections
This module contains some NOT(!) thread safe collections

### Usage

```
go get github.com/PavloVM7/go-collections@v0.1.0
```

## LinkedList

`LinkedList` is an implementation of a doubly-linked list.

### Usage

```go
package main

import (
	"fmt"
	"github.com/PavloVM7/go-collections/pkg/collections/lists"
)

func main() {
	list := lists.NewLinkedList[int]()

	using := func(funcs string) {
		fmt.Printf("=== using %s\n", funcs)
	}
	showList := func() {
		fmt.Printf(">>> list size: %d, items: %v\n", list.Size(), list.ToArray())
	}

	using("AddFirst()")
	for i := 10; i > 0; i-- {
		list.AddFirst(i)
	}
	showList()

	using("AddLast()")
	for i := 11; i <= 20; i++ {
		list.AddLast(i)
	}
	showList()

	using("Get() and Remove()")
	item10, err := list.Get(10)
	fmt.Printf("before remove 10th item = %d, err = %v\n", item10, err)
	item10, err = list.Remove(10) // removes 10th item
	fmt.Printf("removed item10 = %d, err = %v\n", item10, err)
	item10, err = list.Get(10)
	fmt.Printf("after remove 10th item = %d, err = %v\n", item10, err)
	showList()

	using("GetFirst() and RemoveFirst()")
	first, firstOk := list.GetFirst()
	fmt.Printf("before remove first element: %d, exists: %t\n", first, firstOk)
	first, firstOk = list.RemoveFirst()
	fmt.Printf("first element: %d, removed: %t\n", first, firstOk)
	first, firstOk = list.GetFirst()
	fmt.Printf("current first element: %d, exists: %t\n", first, firstOk)
	showList()

	using("GetLast() and RemoveLast()")
	last, lastOk := list.GetLast()
	fmt.Printf("before remove last element: %d, exists: %t\n", last, lastOk)
	last, lastOk = list.RemoveLast()
	fmt.Printf("last element: %d, removed: %t\n", last, lastOk)
	last, lastOk = list.GetLast()
	fmt.Printf("current last element: %d, exists: %t\n", last, lastOk)
	showList()

	using("RemoveFirstOccurrence()")
	rFirst, fIndex := list.RemoveFirstOccurrence(func(value int) bool {
		return value%2 != 0
	})
	fmt.Printf("removed first odd value: %d, index: %d\n", rFirst, fIndex)
	showList()

	using("RemoveLastOccurrence()")
	rLast, lIndex := list.RemoveLastOccurrence(func(value int) bool {
		return value%2 == 0
	})
	fmt.Printf("removed last even value: %d, index: %d\n", rLast, lIndex)
	showList()

	using("RemoveAll()")
	count := list.RemoveAll(func(value int) bool {
		return value%3 == 0
	})
	fmt.Printf("%d elements that are dividable by 3 have been removed\n", count)
	showList()

	using("Clear()")
	list.Clear()
	showList()
}
```

outputs:

```
=== using AddFirst()
>>> list size: 10, items: [1 2 3 4 5 6 7 8 9 10]
=== using AddLast()
>>> list size: 20, items: [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20]
=== using Get() and Remove()
before remove 10th item = 11, err = <nil>
removed item10 = 11, err = <nil>
after remove 10th item = 12, err = <nil>
>>> list size: 19, items: [1 2 3 4 5 6 7 8 9 10 12 13 14 15 16 17 18 19 20]
=== using GetFirst() and RemoveFirst()
before remove first element: 1, exists: true
first element: 1, removed: true
current first element: 2, exists: true
>>> list size: 18, items: [2 3 4 5 6 7 8 9 10 12 13 14 15 16 17 18 19 20]
=== using GetLast() and RemoveLast()
before remove last element: 20, exists: true
last element: 20, removed: true
current last element: 19, exists: true
>>> list size: 17, items: [2 3 4 5 6 7 8 9 10 12 13 14 15 16 17 18 19]
=== using RemoveFirstOccurrence()
removed first odd value: 3, index: 1
>>> list size: 16, items: [2 4 5 6 7 8 9 10 12 13 14 15 16 17 18 19]
=== using RemoveLastOccurrence()
removed last even value: 18, index: 14
>>> list size: 15, items: [2 4 5 6 7 8 9 10 12 13 14 15 16 17 19]
=== using RemoveAll()
4 elements that are dividable by 3 have been removed
>>> list size: 11, items: [2 4 5 7 8 10 13 14 16 17 19]
=== using Clear()
>>> list size: 0, items: []
```

#### Sort linked list

```go
package main

import (
	"fmt"
	"github.com/PavloVM7/go-collections/pkg/collections/lists"
)

func main() {
	list := lists.NewLinkedListItems[int](10, 8, 6, 4, 2, 1, 3, 5, 7, 9)
	fmt.Printf("before sorting the list: %v\n", list.ToArray())
	lists.SortList(list, func(item1, item2 int) bool {
		return item1 < item2
	})
	fmt.Printf("after sorting the list:  %v\n", list.ToArray())
}
```

outputs:

```text
before sorting the list: [10 8 6 4 2 1 3 5 7 9]
after sorting the list:  [1 2 3 4 5 6 7 8 9 10]
```

## Set

`Set` is a collection that does not contain duplicate elements.

### Usage

```go
package main

import (
	"fmt"
	"github.com/PavloVM7/go-collections/pkg/collections"
	"runtime"
)

func main() {
	set := collections.NewSetCapacity[int](3)
	using := func(funcs string) {
		fmt.Printf("=== using %s\n", funcs)
	}
	showSet := func() {
		fmt.Printf(">>> set capacity: %d, size: %d, elements: %v\n", set.Capacity(), set.Size(), set.ToSlice())
	}
	isSetEmpty := func() {
		fmt.Println("~~~ is set empty? -", set.IsEmpty())
	}
	showSet()
	isSetEmpty()
	using("AddAll()")
	if set.AddAll(1, 2, 3) {
		showSet()
		isSetEmpty()
	}
	if !set.AddAll(2, 3) {
		fmt.Println("- the set already contains values 2 and 3")
	}
	using("Add()")
	if set.Add(4) {
		fmt.Println("+ ", 4, "was added to the set")
		showSet()
	}
	if !set.Add(1) {
		fmt.Println("- the set already contains the value 1")
	}

	using("Contains()")
	showSet()
	if set.Contains(3) {
		fmt.Println("+ the set contains the value 3")
	}
	if set.Contains(4) {
		fmt.Println("+ the set contains the value 4")
	}
	if !set.Contains(123) {
		fmt.Println("- there is no value 123 in the set")
	}

	using("Remove()")
	if set.Remove(3) {
		fmt.Printf("+ the value %d was removed from the set\n", 3)
	}
	if set.Remove(4) {
		fmt.Printf("+ the value %d was removed from the set\n", 4)
	}
	if !set.Remove(123) {
		fmt.Printf("- the value %d was not removed from the set because the set did not contain it\n", 123)
	}
	showSet()

	using("Clear()")
	set.Clear()
	showSet()
	isSetEmpty()

	using("TrimToSize()")
	const number = 1_000_000
	for i := 1; i <= number; i++ {
		set.Add(i)
	}

	getMemStats := func() runtime.MemStats {
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		return mem
	}

	memToString := func(mem runtime.MemStats) string { return fmt.Sprintf("%d Kb", mem.Alloc/1024) }

	runtime.GC()

	fmt.Printf(">>> set capacity: %d, size: %d, memory usage: %s\n", set.Capacity(), set.Size(), memToString(getMemStats()))
	for i := 21; i <= number; i++ {
		set.Remove(i)
	}

	runtime.GC()

	fmt.Printf("after removing memory usage: %s, set size: %d\n", memToString(getMemStats()), set.Size())
	showSet()

	set.TrimToSize()

	runtime.GC()

	fmt.Printf("after TrimToSize() memory usage: %s, set size: %d\n", memToString(getMemStats()), set.Size())
	showSet()
}
```

outputs like this:

```text
>>> set capacity: 3, size: 0, elements: []
~~~ is set empty? - true
=== using AddAll()
>>> set capacity: 3, size: 3, elements: [1 2 3]
~~~ is set empty? - false
- the set already contains values 2 and 3
=== using Add()
+  4 was added to the set
>>> set capacity: 3, size: 4, elements: [1 2 3 4]
- the set already contains the value 1
=== using Contains()
>>> set capacity: 3, size: 4, elements: [3 4 1 2]
+ the set contains the value 3
+ the set contains the value 4
- there is no value 123 in the set
=== using Remove()
+ the value 3 was removed from the set
+ the value 4 was removed from the set
- the value 123 was not removed from the set because the set did not contain it
>>> set capacity: 3, size: 2, elements: [1 2]
=== using Clear()
>>> set capacity: 3, size: 0, elements: []
~~~ is set empty? - true
=== using TrimToSize()
>>> set capacity: 3, size: 1000000, memory usage: 21894 Kb
after removing memory usage: 21897 Kb, set size: 20
>>> set capacity: 3, size: 20, elements: [11 17 3 12 7 10 15 6 4 8 2 5 18 20 9 1 14 19 13 16]
after TrimToSize() memory usage: 97 Kb, set size: 20
>>> set capacity: 3, size: 20, elements: [5 19 15 6 18 9 1 17 7 2 13 16 11 3 10 8 20 14 12 4]

```

## Collections Utils

### Usage `CopyMap`

```go
package main

import (
	"fmt"
	"github.com/PavloVM7/go-collections/pkg/collections"
	"sort"
)

func main() {
	src := map[int]string{1: "value 1", 2: "value 2", 3: "value 3"}
	fmt.Println("source map:")
	showIntMap(src)

	cpy := collections.CopyMap(src)
	fmt.Println("copy of the map:")
	showIntMap(cpy)

	src[4] = "value 4"
	fmt.Println("source map after adding new value:")
	showIntMap(src)

	fmt.Println("copy of the map:")
	showIntMap(cpy)
}
func showIntMap(mp map[int]string) {
	keys := make([]int, 0, len(mp))
	for k := range mp {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Printf("\t%d => %v", k, mp[k])
	}
	fmt.Println()
}
```

outputs:

```text
source map:
        1 => value 1    2 => value 2    3 => value 3
copy of the map:
        1 => value 1    2 => value 2    3 => value 3
source map after adding new value:
        1 => value 1    2 => value 2    3 => value 3    4 => value 4
copy of the map:
        1 => value 1    2 => value 2    3 => value 3

```
## ⌨️ Author

[@PavloVM7](https://github.com/PavloVM7) - Idea & Initial work
