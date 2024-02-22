package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"main/config/middleware"
	"main/internal/app/book"
	"main/internal/domain"
	"main/internal/server"
	"os"
)

func main() {
	router := gin.Default()
	router.Use(middleware.Logger())
	handler := buildDependencies()
	server.Routes(&router.RouterGroup, handler)
	if err := router.Run(":8090"); err != nil {
		log.Fatal(err)
	}
}

func buildDependencies() server.Handler {
	books := getBooks(false)
	repo := book.NewJsonRepository(books)
	service := book.NewService(repo)
	return server.NewHandler(service)
}

func getBooks(isMemoryRepo bool) map[string]domain.Book {
	if isMemoryRepo {
		return map[string]domain.Book{}
	}
	return readJsonFile()
}

func readJsonFile() map[string]domain.Book {
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
