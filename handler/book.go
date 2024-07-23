package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type bookHandeler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandeler {
	return &bookHandeler{bookService}
}

func (h *bookHandeler) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Name": "Bian Albaihaqi",
		"Bio":  "Back-end Development",
	})
}

func (h *bookHandeler) HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Kota":   "Garut",
		"Negara": "Indonesia",
	})
}

func (h *bookHandeler) BookHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func (h *bookHandeler) QueryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{"title": title, "price": price})
}

func (h *bookHandeler) PostBookHandler(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {

		errorMessages := []string{}

		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validationErrors {
				errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
		}

		if unmarshalTypeError, ok := err.(*json.UnmarshalTypeError); ok {
			errorMessage := fmt.Sprintf("Error on field %s, condition: should be %s", unmarshalTypeError.Field, unmarshalTypeError.Type)
			errorMessages = append(errorMessages, errorMessage)
		}

		if len(errorMessages) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	book, err := h.bookService.InsertBook(bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}
