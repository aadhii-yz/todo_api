package handler

import (
	"net/http"
	"strconv"

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

func (h *Handler) GetAllTodos(ctx *gin.Context) {
	todos := h.DB.GetAll()

	todosJSON, err := todos.ToJSON()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"todos": string(todosJSON),
	})
}

func (h *Handler) GetTodoById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	todo, err := h.DB.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	todoJSON, err := todo.ToJSON()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"todo": string(todoJSON)})
}

func (h *Handler) PostTodo(ctx *gin.Context)   {}
func (h *Handler) PatchTodo(ctx *gin.Context)  {}
func (h *Handler) DeleteTodo(ctx *gin.Context) {}
