package main

import (
	"fmt"
	"log"

	"github.com/aadhii-yz/todo_api/internal/db"
	"github.com/aadhii-yz/todo_api/internal/handler"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Port   int
	DB     db.TodoStorage
	Router *gin.Engine
}

func (s *Server) Setup() {
	h := handler.NewHandler(s.DB)

	s.Router.GET("/todos", h.GetAllTodos)
	s.Router.GET("/todos/:id", h.GetTodoById)
	s.Router.POST("/todos", h.PostTodo)
	s.Router.PATCH("/todos/:id", h.PatchTodo)
	s.Router.DELETE("/todos/:id", h.DeleteTodo)
}

func (s *Server) Run() {
	addr := fmt.Sprintf(":%d", s.Port)
	fmt.Printf("Listening on port %s\n", addr)
	log.Fatal(s.Router.Run(addr))
}
