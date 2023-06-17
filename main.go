package main
 
import (
   "fmt"
 
   "example.com/go-crud-api/controller"
   "example.com/go-crud-api/database"
   "github.com/gin-gonic/gin"
)
 
func main() {
   fmt.Println("Starting application ...")
   database.DatabaseConnection()
 
   r := gin.Default()
   r.GET("/books/:id", controllers.ReadBook)
   r.GET("/books", controllers.ReadBooks)
   r.POST("/books", controllers.CreateBooks)
   r.POST("/book", controllers.CreateBook)
   r.PUT("/books/:id", controllers.UpdateBook)
   r.DELETE("/books/:id", controllers.DeleteBook)
   r.Run(":5009")
}
