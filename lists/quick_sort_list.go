package lists

func SortList[T any](list *LinkedList[T], less func(item1, item2 T) bool) {
	sortItems[T](list.first, list.last, less)
}

func sortItems[T any](start, end *listItem[T], less func(item1, item2 T) bool) {
	if start == nil || start == end || start == end.next || start.next == nil {
		return
	}
	current := start
	replace := start
	for current != end {
		if less(current.value, end.value) {
			swapListItems[T](replace, current)
			replace = replace.next
		}
		current = current.next
	}
	swapListItems[T](end, replace)
	if replace.prev != nil && start != replace.prev {
		sortItems[T](start, replace.prev, less)
	}
	if replace.next != nil && replace.next.next != nil && replace != end && replace.next != end {
		sortItems[T](replace.next, end, less)
	}
}

func circleLeftShiftIterator[T any](ar []T) func() []T {
	cpy := make([]T, len(ar))
	copy(cpy, ar)
	indexes := make([]int, len(ar))
	for i := 0; i < len(ar); i++ {
		indexes[i] = i
	}
	retArray := func() []T {
		result := make([]T, len(indexes))
		for i := 0; i < len(result); i++ {
			result[i] = cpy[indexes[i]]
		}
		return result
	}

	k := -1
	last := len(indexes) - 1
	left := func() {
		i0 := indexes[0]
		for i := 0; i < k; i++ {
			indexes[i] = indexes[i+1]
		}
		indexes[k] = i0
	}
	return func() []T {
		if k != -1 {
			for {
				left()
				if indexes[k] != k {
					k = last
					break
				} else {
					k--
					if k < 0 {
						k = last
						return retArray()
					}
				}
			}
		} else {
			k = last
		}
		return retArray()
	}
}
