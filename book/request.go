package book

import "encoding/json"

type BookRequest struct {
	Id          json.Number `json:"id" binding:"required,number"`
	Title       string      `json:"title" binding:"required"`
	Description string      `json:"description" binding:"required"`
	Price       json.Number `json:"price" binding:"required,number"`
	Rating      json.Number `json:"rating" binding:"required,number"`
}
