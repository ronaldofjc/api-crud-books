package book

import (
	"errors"
	"github.com/google/uuid"
	"main/internal/domain"
)

type IRepository interface {
	CreateBookRepo(book domain.Book) (*domain.Book, error)
	GetBooksRepo() ([]domain.Book, error)
	GetByIdRepo(id string) (*domain.Book, error)
	RemoveByIdRepo(id string) error
	UpdateByIdRepo(id string, book domain.Book) (*domain.Book, error)
}

type Repository struct {
	books map[string]domain.Book
}

func NewRepository(books map[string]domain.Book) *Repository {
	return &Repository{
		books: books,
	}
}

func (repo *Repository) RemoveByIdRepo(id string) error {
	if _, ok := repo.books[id]; ok {
		delete(repo.books, id)
		return nil
	}

	return errors.New("book not found")
}

func (repo *Repository) GetByIdRepo(id string) (*domain.Book, error) {
	if book, ok := repo.books[id]; ok {
		return &book, nil
	}

	return &domain.Book{}, nil
}

func (repo *Repository) CreateBookRepo(book domain.Book) (*domain.Book, error) {
	id := uuid.New()
	book.ID = id.String()
	repo.books[id.String()] = book
	return &book, nil
}

func (repo *Repository) GetBooksRepo() ([]domain.Book, error) {
	var books = make([]domain.Book, 0)
	for _, value := range repo.books {
		books = append(books, value)
	}

	return books, nil
}

func (repo *Repository) UpdateByIdRepo(id string, book domain.Book) (*domain.Book, error) {
	if _, ok := repo.books[id]; ok {
		book.ID = id
		repo.books[id] = book
		return &book, nil
	}

	return nil, errors.New("book not found")
}
