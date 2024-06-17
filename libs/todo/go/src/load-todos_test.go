package src

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestLoadTodos(t *testing.T) {
	todos, err := LoadTodos("../fixtures/test1.json")
	if err != nil {
		t.Fatalf("An error occured in LoadTodos: %s\n", err.Error())
	}

	expected := []TodoEntry{
		{
			Name:        "thing 1",
			Description: "do thing 1",
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
	if !reflect.DeepEqual(todos, expected) {
		expected, _ := json.Marshal(expected)
		todos, _ := json.Marshal(todos)

		t.Fatalf("Expected >>>>>\n%s\nGot      <<<<<\n%s\n", string(expected), string(todos))
	}
}
