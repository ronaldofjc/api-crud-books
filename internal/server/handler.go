package server

import (
	"github.com/gin-gonic/gin"
	"main/internal/app/book"
	"main/internal/domain"
	"net/http"
)

type Handler struct {
	service book.IService
}

func NewHandler(s book.IService) Handler {
	return Handler{
		service: s,
	}
}

func HelloHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "API Golang it works!"})
}

func PingHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func (handler Handler) CreateBook(ctx *gin.Context) {
	var request domain.Book

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid params"})
		return
	}

	response, err := handler.service.CreateBook(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (handler Handler) GetBooks(ctx *gin.Context) {
	response, err := handler.service.GetBooks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (handler Handler) GetById(ctx *gin.Context) {
	bookId := ctx.Param("id")
	response, err := handler.service.GetById(bookId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if response.IsEmpty() {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (handler Handler) RemoveById(ctx *gin.Context) {
	bookId := ctx.Param("id")
	err := handler.service.RemoveById(bookId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (handler Handler) UpdateById(ctx *gin.Context) {
	bookId := ctx.Param("id")
	var request domain.Book

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid params"})
		return
	}

	response, err := handler.service.UpdateById(bookId, request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
