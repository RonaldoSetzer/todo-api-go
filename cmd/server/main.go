package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

const (
	DO    = "do"
	DONE  = "done"
	DOING = "doing"
)

var todos = []Todo{
	{1, "Study Golang", "Study Golang everyday", DO},
	{2, "Study React", "Study React everyday", DO},
	{3, "Study Vue", "Study Vue everyday", DO},
	{4, "Study Flutter", "Study Flutter everyday", DO},
}

func main() {
	http.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		segments := strings.Split(path, "/")

		switch r.Method {
		case http.MethodGet:
			if len(segments) < 3 || segments[2] == "" {
				handleGetTodos(w)
			} else {
				id, err := strconv.Atoi(segments[2])
				if err != nil {
					http.Error(w, "Invalid ID", http.StatusBadRequest)
					return
				}
				handleGetTodo(id, w)
			}

		case http.MethodPost:
			var todo Todo
			if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			handleAddTodo(todo, w)

		case http.MethodPut:
			if len(segments) < 3 || segments[2] == "" {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			} else {
				id, err := strconv.Atoi(segments[2])
				if err != nil {
					http.Error(w, "Invalid ID", http.StatusBadRequest)
					return
				}
				var todo Todo
				if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				todo.ID = id
				handleUpdateTodo(todo, w)
			}

		case http.MethodDelete:
			if len(segments) < 3 || segments[2] == "" {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			} else {
				id, err := strconv.Atoi(segments[2])
        if err != nil {
          http.Error(w, "Invalid ID", http.StatusBadRequest)
          return
        }
				handleDeleteTodo(id, w)
			}
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.ListenAndServe(":8080", nil)
}

func handleUpdateTodo(todo Todo, w http.ResponseWriter) {
	fmt.Fprintf(w, "Update Todo %s", todo.Title)
	for i, t := range todos {
		if t.ID == todo.ID {
			todos[i] = todo
			break
		}
	}
}

func handleDeleteTodo(id int, w http.ResponseWriter) {
	fmt.Fprintf(w, "Delete Todo %d", id)
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}
}

func handleAddTodo(todo Todo, w http.ResponseWriter) {
	fmt.Fprintf(w, "Add Todo %s", todo.Title)
	todo.ID = len(todos) + 1
	todos = append(todos, todo)
}

func handleGetTodo(id int, w http.ResponseWriter) {
	fmt.Fprintf(w, "Todo %d", id)

	var todo Todo
	for _, t := range todos {
		if t.ID == id {
			todo = t
			break
		}
	}

	fmt.Fprintf(w, "%d. %s - %s - %s\n", todo.ID, todo.Title, todo.Description, todo.Status)
}

func handleGetTodos(w http.ResponseWriter) {
	fmt.Fprintln(w, "Todo List")
	for _, todo := range todos {
		fmt.Fprintf(w, "%d. %s - %s - %s\n", todo.ID, todo.Title, todo.Description, todo.Status)
	}
}
