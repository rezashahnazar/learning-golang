package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ==================== INTRODUCTION TO GO BASICS ====================
// Unlike Python, Go is a statically typed language.
// We need to declare types for all variables and struct fields.
// Go doesn't have classes, but we use structs and methods instead.
// Go doesn't have decorators like @property, but we can achieve similar functionality with methods.

// ==================== CONSTANTS ====================
// In Go, constants are declared using the const keyword
// They are typically placed at package level
const CategoryCode = "BOOK"

// ==================== STRUCT DEFINITION ====================
// Instead of Python classes, Go uses structs
// The struct fields starting with capital letters are exported (public)
// Fields starting with lowercase are private to the package
type Book struct {
    // In Go, we don't need underscores for private fields
    // Lowercase first letter makes them private by default
    title     string
    author    string
    price     float64
    pageCount int
}

// ==================== CONSTRUCTOR ====================
// Go doesn't have constructors like Python's __init__
// Instead, we create constructor functions by convention
// They usually start with "New" followed by the type name
func NewBook(title string, author string, price float64) *Book {
    // Initialize random seed (similar to Python's random)
    rand.Seed(time.Now().UnixNano())
    
    return &Book{
        title:     title,
        author:    author,
        price:     price,
        pageCount: randomPageCount(),
    }
}

// ==================== METHODS ====================
// In Go, we define methods with a receiver parameter
// This is similar to Python's self parameter

// Summary method - equivalent to Python's summary()
func (b *Book) Summary() string {
    // Go's string formatting is similar to Python's f-strings
    // but uses different syntax
    return fmt.Sprintf("%s by %s - $%.2f", b.title, b.author, b.price)
}

// GetPrice - equivalent to Python's get_price()
func (b *Book) GetPrice() float64 {
    return b.price
}

// SetPrice - equivalent to Python's set_price()
func (b *Book) SetPrice(price float64) error {
    // Error handling in Go uses explicit return values
    // instead of raising exceptions
    if price < 0 {
        return fmt.Errorf("price cannot be negative")
    }
    b.price = price
    return nil
}

// ==================== STATIC METHODS ====================
// In Go, static methods are just regular functions
// They don't need a receiver parameter

// GetCategoryCode - equivalent to Python's get_category_code()
func GetCategoryCode() string {
    return CategoryCode
}

// randomPageCount - equivalent to Python's random_page_count()
func randomPageCount() int {
    // Generate random number between 100 and 1000
    return rand.Intn(901) + 100 // 901 = 1000 - 100 + 1
}

// ==================== PROPERTY-LIKE METHODS ====================
// Go doesn't have properties like Python
// We use conventional getter/setter methods instead

func (b *Book) GetPageCount() int {
    return b.pageCount
}

func (b *Book) SetPageCount(value int) {
    b.pageCount = value
}

// ==================== MAIN FUNCTION ====================
func main() {
    // Create a new book instance
    harryPotter := NewBook("Harry Potter", "J.K. Rowling", 10.99)

    // Print initial summary
    fmt.Println(harryPotter.Summary())

    // Demonstrate error handling with SetPrice
    if err := harryPotter.SetPrice(12.99); err != nil {
        fmt.Println("Error:", err)
    }

    // Print updated summary
    fmt.Println(harryPotter.Summary())

    // Demonstrate getter method
    fmt.Printf("Price: %.2f\n", harryPotter.GetPrice())

    // Demonstrate static function
    fmt.Println("Category Code:", GetCategoryCode())

    // Demonstrate property-like methods
    fmt.Println("Page Count:", harryPotter.GetPageCount())

    // Set new page count
    harryPotter.SetPageCount(500)
    fmt.Println("Updated Page Count:", harryPotter.GetPageCount())
}

// ==================== KEY DIFFERENCES FROM PYTHON ====================
// 1. Go uses explicit error handling with return values instead of exceptions
// 2. Go uses structs and methods instead of classes
// 3. Go doesn't have decorators or properties
// 4. Go is statically typed - all types must be declared
// 5. Go uses pointer receivers (*Book) for methods that modify the struct
// 6. Go's visibility is controlled by capitalization (uppercase = public, lowercase = private)
// 7. Go doesn't have a direct equivalent to Python's property deleters
