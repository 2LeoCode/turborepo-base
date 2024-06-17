package src

import (
	"encoding/json"
	"fmt"
	"os"
)

type TodoList struct {
	GetTodos   func() []TodoEntry
	AddTodo    func(string, string) bool
	RemoveTodo func(string) bool
	SaveTodos  func() error
}

func NewTodoList(path string) (*TodoList, error) {
	if todos, err := LoadTodos(path); err != nil {
		return nil, err
	} else {
		return &TodoList{
			GetTodos: func() []TodoEntry {
				return todos
			},
			AddTodo: func(name string, description string) bool {
				for _, todo := range todos {
					if todo.Name == name {
						return false
					}
				}
				todos = append(todos, TodoEntry{
					Name:        name,
					Description: description,
				})
				return true
			},
			RemoveTodo: func(name string) bool {
				for i, todo := range todos {
					if todo.Name == name {
						todos[i] = todos[len(todos)-1]
						todos = todos[:len(todos)-1]
						return true
					}
				}
				return false
			},
			SaveTodos: func() error {
				if encoded, err := json.Marshal(todos); err != nil {
					return fmt.Errorf("Error while encoding todos: %s", err)
				} else if err = os.WriteFile(path, encoded, 0644); err != nil {
					return fmt.Errorf("Error while saving todos: %s", err)
				}
				return nil
			},
		}, nil
	}
}
