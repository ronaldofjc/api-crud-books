package main

import (
	"main/internal/app/book"
	"main/internal/domain"
	"main/internal/server"
)

func buildDependencies() server.Handler {
	books := map[string]domain.Book{}
	repo := book.NewRepository(books)
	service := book.NewService(repo)
	return server.NewHandler(service)
}
