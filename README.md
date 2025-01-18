# Learn Go Through Python Comparisons

A comprehensive tutorial for learning Go, with Python comparisons for better understanding.

By Reza Shahnazar ([@rezashahnazar](https://github.com/rezashahnazar))

## Table of Contents

1. [Introduction](#introduction)
2. [Project Setup](#project-setup)
3. [Basic Syntax and Types](#basic-syntax-and-types)
4. [Object-Oriented Programming](#object-oriented-programming)
5. [Interfaces and Polymorphism](#interfaces-and-polymorphism)
6. [Error Handling](#error-handling)
7. [Concurrency](#concurrency)
8. [Project Structure](#project-structure)
9. [Testing](#testing)
10. [Best Practices](#best-practices)

## Introduction

This tutorial compares Go and Python implementations to help you understand Go's concepts through familiar Python patterns. We'll use a bookstore management system as our learning example.

## Project Setup

### Python Setup

```python
# Create virtual environment
python -m venv venv
source venv/bin/activate  # On Windows: venv\Scripts\activate
pip install mypy  # For type checking
```

### Go Setup

```bash
# Initialize a new Go module
go mod init learn-golang
```

## Basic Syntax and Types

### Variables and Constants

Python:

```python
# Python variable declaration
title = "The Go Programming Language"
price = 29.99
in_stock = True

# Python constants (by convention)
MAX_BOOKS = 1000
CATEGORY_CODE = "BOOK"
```

Go:

```go
package main

// Go variable declaration
var title string = "The Go Programming Language"
price := 29.99  // Type inference
var inStock bool = true

// Go constants
const (
    MaxBooks     = 1000
    CategoryCode = "BOOK"
)
```

### Type System

Let's examine the Python example's type hints (lines 38-70 in example_book.py) and create a Go equivalent:

```go
// book.go
package main

type Book struct {
    title     string
    author    string
    price     float64
    pageCount int
    seller    string
}

// Constructor equivalent
func NewBook(title, author string, price float64, seller string) *Book {
    return &Book{
        title:     title,
        author:    author,
        price:     price,
        pageCount: randomPageCount(),
        seller:    seller,
    }
}
```

## Object-Oriented Programming

### Classes vs Structs

Python's class-based approach (from example_book.py):

```python
class Book:
    def __init__(self, title, author, price, seller=None):
        self._title = title
        self._author = author
        self._price = price
        self.seller = seller
```

Go's struct-based approach:

```go
// book.go
type Book struct {
    title  string  // Private (lowercase)
    author string
    price  float64
    Seller string  // Public (uppercase)
}

// Methods are defined outside the struct
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

### Properties and Methods

Python's property decorators (from example_book.py, lines 118-148) vs Go's method approach:

```go
// book.go
type Book struct {
    pageCount int
}

// Getter
func (b *Book) PageCount() int {
    return b.pageCount
}

// Setter
func (b *Book) SetPageCount(count int) {
    b.pageCount = count
}
```

## Interfaces and Polymorphism

Python's abstract base class (from example_book.py, lines 5-35) vs Go's interface:

```go
// priced_item.go
type PricedItem interface {
    GetPrice() float64
    SetPrice(price float64) error
    CalculateDiscount(percentage float64) (float64, error)
}

// book.go
func (b *Book) CalculateDiscount(percentage float64) (float64, error) {
    if percentage < 0 || percentage > 100 {
        return 0, fmt.Errorf("percentage must be between 0 and 100")
    }
    return b.price * (1 - percentage/100), nil
}

// magazine.go
type Magazine struct {
    name        string
    price       float64
    issueNumber int
}

func (m *Magazine) CalculateDiscount(percentage float64) (float64, error) {
    if percentage < 0 || percentage > 100 {
        return 0, fmt.Errorf("percentage must be between 0 and 100")
    }
    baseDiscount := m.price * (1 - percentage/100)
    if m.price > 10 {
        return baseDiscount * 0.9, nil
    }
    return baseDiscount, nil
}
```

## Error Handling

Python's exception handling vs Go's explicit error handling:

Python:

```python
try:
    book.set_price(-10)
except ValueError as e:
    print(f"Error: {e}")
```

Go:

```go
if err := book.SetPrice(-10); err != nil {
    log.Printf("Error: %v", err)
    return err
}
```

## Concurrency

Python's async/await vs Go's goroutines:

```python
# Python async
import asyncio

async def fetch_book_price(book_id):
    await asyncio.sleep(1)  # Simulate API call
    return 29.99

async def main():
    prices = await asyncio.gather(
        fetch_book_price(1),
        fetch_book_price(2)
    )
```

Go:

```go
func fetchBookPrice(bookID int, prices chan float64) {
    time.Sleep(time.Second) // Simulate API call
    prices <- 29.99
}

func main() {
    prices := make(chan float64, 2)

    go fetchBookPrice(1, prices)
    go fetchBookPrice(2, prices)

    price1 := <-prices
    price2 := <-prices
}
```

## Project Structure

### Python Project Structure

```
project/
├── venv/
├── src/
│   ├── __init__.py
│   ├── models/
│   │   ├── __init__.py
│   │   ├── book.py
│   │   └── magazine.py
│   └── services/
│       ├── __init__.py
│       └── pricing.py
├── tests/
│   └── test_book.py
└── requirements.txt
```

### Go Project Structure

```
project/
├── cmd/
│   └── bookstore/
│       └── main.go
├── internal/
│   ├── book/
│   │   └── book.go
│   └── magazine/
│       └── magazine.go
├── pkg/
│   └── pricing/
│       └── pricing.go
├── go.mod
└── go.sum
```

## Testing

Python test:

```python
import unittest

class TestBook(unittest.TestCase):
    def test_discount_calculation(self):
        book = Book("Test", "Author", 100.0)
        self.assertEqual(book.calculate_discount(20), 80.0)
```

Go test:

```go
package book

import "testing"

func TestDiscountCalculation(t *testing.T) {
    book := NewBook("Test", "Author", 100.0, "")
    discount, err := book.CalculateDiscount(20)

    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    if discount != 80.0 {
        t.Errorf("Expected 80.0, got %f", discount)
    }
}
```

## Best Practices

### Go Best Practices

1. Use error handling instead of exceptions
2. Follow Go's official style guide (gofmt)
3. Use interfaces for flexibility
4. Implement concurrency with goroutines and channels
5. Use meaningful package names
6. Keep interfaces small
7. Use composition over inheritance

### Python vs Go Conventions

1. Naming:
   - Python: snake_case
   - Go: camelCase
2. Visibility:
   - Python: \_private (convention)
   - Go: uppercase/lowercase (enforced)
3. Error Handling:
   - Python: try/except
   - Go: explicit error returns
4. Documentation:
   - Python: docstrings
   - Go: godoc comments

## Exercises

1. Implement a `Library` struct that manages a collection of `PricedItem`s
2. Add concurrent price updates using goroutines
3. Implement a REST API using Go's standard `net/http` package
4. Write comprehensive tests for the `Book` and `Magazine` types
5. Add database integration using Go's `database/sql` package

## Additional Resources

1. [Official Go Documentation](https://golang.org/doc/)
2. [Go by Example](https://gobyexample.com/)
3. [Effective Go](https://golang.org/doc/effective_go)
4. [Go Playground](https://play.golang.org/)

## License

This tutorial is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

For questions or contributions, please contact:

- Email: reza.shahnazar@gmail.com
- GitHub: [@rezashahnazar](https://github.com/rezashahnazar)
