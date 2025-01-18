// ==================== INTRODUCTION TO GO ====================
// Go is a statically-typed, compiled language created by Google.
// Unlike Python's interpreter, Go code is compiled directly to machine code.
// This makes it typically faster than Python but requires explicit type declarations.

// ==================== PACKAGE DECLARATION ====================
// Every Go file must start with a package declaration
// 'package main' is special - it indicates this is an executable program, not a library
package main

// ==================== IMPORTS ====================
// Go's import system is different from Python's
// - In Python: 'import random'
// - In Go: we use the full import path and group imports in parentheses
import (
	// fmt is Go's standard package for formatted I/O (similar to Python's print())
	"fmt"
	// math/rand is for random number generation (similar to Python's random module)
	"math/rand"
	// time is needed for random seed initialization
	"time"
)

// ==================== INTERFACE DEFINITION ====================
// In Python, we used ABC (Abstract Base Class) for PricedItem
// In Go, we use interfaces instead. Key differences:
// 1. Interfaces only declare method signatures
// 2. Types implicitly implement interfaces (no explicit declaration needed)
// 3. Interfaces are typically smaller in Go (following interface segregation principle)
type PricedItem interface {
    // Method declarations include parameter types and return types
    // float64 is Go's double-precision floating-point type (similar to Python's float)
    GetPrice() float64
    
    // Go methods can return multiple values
    // Here we return both float64 and error (Go's error handling mechanism)
    SetPrice(float64) error
    
    // Multiple return values are grouped in parentheses
    CalculateDiscount(float64) (float64, error)
}

// ==================== STRUCT DEFINITION ====================
// Instead of Python classes, Go uses structs
// Structs are collections of fields (similar to class attributes)
type Book struct {
    // Field naming conventions in Go:
    // - lowercase first letter = private (package-level visibility)
    // - uppercase first letter = public (exported, visible outside package)
    title      string   // private, like Python's self._title
    author     string   // private, like Python's self._author
    price      float64  // private, like Python's self._price
    pageCount  int      // private, like Python's self._page_count
    Seller     string   // public, like Python's self.seller
}

// ==================== CONSTANTS ====================
// Constants in Go are declared using the const keyword
// Unlike Python's class-level constants, these are package-level
const CategoryCode = "BOOK"

// ==================== CONSTRUCTOR ====================
// Go doesn't have built-in constructors like Python's __init__
// Instead, we use factory functions, typically prefixed with "New"
// The * before Book means this returns a pointer to a Book
func NewBook(title string, author string, price float64, seller string) *Book {
    // The & operator creates a pointer to a new struct instance
    return &Book{
        // In struct initialization, we assign values to fields
        title:     title,
        author:    author,
        price:     price,
        pageCount: randomPageCount(),
        Seller:    seller,
    }
}

// ==================== METHODS ====================
// Go methods have a "receiver" parameter in parentheses before the method name
// This is similar to Python's self parameter
// (b *Book) means this method operates on a pointer to a Book
func (b *Book) Summary() string {
    // fmt.Sprintf is similar to Python's f-strings
    return fmt.Sprintf("%s by %s - $%.2f", b.title, b.author, b.price)
}

// Implementation of PricedItem interface methods
// Note: Go automatically knows this implements PricedItem because it has all required methods
func (b *Book) GetPrice() float64 {
    return b.price
}

// ==================== ERROR HANDLING ====================
// Go doesn't use exceptions like Python
// Instead, functions return error values that must be checked
func (b *Book) SetPrice(price float64) error {
    if price < 0 {
        // fmt.Errorf creates a new error with formatted text
        return fmt.Errorf("price cannot be negative")
    }
    b.price = price
    // nil is Go's version of None/null
    return nil
}

func (b *Book) CalculateDiscount(percentage float64) (float64, error) {
    if percentage < 0 || percentage > 100 {
        return 0, fmt.Errorf("percentage must be between 0 and 100")
    }
    return b.price * (1 - percentage/100), nil
}

// ==================== UTILITY FUNCTIONS ====================
// Regular functions (not methods) don't have a receiver parameter
func randomPageCount() int {
    // rand.Intn(n) generates numbers from 0 to n-1
    // We add 100 to get a range of 100-1000
    return rand.Intn(901) + 100
}

// ==================== MAGAZINE IMPLEMENTATION ====================
// Another struct implementing the same interface
// This demonstrates Go's interface polymorphism
type Magazine struct {
    name        string
    price       float64
    issueNumber int
}

func NewMagazine(name string, price float64, issueNumber int) *Magazine {
    return &Magazine{
        name:        name,
        price:       price,
        issueNumber: issueNumber,
    }
}

// Magazine's implementation of PricedItem interface
func (m *Magazine) GetPrice() float64 {
    return m.price
}

func (m *Magazine) SetPrice(price float64) error {
    if price < 0 {
        return fmt.Errorf("price cannot be negative")
    }
    m.price = price
    return nil
}

func (m *Magazine) CalculateDiscount(percentage float64) (float64, error) {
    if percentage < 0 || percentage > 100 {
        return 0, fmt.Errorf("percentage must be between 0 and 100")
    }
    baseDiscount := m.price * (1 - percentage/100)
    // Additional 10% off for magazines over $10
    if m.price > 10 {
        return baseDiscount * 0.9, nil
    }
    return baseDiscount, nil
}

