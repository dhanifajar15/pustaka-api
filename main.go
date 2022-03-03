package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"
)

func main() {
	router := gin.Default()
	dsn := "dhanifajar15:dhani@tcp(192.168.1.7:6033)/pustaka_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection error")
	}
	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	//bookFileRepository := book.NewFileRepository()
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	v1 := router.Group("/v1")

	v1.GET("/books", bookHandler.GetBooks)
	v1.POST("/books", bookHandler.PostBooksHandler)

	router.Run()
}
