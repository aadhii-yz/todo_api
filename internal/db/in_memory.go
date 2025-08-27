package db

import "github.com/aadhii-yz/todo_api/internal/todo"

// type InMemoryDB todo.Todos
type InMemoryDB struct {
	data map[int]todo.Todo
}

func CreateInMemoryDB() *InMemoryDB {
	db := InMemoryDB{}
	db.data = make(map[int]todo.Todo)
	return &db
}

func (db *InMemoryDB) GetAll() todo.Todos {
	return todo.Todos{}
}

func (db *InMemoryDB) GetById(id int) todo.Todo {
	return todo.Todo{}
}

func (db *InMemoryDB) Post(todo.Todo) error { return nil }

func (db *InMemoryDB) Patch(id int, todo todo.Todo) error { return nil }

func (db *InMemoryDB) DeleteById(id int) error { return nil }
