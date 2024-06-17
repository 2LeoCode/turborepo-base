package main

import (
	"encoding/json"
	"libs-todo-go"
	"log"
	"net/http"
)

func main() {
	if todoList, err := libs_todo_go.NewTodoList("./todo-list.json"); err != nil {
		log.Fatalf("Failed to create todo-list: %s", err)
	} else {
		http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
			var body []string
			if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
				http.Error(response, "Bad request", http.StatusBadRequest)
			} else {
				switch request.Method {
				case "POST":
					if len(body) == 0 {
						http.Error(response, "Bad request", http.StatusBadRequest)
						break
					}
					encoder := json.NewEncoder(response)
					switch body[0] {
					case "GetTodos":
						if len(body) != 1 {
							http.Error(response, "Bad request", http.StatusBadRequest)
						} else {
							response.Header().Set("Access-Control-Allow-Origin", "*")
							response.Header().Set("Content-Type", "application/json")
							response.WriteHeader(http.StatusCreated)
							if err = encoder.Encode(todoList.GetTodos()); err != nil {
								http.Error(response, "Internal server error", http.StatusInternalServerError)
							}
						}
					case "AddTodo":
						if len(body) != 3 {
							http.Error(response, "Bad request", http.StatusBadRequest)
						} else {
							response.Header().Set("Access-Control-Allow-Origin", "*")
							response.Header().Set("Content-Type", "application/json")
							if err := encoder.Encode(todoList.AddTodo(body[1], body[2])); err != nil {
								http.Error(response, "Internal server error", http.StatusInternalServerError)
							}
						}
					case "RemoveTodo":
						if len(body) != 2 {
							http.Error(response, "Bad request", http.StatusBadRequest)
						} else {
							response.Header().Set("Access-Control-Allow-Origin", "*")
							response.Header().Set("Content-Type", "application/json")
							if err := encoder.Encode(todoList.RemoveTodo(body[1])); err != nil {
								http.Error(response, "Internal server error", http.StatusInternalServerError)
							}
						}
					case "SaveTodos":
						if len(body) != 1 {
							http.Error(response, "Bad request", http.StatusBadRequest)
						} else if err = todoList.SaveTodos(); err != nil {
							http.Error(response, "Internal server error", http.StatusInternalServerError)
						} else {
							response.Header().Set("Access-Control-Allow-Origin", "*")
							response.Header().Set("Content-Type", "application/json")
							encoder.Encode(nil)
						}
					default:
						http.Error(response, "Bad request", http.StatusBadRequest)
					}
				default:
					http.Error(response, "Bad request", http.StatusBadRequest)
				}
			}
		})
		log.Fatalln(http.ListenAndServe(":3000", nil))
	}

}
