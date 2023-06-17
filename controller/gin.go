package controllers
 
import (
   "errors"
   "net/http"
 
   "example.com/go-crud-api/database"
   "github.com/gin-gonic/gin"
)

// CreateBooks creates multiple books.
//	@Summary		Create multiple books
//	@Description	Create multiple books
//	@Tags			Books
//	@Accept			json
//	@Produce		json
//	@Param			books	body		[]database.Book	true	"Books to create"
//	@Success		200		{object}	[]database.Book
//	@Failure		400		{object}	string	"Bad request"
//	@Router			/books [post]
func CreateBooks(c *gin.Context) {
	var books []*database.Book
	err := c.ShouldBind(&books)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	// Save each book in the database
	for _, book := range books {
		res := database.DB.Create(book)
		if res.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "error creating a book",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"books": books,
	})
	return
}

// CreateBook creates a new book.
//	@Summary		Create a book
//	@Description	Create a book
//	@Tags			Books
//	@Accept			json
//	@Produce		json
//	@Param			book	body		database.Book	true	"Book to create"
//	@Success		200		{object}	database.Book
//	@Failure		400		{object}	string	"Bad request"
//	@Router			/book [post]
func CreateBook(c *gin.Context) {
	var book *database.Book
	err := c.ShouldBind(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res := database.DB.Create(book)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating a book",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"book": book,
	})
	return
 }

 // ReadBook retrieves a book by ID.
//	@Summary		Retrieve a book by ID
//	@Description	Retrieve a book by ID
//	@Tags			Books
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Book ID"
//	@Success		200	{object}	database.Book
//	@Failure		404	{object}	string	"Book not found"
//	@Router			/book/{id} [get]
 func ReadBook(c *gin.Context) {
	var book database.Book
	id := c.Param("id")
	res := database.DB.Find(&book, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "book not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"book": book,
	})
	return
 }
 
 // ReadBooks retrieves all books.
//	@Summary		Retrieve all books
//	@Description	Retrieve all books
//	@Tags			Books
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]database.Book
//	@Failure		404	{object}	string	"Authors not found"
//	@Router			/books [get]
 func ReadBooks(c *gin.Context) {
	var books []database.Book
	res := database.DB.Find(&books)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("authors not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"books": books,
	})
	return
 }

 // UpdateBook updates a book by ID.
//	@Summary		Update a book by ID
//	@Description	Update a book by ID
//	@Tags			Books
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string			true	"Book ID"
//	@Param			book	body		database.Book	true	"Updated book data"
//	@Success		200		{object}	database.Book
//	@Failure		400		{object}	string	"Bad request"
//	@Router			/book/{id} [put]
 func UpdateBook(c *gin.Context) {
	var book database.Book
	id := c.Param("id")
	err := c.ShouldBind(&book)
   
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
  
	var updateBook database.Book
	res := database.DB.Model(&updateBook).Where("id = ?", id).Updates(book)
  
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "book not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"book": book,
	})
	return
 }

 // DeleteBook deletes a book by ID.
//	@Summary		Delete a book by ID
//	@Description	Delete a book by ID
//	@Tags			Books
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Book ID"
//	@Success		200	{object}	string	"Book deleted successfully"
//	@Failure		404	{object}	string	"Book not found"
//	@Router			/book/{id} [delete]

 func DeleteBook(c *gin.Context) {
	var book database.Book
	id := c.Param("id")
	res := database.DB.Find(&book, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "book not found",
		})
		return
	}
	database.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{
		"message": "book deleted successfully",
	})
	return
 }
 