package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/config/middleware"
	"main/internal/app/book"
	"main/internal/domain"
	"main/internal/server"
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
	books := map[string]domain.Book{}
	repo := book.NewRepository(books)
	service := book.NewService(repo)
	return server.NewHandler(service)
}
