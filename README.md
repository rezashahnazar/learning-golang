# Go vs Python Tutorial: Learning Go for Python Developers

This tutorial helps Python developers learn Go by comparing common patterns and implementations between the two languages.

## Table of Contents

- [Basic Setup](#basic-setup)
- [Type System](#type-system)
- [Classes vs Structs](#classes-vs-structs)
- [Constructors](#constructors)
- [Methods and Receivers](#methods-and-receivers)
- [Error Handling](#error-handling)
- [Constants and Static Members](#constants-and-static-members)

## Basic Setup

### Python Setup

```python
# Create a virtual environment
python -m venv venv
source venv/bin/activate  # On Windows: venv\Scripts\activate
```

### Go Setup

```bash
# Initialize a new Go module
go mod init your-project-name
```

## Type System

Go is statically typed, while Python is dynamically typed. Let's see the difference:

### Python

```python
# Dynamic typing
x = 42        # type is inferred
x = "hello"   # can change type
```

### Go

```go
// Static typing
var x int = 42
// x = "hello"  // This would cause a compilation error
```

## Classes vs Structs

Python uses classes while Go uses structs. Let's compare the basic structure:

### Python Implementation

```python
class Book:
    def __init__(self, title, author, price):
        self._title = title
        self._author = author
        self._price = price
```

### Go Implementation

```go
type Book struct {
    title  string
    author string
    price  float64
}
```

## Constructors

Python uses `__init__` while Go uses constructor functions by convention:

### Python Constructor

```python
def __init__(self, title, author, price):
    self._title = title
    self._author = author
    self._price = price
    self._page_count = self.random_page_count()
```

### Go Constructor

```go
func NewBook(title string, author string, price float64) *Book {
    return &Book{
        title:     title,
        author:    author,
        price:     price,
        pageCount: randomPageCount(),
    }
}
```

## Methods and Receivers

### Python Methods

```python
def summary(self):
    return f"{self._title} by {self._author} - ${self._price:.2f}"

@property
def price(self):
    return self._price

@price.setter
def price(self, value):
    if value < 0:
        raise ValueError("price cannot be negative")
    self._price = value
```

### Go Methods

```go
func (b *Book) Summary() string {
    return fmt.Sprintf("%s by %s - $%.2f", b.title, b.author, b.price)
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

## Error Handling

### Python Error Handling

```python
def set_price(self, price):
    if price < 0:
        raise ValueError("price cannot be negative")
    self._price = price
```

### Go Error Handling

```go
func (b *Book) SetPrice(price float64) error {
    if price < 0 {
        return fmt.Errorf("price cannot be negative")
    }
    b.price = price
    return nil
}

// Usage:
if err := book.SetPrice(-10); err != nil {
    fmt.Println("Error:", err)
}
```

## Constants and Static Members

### Python Constants

```python
class Book:
    CATEGORY_CODE = "BOOK"  # Class-level constant

    @classmethod
    def get_category_code(cls):
        return cls.CATEGORY_CODE
```

### Go Constants

```go
const CategoryCode = "BOOK"

func GetCategoryCode() string {
    return CategoryCode
}
```

## Key Differences

1. **Type System**

   - Python: Dynamic typing
   - Go: Static typing with type inference

2. **Object-Oriented Programming**

   - Python: Class-based with inheritance
   - Go: Struct-based with composition

3. **Error Handling**

   - Python: Exception-based
   - Go: Explicit error returns

4. **Access Modifiers**

   - Python: Convention-based (\_prefix)
   - Go: Capitalization-based (uppercase for exported)

5. **Method Receivers**
   - Python: Implicit self
   - Go: Explicit receiver parameter

## Practice Exercises

1. Create a new struct/class for `Library` that contains a collection of books
2. Implement methods to add/remove books
3. Create a method to calculate the total value of all books
4. Implement proper error handling for both languages

For more examples and documentation, contact:

- GitHub: [@rezashahnazar](https://github.com/rezashahnazar)
- Email: reza.shahnazar@gmail.com
