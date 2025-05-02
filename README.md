## Code Explanation

- **package main**: Defines the program as an executable package, producing a standalone binary.
- **import "fmt"**: Imports the `fmt` package for formatting and printing text to the console.
- **main function**: The starting point of the program, executed when the program runs.

# My First Go Program

Welcome to my first Go program! This simple application prints a greeting message to the console using the Go programming language.

## Overview

This program demonstrates the basic structure of a Go program, including the `main` package, the `fmt` package for output, and the `main` function as the entry point.

## Code

```go
package main

import "fmt"

func main() {
    fmt.Println("Hi there!")
}


In Go (Golang), data types are categorized into primitive and non-primitive types. Let me explain both with examples:

## Primitive Types in Go

Primitive types are basic data types built into the language. They're stored directly in memory and operated on directly.

### Numeric Types
```go
// Integers
var age int = 30
var count int8 = 127     // -128 to 127
var population uint64 = 7800000000

// Floating point
var height float32 = 5.9
var weight float64 = 70.5

// Complex numbers
var complex1 complex64 = 1 + 2i
var complex2 complex128 = 3.14 + 2.71i
```

### Boolean Type
```go
var isValid bool = true
var isCompleted bool = false
```

### String Type
```go
var name string = "John"
var message string = `This is a 
multi-line string`
```

### Other Primitive Types
```go
// Byte (alias for uint8)
var b byte = 'A'  // 65

// Rune (alias for int32, represents Unicode code point)
var r rune = '世'  // 19990
```

## Non-Primitive Types in Go

Non-primitive types are composite types constructed from primitive types.

### Arrays
```go
// Fixed size collection of elements
var numbers [5]int = [5]int{1, 2, 3, 4, 5}
var matrix [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
```

### Slices
```go
// Dynamic size collection (reference to arrays)
var fruits []string = []string{"apple", "banana", "orange"}
numbers := make([]int, 5, 10)  // length 5, capacity 10
```

### Maps
```go
// Key-value pairs
var studentGrades map[string]int = map[string]int{
    "Alice": 95,
    "Bob":   87,
    "Carol": 92,
}
```

### Structs
```go
// Custom data type with named fields
type Person struct {
    Name    string
    Age     int
    Address string
}

var employee Person = Person{
    Name:    "Alice",
    Age:     30,
    Address: "123 Main St",
}
```

### Pointers
```go
// Stores memory address of another variable
var x int = 10
var ptr *int = &x  // ptr holds the memory address of x
```

### Functions
```go
// Functions can be assigned to variables
var add func(int, int) int = func(a, b int) int {
    return a + b
}
result := add(5, 3)  // 8
```

### Interfaces
```go
// Defines behavior
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Circle implements Shape interface
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14 * c.Radius
}
```

### Channels
```go
// Used for communication between goroutines
ch := make(chan int)
go func() {
    ch <- 42  // Send value to channel
}()
value := <-ch  // Receive value from channel
```

Go's type system is designed for clarity and efficiency, with a focus on compile-time type safety.