package book

type Service interface {
	ViewBook() ([]Book, error)
	ViewbookById(Id int) (Book, error)
	InsertBook(bookRequest BookRequest) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) ViewBook() ([]Book, error) {
	books, err := s.repository.ViewBook()
	if err != nil {
		return books, err
	}

	return books, nil
}

func (s *service) ViewbookById(Id int) (Book, error) {
	book, err := s.repository.ViewbookById(Id)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (s *service) InsertBook(bookRequest BookRequest) (Book, error) {
	id, _ := bookRequest.Id.Int64()
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()

	book := Book{
		Id:          int(id),
		Title:       bookRequest.Title,
		Description: bookRequest.Description,
		Price:       int(price),
		Rating:      int(rating),
	}

	newBook, err := s.repository.InsertBook(book)
	return newBook, err
}
