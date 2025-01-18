// =================== GO PROGRAMMING TUTORIAL ===================
// This tutorial will teach you Go by comparing it with Python concepts
// We'll build a book management system while learning core Go concepts

// ------------------- PACKAGE DECLARATION ---------------------
// Every Go file must start with a package declaration
// 'main' is a special package name - it tells Go this is an executable program
// If this was a library, we might name it something like 'bookstore'
package main

// ------------------- IMPORTS -------------------------------
// Go uses explicit imports for each package needed
// Unlike Python, you can't use a package without importing it
// The import syntax uses quotes, different from Python's plain imports
import (
	// fmt is Go's standard package for formatted I/O operations
	// similar to Python's print() and string formatting
	"fmt"

	// math/rand is for random number generation
	// notice how sub-packages use "/" unlike Python's "."
	"math/rand"
)

// ------------------- INTERFACES ---------------------------
// Interfaces in Go are fundamentally different from Python's ABC
// 1. They only declare method signatures (no implementations)
// 2. They are implemented implicitly (no "implements" keyword needed)
// 3. They are typically small, often just 1-2 methods
// 4. They are satisfied by any type that implements all their methods
type PricedItem interface {
    // Method declarations show:
    // - Name of method
    // - Return type(s) after the parentheses
    // - No function body (just declarations)
    Price() float64
    SetPrice(price float64) error
    CalculateDiscount(percentage float64) (float64, error)
}

// ------------------- STRUCTS -----------------------------
// Structs are Go's way of defining custom data types
// Unlike Python classes:
// 1. No inheritance (Go favors composition over inheritance)
// 2. No constructor method
// 3. No instance methods inside the struct definition
// 4. Fields must have explicit types
type Book struct {
    // Go's field visibility is controlled by capitalization:
    // lowercase = private (package-level)
    // uppercase = public (exported)
    title      string  // private, like Python's _title
    author     string  // private, like Python's _author
    price      float64 // private, like Python's _price
    pageCount  int     // private, like Python's _page_count
    Seller     string  // public, like Python's seller (no underscore)
}

// ------------------- CONSTANTS ---------------------------
// Constants in Go are declared at package level
// Unlike Python, Go has true constants that cannot be changed
// Naming convention: Use MixedCaps or ALL_CAPS for constants
const CategoryCode = "BOOK"

// ------------------- CONSTRUCTORS ------------------------
// Go doesn't have built-in constructors like Python's __init__
// Instead, we use factory functions, typically prefixed with "New"
// This is a common Go pattern for object creation
func NewBook(title, author string, price float64, seller string) *Book {
    // The * before Book means this returns a pointer
    // Pointers are a core Go concept with no Python equivalent
    // They hold the memory address of values
    
    // Return a new Book instance
    // The & operator creates a pointer to the struct
    return &Book{
        // Field initialization uses name: value syntax
        // Similar to Python's kwargs but with colons
        title:     title,
        author:    author,
        price:     price,
        pageCount: randomPageCount(),
        Seller:    seller,
    }
}

// ------------------- METHODS -----------------------------
// Go methods are declared outside the struct
// The (b *Book) is called a "receiver" - it's like Python's self
// But in Go, we explicitly say if we're using a pointer (*Book)
// or value (Book) receiver
func (b *Book) Summary() string {
    // fmt.Sprintf is like Python's f-strings
    // %.2f formats float with 2 decimal places
    return fmt.Sprintf("%s by %s - $%.2f", b.title, b.author, b.price)
}

// Interface implementation for Book
// Notice how we don't need to explicitly state that we're
// implementing PricedItem - Go does this implicitly
func (b *Book) Price() float64 {
    return b.price
}

// ------------------- ERROR HANDLING ----------------------
// Go handles errors very differently from Python:
// 1. No try/except blocks
// 2. Errors are return values, not exceptions
// 3. Multiple return values are common (value, error)
func (b *Book) SetPrice(price float64) error {
    // Error checking is explicit
    if price < 0 {
        // fmt.Errorf creates a new error with formatted text
        return fmt.Errorf("price cannot be negative")
    }
    b.price = price
    // nil is Go's equivalent of None
    return nil
}

func (b *Book) CalculateDiscount(percentage float64) (float64, error) {
    // Multiple return values are idiomatic in Go
    // This is different from Python's single return value
    if percentage < 0 || percentage > 100 {
        return 0, fmt.Errorf("percentage must be between 0 and 100")
    }
    return b.price * (1 - percentage/100), nil
}

// ------------------- HELPER FUNCTIONS --------------------
// Package-level functions (not methods) don't have receivers
// They're like Python's module-level functions
func GetCategoryCode() string {
    return CategoryCode
}

// Private helper function (lowercase name)
func randomPageCount() int {
    // rand.Intn(n) generates 0 to n-1
    // Adding 100 gives us 100 to 1000
    return rand.Intn(901) + 100
}

// ------------------- MULTIPLE TYPES ---------------------
// Go encourages small, focused types that satisfy interfaces
type Magazine struct {
    name        string
    price       float64
    issueNumber int
}

// Constructor for Magazine
func NewMagazine(name string, price float64, issueNumber int) *Magazine {
    return &Magazine{
        name:        name,
        price:       price,
        issueNumber: issueNumber,
    }
}

// Magazine methods implementing PricedItem interface
func (m *Magazine) Price() float64 {
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
    if m.price > 10 {
        return baseDiscount * 0.9, nil
    }
    return baseDiscount, nil
}

