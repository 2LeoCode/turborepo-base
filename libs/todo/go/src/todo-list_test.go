package src

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func TestTodoList(t *testing.T) {
	const todosPath = "../fixtures/test1.json"
	if todoList, err := NewTodoList(todosPath); err != nil {
		t.Fatalf("An error occured while creating a TodoList: %s\n", err.Error())
	} else {
		backup := todoList.GetTodos()
		t.Cleanup(func() {
			encoded, _ := json.Marshal(backup)
			os.WriteFile(todosPath, encoded, 0644)
		})
		if todoList.AddTodo("something", "something to do") == false {
			t.Fatal("AddTodo returned false, expected true")
		}
		if todoList.AddTodo("something", "this entry has a duplicate name") == true {
			t.Fatal("AddTodo returned true, expected false")
		}
		if todoList.RemoveTodo("thing 1") == false {
			t.Fatal("RemoveTodo returned false, expected true")
		}
		if todoList.RemoveTodo("noexist") == true {
			t.Fatal("RemoveTodo returned true, expected false")
		}
		expected := []TodoEntry{
			{
				Name:        "something",
				Description: "something to do",
			},
			{
				Name:        "thing 2",
				Description: "do thing 2",
			},
			{
				Name:        "thing 3",
				Description: "do thing 3",
			},
		}
		todos := todoList.GetTodos()
		if !reflect.DeepEqual(todos, expected) {
			expected, _ := json.Marshal(expected)
			todos, _ := json.Marshal(todos)

			t.Fatalf("Expected >>>>>\n%s\nGot      <<<<<\n%s\n", string(expected), string(todos))
		}
		if err = todoList.SaveTodos(); err != nil {
			t.Fatalf("SaveTodos failed %s\n", err.Error())
		}
		if content, err := os.ReadFile(todosPath); err != nil {
			t.Fatalf("Unexpected error: cannot read file: %s", err.Error())
		} else {
			var parsed []TodoEntry
			if err = json.Unmarshal(content, &parsed); err != nil {
				t.Fatalf("Failed to parse todos file: %s\n", err.Error())
			}
			if !reflect.DeepEqual(todos, expected) {
				expected, _ := json.Marshal(expected)
				todos, _ := json.Marshal(todos)

				t.Fatalf("Expected >>>>>\n%sGot      <<<<<\n%s\n", string(expected), string(todos))
			}
		}
	}
}
