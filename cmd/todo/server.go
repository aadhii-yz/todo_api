package main

import "github.com/aadhii-yz/todo_api/internal/db"

type Server struct {
	Port int
	DB   db.TodoStorage
}
