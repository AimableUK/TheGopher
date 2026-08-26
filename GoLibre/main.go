package main

import (
	"context"
	"go-mongo-api/configs"
	"go-mongo-api/controllers"
	"go-mongo-api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	ctx := context.Background()

	configs.ConnectDB(ctx)
	go services.WatchBooksChanges() // Start watching changes

	// Routes
	router.POST("/books", controllers.CreateBook)
	router.GET("/books/:id", controllers.GetBook)
	router.GET("/list-books", controllers.ListBooks)
	router.GET("/books/count-by-category", controllers.ListBooksByCategory)
	router.GET("/search-books", controllers.SearchBooks)
	router.GET("/books/author/:authorId", controllers.GetAuthorBooks)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.PUT("/books/:id/details", controllers.UpdateBookDetails)
	router.DELETE("/books/:id", controllers.DeleteBook)

	router.Run(":8080")
}
