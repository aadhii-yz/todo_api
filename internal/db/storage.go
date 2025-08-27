package db

import "github.com/aadhii-yz/todo_api/internal/todo"

type TodoStorage interface {
	GetAll() todo.Todos
	GetById(id int) todo.Todo
	Post(todo.Todo) error
	Patch(id int, todo todo.Todo) error
	DeleteById(id int) error
}
