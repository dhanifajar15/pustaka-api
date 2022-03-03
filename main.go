package main

import (
	"fmt"
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
	//CRUD

	book := book.Book{}
	book.Title = "Atomic Habits"
	book.Price = 12000
	book.Discount = 10
	book.Rating = 12
	book.Description = "Buku tentang membangun kebiasaan baik dan menghilangkan kebiasaan buruk"

	err = db.Create(&book).Error

	if err != nil {
		fmt.Println("===================")
		fmt.Println("Error Creating Book")
		fmt.Println("===================")
	}
	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BookHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run()
}
