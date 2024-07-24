package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pustaka-api/book"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type bookHandeler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandeler {
	return &bookHandeler{bookService}
}

func (h *bookHandeler) GetBooks(c *gin.Context) {
	books, err := h.bookService.ViewBook()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var booksResponse []book.BookResponse

	for _, response := range books {
		bookResponse := convertToBookResponse(response)

		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func (h *bookHandeler) GetBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	response, err := h.bookService.ViewbookById(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := convertToBookResponse(response)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandeler) CreateBook(c *gin.Context) {
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

func convertToBookResponse(response book.Book) book.BookResponse {
	return book.BookResponse{
		Id:          response.Id,
		Title:       response.Title,
		Price:       response.Price,
		Description: response.Description,
		Rating:      response.Rating,
	}
}
