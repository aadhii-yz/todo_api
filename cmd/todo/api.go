package main

import (
	"github.com/aadhii-yz/todo_api/internal/db"
	"github.com/gin-gonic/gin"
)

func main() {
	s := Server{
		Port:   6060,
		DB:     db.CreateInMemoryDB(),
		Router: gin.Default(),
	}
	s.Setup()
	s.Run()
}
