from abc import ABC, abstractmethod
import random


# Add this interface-like abstract class
class PricedItem(ABC):
    """
    Abstract base class representing items that have a price.
    This demonstrates a concept similar to Go interfaces.
    In Go, interfaces are implicit - any type that implements
    the required methods automatically satisfies the interface.
    """

    @abstractmethod
    def get_price(self) -> float:
        """Get the item's price"""
        pass

    @abstractmethod
    def set_price(self, price: float) -> None:
        """Set the item's price"""
        pass

    @abstractmethod
    def calculate_discount(self, percentage: float) -> float:
        """Calculate discounted price"""
        pass


# Modify the Book class to implement the interface
class Book(PricedItem):
    # Class-level constant representing the category code for all books
    CATEGORY_CODE = "BOOK"

    def __init__(self, title, author, price, seller=None):
        """
        Initialize a new Book instance.
        Args:
            title (str): The title of the book
            author (str): The author's name
            price (float): The book's price
            seller (str): The seller's name (public attribute, can be modified directly)
        """
        self._title = title
        self._author = author
        self._price = price
        self._page_count = self.random_page_count()
        # seller is intentionally not private (no underscore)
        # to demonstrate different attribute access patterns
        self.seller = seller

    def summary(self):
        """
        Returns a formatted string containing the book's details.
        Returns:
            str: A string with title, author, and price information
        """
        return f"{self._title} by {self._author} - ${self._price:.2f}"

    def get_price(self) -> float:
        """
        Getter method for the book's price.
        Returns:
            float: The current price of the book
        """
        return self._price

    def set_price(self, price: float) -> None:
        """
        Setter method for the book's price.
        Args:
            price (float): The new price to set
        Raises:
            ValueError: If the price is negative
        """
        if price < 0:
            raise ValueError("Price cannot be negative")
        self._price = price

    @classmethod
    def get_category_code(cls):
        """
        Class method to retrieve the category code.
        Returns:
            str: The category code for all books
        """
        return cls.CATEGORY_CODE

    @staticmethod
    def random_page_count():
        """
        Static method to generate a random page count.
        Returns:
            int: A random number between 100 and 1000
        """
        return random.randint(100, 1000)

    @property
    def page_count(self):
        """
        Property getter for page count.
        Returns:
            int: The current page count
        """
        return self._page_count

    @page_count.setter
    def page_count(self, value):
        """
        Property setter for page count.
        Args:
            value (int): The new page count to set
        """
        self._page_count = value

    @page_count.deleter
    def page_count(self):
        """
        Property deleter for page count.
        Removes the page count attribute from the instance
        """
        del self._page_count

    # Add this method to fulfill the PricedItem interface
    def calculate_discount(self, percentage: float) -> float:
        """
        Calculate the discounted price
        Args:
            percentage (float): Discount percentage (0-100)
        Returns:
            float: Discounted price
        """
        if not 0 <= percentage <= 100:
            raise ValueError("Percentage must be between 0 and 100")
        return self._price * (1 - percentage / 100)


# Add another class that implements the same interface
class Magazine(PricedItem):
    """
    Magazine class implementing the same PricedItem interface.
    This demonstrates how different types can implement the same interface.
    """

    def __init__(self, name: str, price: float, issue_number: int):
        self._name = name
        self._price = price
        self._issue_number = issue_number

    def get_price(self) -> float:
        return self._price

    def set_price(self, price: float) -> None:
        if price < 0:
            raise ValueError("Price cannot be negative")
        self._price = price

    def calculate_discount(self, percentage: float) -> float:
        if not 0 <= percentage <= 100:
            raise ValueError("Percentage must be between 0 and 100")
        # Magazines have a different discount calculation
        # (just as an example of different implementations)
        base_discount = self._price * (1 - percentage / 100)
        # Additional 10% off for magazines over $10
        if self._price > 10:
            return base_discount * 0.9
        return base_discount


def print_item_price_info(item: PricedItem):
    """
    Function demonstrating interface-like behavior.
    This is similar to how Go interfaces work - we can pass
    any object that implements the PricedItem interface.
    """
    print(f"Original price: ${item.get_price():.2f}")
    print(f"Price with 20% discount: ${item.calculate_discount(20):.2f}")


def main():
    """
    Main function demonstrating the usage of the Book class.
    Shows various operations including:
    - Creating a book
    - Getting and setting prices
    - Using properties
    - Handling errors
    - Working with the public seller attribute
    """
    # Create a new book instance with a seller
    harry_potter = Book("Harry Potter", "J.K. Rowling", 10.99, "Flourish & Blotts")

    # Print the initial summary
    print(harry_potter.summary())

    # Demonstrate direct access to public seller attribute
    print("Original Seller:", harry_potter.seller)

    # Modify seller directly (possible because it's a public attribute)
    harry_potter.seller = "Obscurus Books"
    print("New Seller:", harry_potter.seller)

    # Demonstrate price setting with error handling
    try:
        harry_potter.set_price(12.99)
    except ValueError as e:
        print("Error:", e)

    # Print the updated summary
    print(harry_potter.summary())

    # Demonstrate getter method
    print("Price:", harry_potter.get_price())

    # Demonstrate class method
    print("Category Code:", Book.get_category_code())

    # Demonstrate property getter
    print("Page Count:", harry_potter.page_count)

    # Demonstrate property setter
    harry_potter.page_count = 500
    print("Updated Page Count:", harry_potter.page_count)

    # Demonstrate property deleter and error handling
    del harry_potter.page_count
    try:
        print("Deleted Page Count:", harry_potter.page_count)
    except AttributeError as e:
        print("Error:", e)

    # Add magazine demo
    vogue = Magazine("Vogue", 12.99, 123)

    print("\n=== Demonstrating interface-like behavior ===")
    print("Book pricing:")
    print_item_price_info(harry_potter)

    print("\nMagazine pricing:")
    print_item_price_info(vogue)


if __name__ == "__main__":
    main()

# Example Output:
"""
Harry Potter by J.K. Rowling - $10.99
Original Seller: Flourish & Blotts
New Seller: Obscurus Books
Harry Potter by J.K. Rowling - $12.99
Price: 12.99
Category Code: BOOK
Page Count: 307
Updated Page Count: 500
Error: 'Book' object has no attribute '_page_count'

=== Demonstrating interface-like behavior ===
Book pricing:
Original price: $12.99
Price with 20% discount: $10.39

Magazine pricing:
Original price: $12.99
Price with 20% discount: $9.35
"""
