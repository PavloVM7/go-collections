[![Go](https://github.com/PavloVM7/go-collections/actions/workflows/go.yml/badge.svg)](https://github.com/PavloVM7/go-collections/actions/workflows/go.yml)
# go-collections
This module contains some NOT(!) thread safe collections

## LinkedList

`LinkedList` is an implementation of a doubly-linked list.

### Usage

```go
package main

import "github.com/PavloVM7/go-collections/lists"

list := NewLinkedList[int]()
list.AddLast(2) // Adds 2 to the end of the list
list.AddLast(3)  // Add 3 to the end of the list
list.AddFirst(1) // Insert 1 at the beginning of the list (before 2)

```

## ⌨️ Author

[@PavloVM7](https://github.com/PavloVM7) - Idea & Initial work
