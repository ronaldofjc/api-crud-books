package book

import (
	"errors"
	"github.com/google/uuid"
	"main/internal/domain"
)

type JsonRepository struct {
	books map[string]domain.Book
}

func NewJsonRepository(books map[string]domain.Book) *JsonRepository {
	return &JsonRepository{
		books: books,
	}
}

func (repo *JsonRepository) RemoveByIdRepo(id string) error {
	if _, ok := repo.books[id]; ok {
		delete(repo.books, id)
		return nil
	}

	return errors.New("book not found")
}

func (repo *JsonRepository) GetByIdRepo(id string) (*domain.Book, error) {
	if book, ok := repo.books[id]; ok {
		return &book, nil
	}

	return &domain.Book{}, nil
}

func (repo *JsonRepository) CreateBookRepo(book domain.Book) (*domain.Book, error) {
	id := uuid.New()
	book.ID = id.String()
	repo.books[id.String()] = book
	return &book, nil
}

func (repo *JsonRepository) GetBooksRepo() ([]domain.Book, error) {
	var books = make([]domain.Book, 0)
	for _, value := range repo.books {
		books = append(books, value)
	}

	return books, nil
}

func (repo *JsonRepository) UpdateByIdRepo(id string, book domain.Book) (*domain.Book, error) {
	if _, ok := repo.books[id]; ok {
		book.ID = id
		repo.books[id] = book
		return &book, nil
	}

	return nil, errors.New("book not found")
}
