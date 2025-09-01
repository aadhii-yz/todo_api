package handler

import (
	"net/http"
	"strconv"

	"github.com/aadhii-yz/todo_api/internal/db"
	"github.com/aadhii-yz/todo_api/internal/todo"
	"github.com/gin-gonic/gin"
)

func HandleErr(err error, ctx *gin.Context, statusCode int) {
	ctx.JSON(statusCode, gin.H{
		"error": err.Error(),
	})
}

type Handler struct {
	DB db.TodoStorage
}

func NewHandler(db db.TodoStorage) *Handler {
	return &Handler{
		DB: db,
	}
}

// GET /todos
func (h *Handler) GetAllTodos(ctx *gin.Context) {
	todos := h.DB.GetAll()
	ctx.JSON(http.StatusOK, todos)
}

// GET /todos/:id
func (h *Handler) GetTodoById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		HandleErr(err, ctx, http.StatusInternalServerError)
		return
	}

	t, err := h.DB.GetById(id)
	if err != nil {
		HandleErr(err, ctx, http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, t)
}

// POST /todos
func (h *Handler) PostTodo(ctx *gin.Context) {
	var t todo.Todo
	if err := ctx.BindJSON(&t); err != nil {
		HandleErr(err, ctx, http.StatusBadRequest)
		return
	}

	id := h.DB.Post(t)
	t.Id = id
	ctx.JSON(http.StatusOK, t)
}

// PATCH /todos/:id
func (h *Handler) PatchTodo(ctx *gin.Context) {
	var t todo.Todo

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		HandleErr(err, ctx, http.StatusInternalServerError)
		return
	}

	if err := ctx.BindJSON(&t); err != nil {
		HandleErr(err, ctx, http.StatusBadRequest)
		return
	}

	if _, err := h.DB.GetById(id); err != nil {
		HandleErr(err, ctx, http.StatusNotFound)
		return
	}

	if err := h.DB.Patch(id, t); err != nil {
		HandleErr(err, ctx, http.StatusInternalServerError)
		return
	}

	t.Id = id
	ctx.JSON(http.StatusOK, t)
}

// DELETE /todos/:id
func (h *Handler) DeleteTodo(ctx *gin.Context) {
}
