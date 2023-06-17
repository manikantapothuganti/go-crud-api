package main

import (
	"fmt"

	"example.com/go-crud-api/controller"
	"example.com/go-crud-api/database"
	"example.com/go-crud-api/docs"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@termsOfService	http://swagger.io/terms/
//	@contact.name	Mani
//	@contact.url	http://www.swagger.io/support
//	@contact.email	manidpothuganti@gmail.com
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html


//	@securityDefinitions.basic	No Auth

func main() {

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Go Books API CRUD"
	docs.SwaggerInfo.Description = "This is a sample server books server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:5009"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	fmt.Println("Starting application ...")
	database.DatabaseConnection()

	r := gin.Default()

	// ReadBook retrieves a book by ID.
	r.GET("/books/:id", controllers.ReadBook)

	// ReadBooks retrieves all books.
	r.GET("/books", controllers.ReadBooks)

	// CreateBooks creates multiple books.
	r.POST("/books", controllers.CreateBooks)

	// CreateBook creates a new book.
	r.POST("/book", controllers.CreateBook)

	// UpdateBook updates a book by ID.
	r.PUT("/books/:id", controllers.UpdateBook)

	// DeleteBook deletes a book by ID.
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":5009")
}
