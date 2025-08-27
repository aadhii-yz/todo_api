package db

import (
	"errors"

	"github.com/aadhii-yz/todo_api/internal/todo"
)

type InMemoryDB struct {
	Data   map[int]todo.Todo
	NextID int
}

func CreateInMemoryDB() *InMemoryDB {
	db := InMemoryDB{}
	db.Data = make(map[int]todo.Todo)
	return &db
}

func (db *InMemoryDB) GetAll() todo.Todos {
	todos := todo.Todos{}
	for _, v := range db.Data {
		todos = append(todos, v)
	}

	return todos
}

func (db *InMemoryDB) GetById(id int) (todo.Todo, error) {
	todo, ok := db.Data[id]
	if !ok {
		return todo, errors.New("Todo Not Found")
	}
	return todo, nil
}

func (db *InMemoryDB) Post(todo todo.Todo) int {
	todo.Id = db.NextID
	db.Data[db.NextID] = todo
	db.NextID++
	return db.NextID - 1
}

func (db *InMemoryDB) Patch(id int, todo todo.Todo) error {
	if _, ok := db.Data[id]; !ok {
		return errors.New("Todo Not Found")
	}
	todo.Id = id
	db.Data[id] = todo
	return nil
}

func (db *InMemoryDB) DeleteById(id int) error {
	if _, ok := db.Data[id]; !ok {
		return errors.New("Todo Not Found")
	}
	delete(db.Data, id)
	return nil
}
