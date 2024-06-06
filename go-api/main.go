package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// create the book struct
type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books) // specify the status
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	//show the new book list with the added book
	c.IndentedJSON(http.StatusCreated, books)

}

func getBookByID(c *gin.Context) {
	var id string = c.Param("id")
	rBook, err := getBook(id)
	if err != nil {
		c.String(http.StatusNotFound, err.Error()) // err.Error() to get the error message
		return
	}
	c.IndentedJSON(http.StatusFound, rBook)

}
func getBook(id string) (*book, error) {
	for i, book := range books {
		if book.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func checkOutBook(c *gin.Context) {
	id := c.Query("b")
	borrow := c.Param("borr")
	book, err := updateList(id, borrow)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, book)

}

func updateList(b string, borrow string) (*book, error) {

	rBook, err := getBook(b)
	if err != nil {
		return nil, err
	}
	if borrow == "1" && rBook.Quantity <= 0 {
		return nil, errors.New("book is out of stock")
	}
	if borrow == "1" {
		rBook.Quantity -= 1
	} else {
		rBook.Quantity += 1
	}
	return rBook, nil
}

func main() {
	//define the router for the gin

	router := gin.Default() // return a *gin.engine
	//specify the route and the function to be called
	router.GET("/books", getBooks)

	//use >localhost:8080/books --include --header "Content-Type: application/json" -d @body.json --request "POST"
	//where @body.json is the file
	router.POST("/books", createBook)

	//fetch the book by the specified id as /books/:id
	//when giving the request dont put : i.e /books/1
	router.GET("/books/:id", getBookByID)

	//check out the book and this reduces the quantity in the list by 1
	//or return the book using the same route
	// localhost:8080/check/1/?b=2 this is how you do the request
	router.PATCH("/check/:borr", checkOutBook)

	//Listen to the port 8080
	router.Run("localhost:8080")

}
