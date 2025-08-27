package main

import (
	"fmt"
	"log"

	"github.com/aadhii-yz/todo_api/internal/db"
	"github.com/aadhii-yz/todo_api/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	s := Server{
		Port: 6060,
		DB:   db.InMemoryDB{},
	}
	addr := fmt.Sprintf(":%d", s.Port)
	router := gin.Default()

	router.GET("/todos", handler.GetAllTodos)
	router.GET("/todos/:id", handler.GetTodoById)
	router.POST("/todos", handler.PostTodo)
	router.PATCH("/todos/:id", handler.PatchTodo)
	router.DELETE("/todos/:id", handler.DeleteTodo)

	log.Fatal(router.Run(addr))
}
