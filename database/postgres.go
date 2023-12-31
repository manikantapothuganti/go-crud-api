package database
 
import (
   "fmt"
   "log"
 
   "gorm.io/driver/postgres"
   "gorm.io/gorm"
)
 
var DB *gorm.DB
var err error
 
// Book represents a book in the database.
type Book struct {
   gorm.Model
   Title  string `json:"title" example:"One Piece"`
   Author string `json:"author" example:"Oda"`
}
 
// DatabaseConnection establishes a connection to the database.
func DatabaseConnection() {
   host := "localhost"
   port := "5432"
   dbName := "postgres"
   dbUser := "postgres"
   password := "mypassword"
   dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
       host,
       port,
       dbUser,
       dbName,
       password,
   )
 
   DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
   DB.AutoMigrate(Book{})
   if err != nil {
       log.Fatal("Error connecting to the database...", err)
   }
   fmt.Println("Database connection successful...")
}
