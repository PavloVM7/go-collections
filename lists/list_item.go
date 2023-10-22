package lists

import "fmt"

type listItem[T any] struct {
	prev  *listItem[T]
	next  *listItem[T]
	value T
}

func (li *listItem[T]) insert(item *listItem[T]) {
	item.prev = li.prev
	item.next = li
	li.prev = item
}
func (li *listItem[T]) append(item *listItem[T]) {
	item.prev = li
	item.next = li.next
	li.next = item
}
func (li *listItem[T]) remove() {
	if li.prev != nil {
		li.prev.next = li.next
	}
	if li.next != nil {
		li.next.prev = li.prev
	}
}
func (li *listItem[T]) String() string {
	vs := ""
	if s, ok := any(li.value).(string); ok {
		vs = "'" + s + "'"
	} else {
		vs = fmt.Sprint(li.value)
	}
	return fmt.Sprintf("listItem[%T]{value: %s, prev: %v, next: %v}", li.value, vs, li.prev, li.next)
}
