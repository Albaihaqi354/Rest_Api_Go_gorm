package main

import (
	"fmt"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db := connctDb()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Connected error!")
	}
	defer sqlDB.Close()

	bookRepository := book.NewBookRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	route := gin.Default()
	v1 := route.Group("/v1")
	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)
	v1.POST("/books", bookHandler.CreateBook)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)

	route.Run()
}

func connctDb() *gorm.DB {
	var dsn = "host=localhost user=postgres password=Whobay123@ dbname=pustaka-api port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connected error!")
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Connected error!")
	}

	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Conncted!")
	}

	return db
}
