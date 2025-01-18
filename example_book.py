import random


class Book:
    # Class-level constant representing the category code for all books
    CATEGORY_CODE = "BOOK"

    def __init__(self, title, author, price):
        """
        Initialize a new Book instance.
        Args:
            title (str): The title of the book
            author (str): The author's name
            price (float): The book's price
        """
        self._title = title
        self._author = author
        self._price = price
        self._page_count = self.random_page_count()

    def summary(self):
        """
        Returns a formatted string containing the book's details.
        Returns:
            str: A string with title, author, and price information
        """
        return f"{self._title} by {self._author} - ${self._price:.2f}"

    def get_price(self):
        """
        Getter method for the book's price.
        Returns:
            float: The current price of the book
        """
        return self._price

    def set_price(self, price):
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


def main():
    """
    Main function demonstrating the usage of the Book class.
    Shows various operations including:
    - Creating a book
    - Getting and setting prices
    - Using properties
    - Handling errors
    """
    # Create a new book instance
    harry_potter = Book("Harry Potter", "J.K. Rowling", 10.99)

    # Print the initial summary
    print(harry_potter.summary())

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


if __name__ == "__main__":
    main()

# Example Output:
# Harry Potter by J.K. Rowling - $10.99
# Harry Potter by J.K. Rowling - $12.99
# Price: 12.99
# Category Code: BOOK
# Page Count: 818
# Updated Page Count: 500
# Error: 'Book' object has no attribute '_page_count'