// ==================== INTERFACE USAGE ====================
// This function demonstrates Go's interface polymorphism
// It can accept any type that implements PricedItem
func printItemPriceInfo(item PricedItem) {
    // Get the original price
    price := item.GetPrice()
    
    // Calculate discount, checking for errors
    // The := operator is a shorthand for declaring and initializing variables
    discounted, err := item.CalculateDiscount(20)
    
    // Error handling in Go is explicit
    if err != nil {
        fmt.Printf("Error calculating discount: %v\n", err)
        return
    }
    
    // Printf is similar to Python's formatted strings
    // Note the %% to print a literal % symbol
    fmt.Printf("Original price: $%.2f\n", price)
    fmt.Printf("Price with 20%% discount: $%.2f\n", discounted)
}

// ==================== ADDITIONAL BOOK METHODS ====================
// GetCategoryCode is similar to Python's class method
// In Go, we just use a regular function since we don't need
// class-level functionality like Python's @classmethod
func GetCategoryCode() string {
    return CategoryCode
}

// GetPageCount is similar to Python's @property decorator
// In Go, we use regular methods for property-like access
func (b *Book) GetPageCount() int {
    return b.pageCount
}

// SetPageCount is similar to Python's @property.setter
func (b *Book) SetPageCount(value int) {
    b.pageCount = value
}

// ==================== MAIN FUNCTION ====================
// The main() function is the entry point of the program
// Similar to Python's if __name__ == "__main__":
func main() {
    // Initialize random seed for random number generation
    // This is similar to Python's random.seed()
    rand.Seed(time.Now().UnixNano())

    // Create a new book instance
    // := is used for declaring and initializing variables in one line
    harryPotter := NewBook("Harry Potter", "J.K. Rowling", 10.99, "Flourish & Blotts")

    // Demonstrate various operations
    fmt.Println(harryPotter.Summary())

    // Accessing public fields (notice the capital letter)
    fmt.Println("Original Seller:", harryPotter.Seller)
    harryPotter.Seller = "Obscurus Books"
    fmt.Println("New Seller:", harryPotter.Seller)

    // Error handling example
    // In Go, we must check errors explicitly
    if err := harryPotter.SetPrice(12.99); err != nil {
        fmt.Println("Error:", err)
    }

    fmt.Println(harryPotter.Summary())

    // Create a magazine instance
    vogue := NewMagazine("Vogue", 12.99, 123)

    // Demonstrate interface usage with both types
    fmt.Println("\n=== Demonstrating interface behavior ===")
    fmt.Println("Book pricing:")
    printItemPriceInfo(harryPotter)

    fmt.Println("\nMagazine pricing:")
    printItemPriceInfo(vogue)

    // Additional demonstrations to match Python output
    fmt.Printf("Price: %.2f\n", harryPotter.GetPrice())
    fmt.Printf("Category Code: %s\n", GetCategoryCode())
    
    // Demonstrate page count operations
    fmt.Printf("Page Count: %d\n", harryPotter.GetPageCount())
    harryPotter.SetPageCount(500)
    fmt.Printf("Updated Page Count: %d\n", harryPotter.GetPageCount())

    // Note: Go doesn't have direct equivalent to Python's property deleter
    // Memory management is handled differently in Go

    // ==================== EXPECTED OUTPUT EXPLANATION ====================
    /*
    Expected Output:
    Harry Potter by J.K. Rowling - $10.99        // From harryPotter.Summary()
    Original Seller: Flourish & Blotts           // Direct access to public Seller field
    New Seller: Obscurus Books                   // After modifying Seller field
    Harry Potter by J.K. Rowling - $12.99        // After SetPrice(12.99)

    === Demonstrating interface behavior ===
    Book pricing:                                // From printItemPriceInfo(harryPotter)
    Original price: $12.99                       // From GetPrice()
    Price with 20% discount: $10.39              // From CalculateDiscount(20)

    Magazine pricing:                            // From printItemPriceInfo(vogue)
    Original price: $12.99                       // From GetPrice()
    Price with 20% discount: $9.35               // From CalculateDiscount(20) with extra 10% off
    Price: 12.99                                 // Direct GetPrice() call
    Category Code: BOOK                          // From GetCategoryCode()
    Page Count: 418                              // From GetPageCount() (random value)
    Updated Page Count: 500                      // After SetPageCount(500)
    */

    // ==================== OUTPUT BREAKDOWN ====================
    // 1. Initial output shows the book's creation and basic string representation
    // 2. Seller modification demonstrates public field access
    // 3. Price update shows the effect of SetPrice method
    // 4. Interface behavior shows how both Book and Magazine implement PricedItem
    // 5. Magazine's special discount (additional 10% off) is visible in its output
    // 6. Final section shows property-like access to various fields

    // Note: The random page count (418 in example) will vary in each run
    // because it's generated using randomPageCount()
}

// ==================== KEY DIFFERENCES FROM PYTHON ====================
// 1. Explicit type declarations
// 2. Pointers and memory management
// 3. Error handling instead of exceptions
// 4. Public/private determined by capitalization
// 5. Interfaces instead of abstract classes
// 6. No inheritance (composition over inheritance)
// 7. No decorators or properties
// 8. Compiled vs interpreted
// 9. Strict formatting rules (enforced by go fmt)
// 10. Built-in concurrency support (not shown in this example)

// ==================== IMPORTANT NOTE ====================
// Some Python features don't have direct equivalents in Go:
// 1. Property decorators (@property) - Use methods instead
// 2. Deleters - Go uses garbage collection
// 3. Class methods (@classmethod) - Use package-level functions
// 4. Static methods (@staticmethod) - Use package-level functions
// 5. Dynamic attribute deletion - Not available in Go
