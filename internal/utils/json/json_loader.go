package utils

import (
	"encoding/json"
	"log"
	"main/internal/domain"
	"os"
)

func ReadJsonFileListBooks() []domain.Book {
	file, err := os.Open("internal/app/book/books.json")
	if err != nil {
		log.Fatal("error on read json file. : ", err.Error())
	}

	defer file.Close()
	books := make([]domain.Book, 0)
	decoder := json.NewDecoder(file)
	data := domain.Book{}
	decoder.Token()
	for decoder.More() {
		decoder.Decode(&data)
		books = append(books, data)
	}

	return books
}

func ReadJsonFileBooks() map[string]domain.Book {
	file, err := os.Open("internal/app/book/books.json")
	if err != nil {
		log.Fatal("error on read json file. : ", err.Error())
	}

	defer file.Close()
	books := make(map[string]domain.Book)
	decoder := json.NewDecoder(file)
	data := domain.Book{}
	decoder.Token()
	for decoder.More() {
		decoder.Decode(&data)
		books[data.ID] = data
	}

	return books
}