// ------------------- INTERFACE USAGE -------------------
// This function demonstrates polymorphism in Go
// It accepts any type that implements PricedItem
func printItemPriceInfo(item PricedItem) {
    // Direct price access through interface method
    fmt.Printf("Original price: $%.2f\n", item.Price())
    
    // Error handling in Go is explicit and required
    discounted, err := item.CalculateDiscount(20)
    // if err != nil is the most common error check in Go
    if err != nil {
        fmt.Printf("Error calculating discount: %v\n", err)
        return
    }
    fmt.Printf("Price with 20%% discount: $%.2f\n", discounted)
}

// ------------------- MAIN FUNCTION ---------------------
// main() is the entry point of a Go program
// Like Python's if __name__ == "__main__":
func main() {
    // := is a shorthand declaration operator
    // It declares and initializes variables in one step
    harryPotter := NewBook("Harry Potter", "J.K. Rowling", 10.99, "Flourish & Blotts")

    // Calling methods uses dot notation like Python
    fmt.Println(harryPotter.Summary())

    // Public fields can be accessed directly
    fmt.Println("Original Seller:", harryPotter.Seller)
    harryPotter.Seller = "Obscurus Books"
    fmt.Println("New Seller:", harryPotter.Seller)

    // Error handling pattern in Go:
    // 1. Call function that returns error
    // 2. Check if error is nil
    // 3. Handle error if present
    if err := harryPotter.SetPrice(12.99); err != nil {
        fmt.Println("Error:", err)
    }

    fmt.Println(harryPotter.Summary())
    fmt.Println("Price:", harryPotter.Price())
    fmt.Println("Category Code:", GetCategoryCode())

    // Creating a magazine instance
    vogue := NewMagazine("Vogue", 12.99, 123)

    fmt.Println("\n=== Demonstrating interface behavior ===")
    fmt.Println("Book pricing:")
    printItemPriceInfo(harryPotter)

    fmt.Println("\nMagazine pricing:")
    printItemPriceInfo(vogue)
}

/* ------------------- EXAMPLE OUTPUT -------------------

Running this program will produce output similar to:

Harry Potter by J.K. Rowling - $10.99
Original Seller: Flourish & Blotts
New Seller: Obscurus Books
Harry Potter by J.K. Rowling - $12.99
Price: 12.99
Category Code: BOOK

=== Demonstrating interface behavior ===
Book pricing:
Original price: $12.99
Price with 20% discount: $10.39

Magazine pricing:
Original price: $12.99
Price with 20% discount: $9.35

Note: The page count will be random each time you run the program.

------------------- ADDITIONAL GO CONCEPTS -------------------

1. GOROUTINES AND CONCURRENCY
   - Go has built-in concurrency support using goroutines
   - Much simpler than Python's threading/multiprocessing
   - Use 'go' keyword to start a goroutine
   - Channels for communication between goroutines

2. DEFER STATEMENT
   - defer delays execution until surrounding function returns
   - Commonly used for cleanup operations
   - Similar to Python's context managers but more flexible

3. SLICES AND ARRAYS
   - Arrays have fixed size
   - Slices are dynamic (like Python lists)
   - Powerful slice operations with memory efficiency

4. MAPS
   - Similar to Python dictionaries
   - Must specify key and value types
   - Written as map[KeyType]ValueType

5. ZERO VALUES
   - Every type has a zero value
   - numbers: 0
   - strings: ""
   - pointers: nil
   - etc.

6. BUILD SYSTEM
   - go build: compiles packages
   - go run: compiles and runs
   - go test: runs tests
   - go mod: manages dependencies

7. TESTING
   - Built-in testing framework
   - Files end in _test.go
   - Run with 'go test'

8. DOCUMENTATION
   - Documentation comments start with //
   - godoc generates documentation
   - View with 'go doc'

9. FORMATTING
   - gofmt automatically formats code
   - No debates about style
   - Run with 'go fmt'

10. MODULES
    - go.mod defines module
    - Similar to Python's requirements.txt
    - But more integrated with language

------------------- COMMON GO PATTERNS -------------------

1. ERROR HANDLING
   if err != nil {
       return nil, err
   }

2. BUILDER PATTERN
   options := NewOptions().
       WithName("name").
       WithAge(25)

3. FUNCTIONAL OPTIONS
   func WithTimeout(t time.Duration) Option {
       return func(o *Options) {
           o.timeout = t
       }
   }

4. INTERFACE COMPOSITION
   type Reader interface {
       Read(p []byte) (n int, err error)
   }

5. EMBEDDING
   type Client struct {
       *http.Client
   }

------------------- GO VS PYTHON SUMMARY -------------------

1. TYPING
   Python: Dynamic typing
   Go: Static typing

2. CONCURRENCY
   Python: GIL, multiprocessing
   Go: Goroutines, channels

3. ERROR HANDLING
   Python: try/except
   Go: explicit error returns

4. OBJECTS
   Python: Classes with inheritance
   Go: Structs with composition

5. INTERFACES
   Python: Explicit ABC
   Go: Implicit satisfaction

6. PACKAGING
   Python: pip, virtualenv
   Go: go modules

7. FORMATTING
   Python: PEP 8 (convention)
   Go: gofmt (enforced)

8. VISIBILITY
   Python: _underscore convention
   Go: Capitalization rules

9. GENERICS
   Python: Built-in
   Go: Added in Go 1.18+

10. PHILOSOPHY
    Python: "There should be one obvious way to do it"
    Go: "Keep it simple and straightforward"
*/
