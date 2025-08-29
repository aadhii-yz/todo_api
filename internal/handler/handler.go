package handler

import (
	"github.com/aadhii-yz/todo_api/internal/db"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	DB db.TodoStorage
}

func NewHandler(db db.TodoStorage) *Handler {
	return &Handler{
		DB: db,
	}
}

func (h *Handler) GetAllTodos(ctx *gin.Context) {}
func (h *Handler) GetTodoById(ctx *gin.Context) {}
func (h *Handler) PostTodo(ctx *gin.Context)    {}
func (h *Handler) PatchTodo(ctx *gin.Context)   {}
func (h *Handler) DeleteTodo(ctx *gin.Context)  {}
