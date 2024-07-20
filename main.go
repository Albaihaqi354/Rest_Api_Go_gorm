package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BookInput struct {
	Title string      `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required,number"`
}

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
	var bookInput BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {

		if validationError, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validationError {
				errorMassage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
				c.JSON(http.StatusBadRequest, errorMassage)
				return
			}
		}

		if UnmarshalTypeError, ok := err.(*json.UnmarshalTypeError); ok {
			errorMassage := fmt.Sprintf("Error on field %s, condition: %s", UnmarshalTypeError.Field, UnmarshalTypeError.Value)
			c.JSON(http.StatusBadRequest, errorMassage)
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})
}
