# Go Programming Tutorial: From Python to Go

This tutorial demonstrates the transition from Python to Go programming, highlighting key differences and similarities between the two languages.

## Table of Contents

- [Installation](#installation)
- [Project Setup](#project-setup)
- [Basic Concepts](#basic-concepts)
- [Structs vs Classes](#structs-vs-classes)
- [Methods and Functions](#methods-and-functions)
- [Error Handling](#error-handling)
- [Memory Management](#memory-management)

## Installation

1. Download Go from [golang.org](https://golang.org/dl/)
2. Verify installation:

```bash
go version
```

## Project Setup

1. Create a new Go module:

```bash
go mod init learn-golang
```

2. Your project structure should look like:

```
learn-golang/
├── go.mod
├── main.go
└── README.md
```

## Basic Concepts

### Key Differences from Python

1. **Static Typing**: Go requires explicit type declarations
2. **Package System**: Every Go file must belong to a package
3. **No Classes**: Go uses structs and methods instead
4. **Error Handling**: Go uses explicit error returns instead of exceptions
5. **Exported Names**: Capitalization determines visibility

### Example Comparison

Let's look at how a simple Book implementation differs between Python and Go:

Python Version (reference from example_book.py):

```python:example_book.py
startLine: 4
endLine: 20
```

Go Version (reference from main.go):

```go:main.go
startLine: 24
endLine: 31
```

## Structs vs Classes

### Defining Types

In Go, we use structs instead of classes. Here's how to create a new Book:

```go
type Book struct {
    title     string
    author    string
    price     float64
    pageCount int
}
```

### Constructors

Python uses `__init__`, while Go uses constructor functions:

```go
func NewBook(title string, author string, price float64) *Book {
    return &Book{
        title:     title,
        author:    author,
        price:     price,
        pageCount: rand.Intn(1000), // Example random page count
    }
}
```

## Methods and Functions

### Method Definition

In Go, methods are defined with a receiver parameter:

```go
func (b *Book) Summary() string {
    return fmt.Sprintf("%s by %s - $%.2f", b.title, b.author, b.price)
}
```

### Getters and Setters

While Python uses properties, Go uses explicit methods:

```go
// Getter
func (b *Book) GetPrice() float64 {
    return b.price
}

// Setter
func (b *Book) SetPrice(price float64) error {
    if price < 0 {
        return fmt.Errorf("price cannot be negative")
    }
    b.price = price
    return nil
}
```

## Error Handling

Go uses explicit error returns instead of exceptions:

```go
func (b *Book) SetPrice(price float64) error {
    if price < 0 {
        return fmt.Errorf("price cannot be negative")
    }
    b.price = price
    return nil
}

// Usage
if err := book.SetPrice(-10); err != nil {
    fmt.Println("Error:", err)
}
```

## Memory Management

- Go uses garbage collection
- Pointers are common but simpler than in C/C++
- The `&` operator creates a pointer
- The `*` operator dereferences a pointer

## Running the Code

Execute your Go program:

```bash
go run main.go
```

## Best Practices

1. Use meaningful variable names
2. Follow Go's official style guide
3. Use proper error handling
4. Capitalize exported names
5. Use comments for package documentation

## Resources

- [Official Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go)

## Author

Reza Shahnazar

- GitHub: [@rezashahnazar](https://github.com/rezashahnazar)
- Email: reza.shahnazar@gmail.com
