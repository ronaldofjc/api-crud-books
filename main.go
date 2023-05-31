package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/internal/server"
)

func main() {
	router := gin.Default()
	handler := buildDependencies()
	server.Routes(&router.RouterGroup, handler)
	if err := router.Run(":8090"); err != nil {
		log.Fatal(err)
	}
}
