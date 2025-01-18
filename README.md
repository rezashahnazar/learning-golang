# Learn Go Through Python Comparisons

A comprehensive guide for Python developers learning Go, with practical examples and side-by-side comparisons.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Project Setup](#project-setup)
- [1. Basic Syntax Comparison](#1-basic-syntax-comparison)
- [2. Variables and Types](#2-variables-and-types)
- [3. Control Structures](#3-control-structures)
- [4. Functions and Methods](#4-functions-and-methods)
- [5. Interfaces and Types](#5-interfaces-and-types)
- [6. Object-Oriented Programming](#6-object-oriented-programming)
- [7. Error Handling](#7-error-handling)
- [8. Collections and Data Structures](#8-collections-and-data-structures)
- [9. Concurrency](#9-concurrency)
- [10. Project Examples](#10-project-examples)

## Prerequisites

- Go 1.23.5 or later
- Python 3.x (for comparison)
- Basic understanding of Python
- A code editor (VS Code recommended)
- Git for version control

## Project Setup

1. Initialize your Go module:

```bash
go mod init learn-golang
```

2. Create project structure:

```bash
mkdir src
touch main.go
touch src/book.go
touch src/magazine.go
```

## 1. Basic Syntax Comparison

### Python Hello World

```python
def main():
    print("Hello, World!")

if __name__ == "__main__":
    main()
```

### Go Hello World

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Key differences:

- Go requires explicit package declaration
- No need for `if __name__ == "__main__"`
- Functions use `func` keyword
- Semicolons (though usually optional)
- Static typing

## 2. Variables and Types

### Python Variables

```python
name = "John"
age = 30
height = 1.75
is_student = True

# Type hints (optional)
name: str = "John"
age: int = 30
```

### Go Variables

```go
var name string = "John"
age := 30  // Type inference
var height float64 = 1.75
var isStudent bool = true

// Constants
const MaxAge = 120
```

Key differences:

- Go is statically typed
- Type inference with `:=`
- Constants with `const`
- No implicit type conversion
- Boolean is `bool`, not `True/False`

## 3. Control Structures

### Python Control Flow

```python
# If statement
if age < 18:
    print("Minor")
elif age < 65:
    print("Adult")
else:
    print("Senior")

# For loop
for i in range(5):
    print(i)

# While loop
while count > 0:
    count -= 1
```

### Go Control Flow

```go
// If statement
if age < 18 {
    fmt.Println("Minor")
} else if age < 65 {
    fmt.Println("Adult")
} else {
    fmt.Println("Senior")
}

// For loop (only loop type in Go)
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// While-like loop
for count > 0 {
    count--
}
```

## 4. Functions and Methods

### Python Functions

```python
def add(a: int, b: int) -> int:
    return a + b

# Multiple returns
def divide_mod(a: int, b: int) -> tuple[int, int]:
    return a // b, a % b

# Method in class
class Calculator:
    def multiply(self, x: int, y: int) -> int:
        return x * y
```

### Go Functions

```go
func add(a, b int) int {
    return a + b
}

// Multiple returns
func divideMod(a, b int) (int, int) {
    return a / b, a % b
}

// Method on struct
type Calculator struct{}

func (c Calculator) Multiply(x, y int) int {
    return x * y
}
```

## 5. Interfaces and Types

### Python Abstract Base Class

```python
from abc import ABC, abstractmethod

class PricedItem(ABC):
    @abstractmethod
    def get_price(self) -> float:
        pass

    @abstractmethod
    def set_price(self, price: float) -> None:
        pass

    @abstractmethod
    def calculate_discount(self, percentage: float) -> float:
        pass
```

### Go Interface

```go
type PricedItem interface {
    GetPrice() float64
    SetPrice(price float64) error
    CalculateDiscount(percentage float64) (float64, error)
}
```

## 6. Object-Oriented Programming

### Python Class

```python
class Book:
    def __init__(self, title: str, author: str, price: float):
        self.title = title
        self.author = author
        self._price = price

    @property
    def price(self) -> float:
        return self._price

    @price.setter
    def price(self, value: float) -> None:
        if value < 0:
            raise ValueError("Price cannot be negative")
        self._price = value
```

### Go Struct

```go
type Book struct {
    title  string
    author string
    price  float64
}

func NewBook(title, author string, price float64) *Book {
    return &Book{
        title:  title,
        author: author,
        price:  price,
    }
}

func (b *Book) GetPrice() float64 {
    return b.price
}

func (b *Book) SetPrice(price float64) error {
    if price < 0 {
        return fmt.Errorf("price cannot be negative")
    }
    b.price = price
    return nil
}
```

## 7. Error Handling

### Python Error Handling

```python
try:
    book.price = -10
except ValueError as e:
    print(f"Error: {e}")
```

### Go Error Handling

```go
if err := book.SetPrice(-10); err != nil {
    fmt.Printf("Error: %v\n", err)
    return
}
```

## 8. Collections and Data Structures

### Python Collections

```python
# Lists
numbers = [1, 2, 3]
numbers.append(4)

# Dictionaries
person = {
    "name": "John",
    "age": 30
}

# Sets
unique_nums = {1, 2, 3}
```

### Go Collections

```go
// Slices
numbers := []int{1, 2, 3}
numbers = append(numbers, 4)

// Maps
person := map[string]interface{}{
    "name": "John",
    "age":  30,
}

// No built-in set type (use map[type]bool)
uniqueNums := map[int]bool{
    1: true,
    2: true,
    3: true,
}
```

## 9. Concurrency

### Python Threading/Async

```python
import asyncio

async def process_item(item):
    await asyncio.sleep(1)
    return item * 2

async def main():
    tasks = [process_item(i) for i in range(5)]
    results = await asyncio.gather(*tasks)
```

### Go Goroutines

```go
func processItem(item int, ch chan int) {
    time.Sleep(time.Second)
    ch <- item * 2
}

func main() {
    ch := make(chan int)
    for i := 0; i < 5; i++ {
        go processItem(i, ch)
    }

    for i := 0; i < 5; i++ {
        result := <-ch
        fmt.Println(result)
    }
}
```

## 10. Project Examples

### Book Management System

See the complete implementation in the repository:

- `src/book.go` - Book type implementation
- `src/magazine.go` - Magazine type implementation
- `main.go` - Main program demonstrating usage

## Running the Examples

### Python

```bash
python example_book.py
```

### Go

```bash
go run .
```

## Key Takeaways

1. Go is statically typed vs Python's dynamic typing
2. Go uses explicit error handling instead of exceptions
3. Interfaces are implicit in Go
4. Go promotes composition over inheritance
5. Public/private access is determined by capitalization
6. Go has built-in concurrency support with goroutines
7. Go uses pointers for efficiency
8. No classes in Go - structs and methods instead

## Next Steps

1. Implement more complex data structures
2. Practice error handling patterns
3. Explore Go's concurrency features
4. Build a REST API
5. Learn about Go modules and dependency management

## Common Gotchas for Python Developers

1. No exceptions for flow control
2. No operator overloading
3. No optional parameters (use structs for options)
4. No inheritance (use composition)
5. Must handle all errors explicitly
6. No globals or dynamic imports
7. Package names must match directory structure

## Resources

- [Official Go Documentation](https://golang.org/doc/)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
- [Go Playground](https://play.golang.org/)
- [Go Tour](https://tour.golang.org/)

## Contributing

Feel free to submit issues and enhancement requests!

Contact: Reza Shahnazar (reza.shahnazar@gmail.com)
GitHub: [@rezashahnazar](https://github.com/rezashahnazar)

## License

This project is licensed under the MIT License - see the LICENSE file for details.
