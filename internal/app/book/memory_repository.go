package book

import (
	"errors"
	"github.com/google/uuid"
	"main/internal/domain"
)

type MemoryRepository struct {
	books map[string]domain.Book
}

func NewMemoryRepository(books map[string]domain.Book) *MemoryRepository {
	return &MemoryRepository{
		books: books,
	}
}

func (repo *MemoryRepository) RemoveByIdRepo(id string) error {
	if _, ok := repo.books[id]; ok {
		delete(repo.books, id)
		return nil
	}

	return errors.New("book not found")
}

func (repo *MemoryRepository) GetByIdRepo(id string) (*domain.Book, error) {
	if book, ok := repo.books[id]; ok {
		return &book, nil
	}

	return &domain.Book{}, nil
}

func (repo *MemoryRepository) CreateBookRepo(book domain.Book) (*domain.Book, error) {
	id := uuid.New()
	book.ID = id.String()
	repo.books[id.String()] = book
	return &book, nil
}

func (repo *MemoryRepository) GetBooksRepo() ([]domain.Book, error) {
	var books = make([]domain.Book, 0)
	for _, value := range repo.books {
		books = append(books, value)
	}

	return books, nil
}

func (repo *MemoryRepository) UpdateByIdRepo(id string, book domain.Book) (*domain.Book, error) {
	if _, ok := repo.books[id]; ok {
		book.ID = id
		repo.books[id] = book
		return &book, nil
	}

	return nil, errors.New("book not found")
}
