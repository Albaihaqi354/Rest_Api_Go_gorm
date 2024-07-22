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
	// viewBook()
	// insertBook()
	// updateBook()
	// deleteBook()

	route := gin.Default()
	v1 := route.Group("/v1")
	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BookHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBookHandler)

	route.Run()
}

func viewBook() []book.Book {
	db := connctDb()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Connected error!")
	}
	defer sqlDB.Close()

	var books []book.Book

	err = db.Debug().Where("Id = ?", "1").Find(&books).Error
	if err != nil {
		fmt.Println("Error Book not Found")
	}

	for _, b := range books {
		fmt.Println("Title:", b.Id)
		fmt.Println("Book Object:", b)
	}

	return books
}

func insertBook() {
	db := connctDb()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Connected error!")
	}
	defer sqlDB.Close()

	db.AutoMigrate(&book.Book{})

	newBook := book.Book{
		Id:          2,
		Title:       "Google AI",
		Description: "Neuron AI From Google",
		Price:       100000,
		Rating:      5,
	}

	err = db.Create(&newBook).Error
	if err != nil {
		fmt.Println("Error Book not Cerated")
	} else {
		fmt.Println("Book Created")
	}

	return
}

func updateBook() {
	db := connctDb()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Connected error!")
	}
	defer sqlDB.Close()

	var book book.Book

	err = db.Debug().Where("Id = ?", 1).First(&book).Error
	if err != nil {
		fmt.Println("Error Book not Found")
	}

	book.Title = "Neuron"
	err = db.Save(&book).Error
	if err != nil {
		fmt.Println("Error Update Book")
	}

	return
}

func deleteBook() {
	db := connctDb()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Connected error!")
	}
	defer sqlDB.Close()

	var book book.Book

	err = db.Debug().Where("Id = ?", 2).First(&book).Error
	if err != nil {
		fmt.Println("Error Book not Found")
	}

	err = db.Delete(&book).Error
	if err != nil {
		fmt.Println("Error Delete data")
	}

	return
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
