package server

import "github.com/gin-gonic/gin"

func Routes(router *gin.RouterGroup, handler Handler) {
	router.GET("/", HelloHandler)
	router.GET("/ping", PingHandler)
	router.POST("/books", handler.CreateBook)
	router.GET("/books", handler.GetBooks)
	router.GET("books/:id", handler.GetById)
	router.DELETE("/books/:id", handler.RemoveById)
	router.PUT("/books/:id", handler.UpdateById)
}
