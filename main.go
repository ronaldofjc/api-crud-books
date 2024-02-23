package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/config/middleware"
	"main/internal/app/book"
	"main/internal/domain"
	"main/internal/server"
	utils "main/internal/utils/json"
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
	repo := getRepository(false)
	service := book.NewService(repo)
	return server.NewHandler(service)
}

func getRepository(isMemoryRepo bool) book.IRepository {
	if isMemoryRepo {
		books := map[string]domain.Book{}
		return book.NewMemoryRepository(books)
	}

	books := utils.ReadJsonFileBooks()
	return book.NewJsonRepository(books)
}
