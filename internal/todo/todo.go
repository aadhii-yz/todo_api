package todo

import (
	"encoding/json"
	"fmt"
)

type Todo struct {
	Id        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"done"`
}

type Todos []Todo

func CreateTodo(task string) Todo {
	return Todo{Task: task}
}

func (t *Todo) SetStatus(completed bool) {
	t.Completed = completed
}

func (t *Todo) SetTask(task string) {
	t.Task = task
}

func (t *Todo) Print() {
	fmt.Printf("Id: %d\nTask: %q\nDone: %t", t.Id, t.Task, t.Completed)
}

func (t *Todo) ToJSON() ([]byte, error) {
	return json.Marshal(t)
}

func (t *Todo) FromJSON(b []byte) error {
	return json.Unmarshal(b, t)
}
