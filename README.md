# Go Collections

## Overview

`go-collections` is a Go package that provides a set of generic data structures inspired by the Java Collections Framework. This package includes implementations for commonly used data structures such as `LinkedList`, `Queue`, and `Stack`, with support for both linked-list and slice-based implementations.

## Features

- **LinkedList:** A generic linked list implementation supporting basic operations like add, remove, and contains.
- **Queue:** A generic queue implementation using both linked lists and slices.
- **Stack:** A generic stack implementation using both linked lists and slices.

## Inspiration

This package draws inspiration from the Java Collections Framework, aiming to provide similar functionality in Go while taking advantage of Go's type system and concurrency features.

## Installation

To install the package, run:

```bash
go get https://github.com/nidhey27/go-collections.git
```

### Usage
Here is an example of how to use the LinkedList, Queue, and Stack implementations:

#### LinkedList
```go
package main

import (
    "fmt"
    "github.com/yourusername/your_project/collections"
)

func main() {
    ll := collections.LinkedList[int]{}
    ll.Add(1)
    ll.Add(2)
    ll.Add(3)

    fmt.Println("Contains 2:", ll.Contains(2))
    ll.Remove(2)
    fmt.Println("Contains 2 after removal:", ll.Contains(2))
}
```

#### Queue
```go
package main

import (
    "fmt"
    "github.com/yourusername/your_project/collections"
)

func main() {
    queue := collections.NewSliceQueue[int]()
    queue.Enqueue(1)
    queue.Enqueue(2)
    queue.Enqueue(3)

    fmt.Println("Dequeue:", queue.Dequeue())
    fmt.Println("Peek:", queue.Peek())
}
```

#### Stack
```go
package main

import (
    "fmt"
    "github.com/yourusername/your_project/collections"
)

func main() {
    stack := collections.NewSliceStack[int]()
    stack.Push(1)
    stack.Push(2)
    stack.Push(3)

    fmt.Println("Pop:", stack.Pop())
    fmt.Println("Peek:", stack.Peek())
}
```

### Testing
To run tests for the package, use:
```bash
go test ./collections
```

### Contributing
Contributions are welcome! Please open an issue or submit a pull request if you have suggestions or improvements.