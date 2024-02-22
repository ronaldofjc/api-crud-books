package book

import (
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
