package book

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"main/internal/domain"
	utils "main/internal/utils/json"
	"os"
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
	books := utils.ReadJsonFileListBooks()
	for index, data := range books {
		if data.ID == id {
			books = append(books[:index], books[index+1:]...)
			jsonString, _ := json.Marshal(books)
			err := os.WriteFile("internal/app/book/books.json", jsonString, 0644)
			if err != nil {
				return errors.New("error on remove book on repository")
			}
			return nil
		}
	}

	return errors.New("book not found")
}

func (repo *JsonRepository) GetByIdRepo(id string) (*domain.Book, error) {
	books := utils.ReadJsonFileListBooks()
	for _, data := range books {
		if data.ID == id {
			return &data, nil
		}
	}

	return &domain.Book{}, nil
}

func (repo *JsonRepository) CreateBookRepo(book domain.Book) (*domain.Book, error) {
	books := utils.ReadJsonFileListBooks()
	id := uuid.New()
	book.ID = id.String()
	books = append(books, book)
	jsonString, _ := json.Marshal(books)
	err := os.WriteFile("internal/app/book/books.json", jsonString, 0644)
	if err != nil {
		return nil, errors.New("error on save new book on repository")
	}

	return &book, nil
}

func (repo *JsonRepository) GetBooksRepo() ([]domain.Book, error) {
	books := utils.ReadJsonFileListBooks()
	return books, nil
}

func (repo *JsonRepository) UpdateByIdRepo(id string, book domain.Book) (*domain.Book, error) {
	books := utils.ReadJsonFileListBooks()
	for index, data := range books {
		if data.ID == id {
			book.ID = id
			books = append(books[:index], books[index+1:]...)
			books = append(books, book)
			jsonString, _ := json.Marshal(books)
			err := os.WriteFile("internal/app/book/books.json", jsonString, 0644)
			if err != nil {
				return nil, errors.New("error on remove book on repository")
			}
			return &book, nil
		}
	}

	return nil, errors.New("book not found")
}
