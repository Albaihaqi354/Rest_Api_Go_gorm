package book

import (
	"gorm.io/gorm"
)

type Repository interface {
	ViewBook() ([]Book, error)
	ViewbookById(Id int) (Book, error)
	InsertBook(book Book) (Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) ViewBook() ([]Book, error) {
	var books []Book
	err := r.db.Find(&books).Error
	if err != nil {
		return books, err
	}
	return books, nil
}

func (r *repository) ViewbookById(Id int) (Book, error) {
	var book Book
	err := r.db.Find(&book, Id).Error
	if err != nil {
		return book, err
	}
	return book, nil
}

func (r *repository) InsertBook(book Book) (Book, error) {
	err := r.db.Create(&book).Error
	if err != nil {
		return book, err
	}
	return book, nil
}
