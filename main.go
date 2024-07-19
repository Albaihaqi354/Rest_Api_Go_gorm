package main

import (
	"net/http"
	"pustaka-api/entity/books"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.GET("/", rootHandler)
	route.GET("/hello", helloHandler)
	route.GET("/books/:id/:title", bookHandler)
	route.GET("/query", queryHandler)
	route.POST("/books", postBookHandle)

	route.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Name": "Bian Albaihaqi",
		"Bio":  "Back-end Development",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Kota":   "Garut",
		"Negara": "Indonesia",
	})
}

func bookHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{"title": title, "price": price})
}

func postBookHandle(c *gin.Context) {
	var bookInput books.BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"title":     bookInput.Title,
		"price":     bookInput.Price,
		"sub_title": bookInput.SubTitle,
	})
}
