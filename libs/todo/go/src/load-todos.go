package src

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type TodoEntry struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func LoadTodos(path string) ([]TodoEntry, error) {
	if !strings.HasSuffix(path, ".json") {
		return nil, errors.New("The todo-list file must be a JSON file")
	}
	if file, err := os.ReadFile(path); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []TodoEntry{}, nil
		}
		return nil, fmt.Errorf("Failed to read file %s: %s", path, err)
	} else {
		arr := []TodoEntry{}
		if err = json.Unmarshal(file, &arr); err != nil {
			return nil, fmt.Errorf("Failed to parse todos: %s", err)
		}
		for i, a := range arr {
			for j, b := range arr {
				if i != j && a.Name == b.Name {
					return nil, errors.New("Duplicate names in todo-list file")
				}
			}
		}
		return arr, nil
	}
}
