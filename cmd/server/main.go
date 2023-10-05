package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/RonaldoSetzer/todo-api-go/internal/domain"
	"github.com/RonaldoSetzer/todo-api-go/internal/infrastructure"
)

func main() {
	localRepository := infrastructure.NewLocalRepository()
  localRepository.AddTodo(
    domain.Todo{ Title: "First Todo", Description: "This is the first todo", Status: domain.DO, },
  )
  localRepository.AddTodo(
    domain.Todo{ Title: "Second Todo", Description: "This is the second todo", Status: domain.DO, },
  )
  localRepository.AddTodo(
    domain.Todo{ Title: "Third Todo", Description: "This is the third todo", Status: domain.DO, },
  )

	http.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		segments := strings.Split(path, "/")

		switch r.Method {
		case http.MethodGet:
			if len(segments) < 3 || segments[2] == "" {
        todos, err := localRepository.GetTodos()
        if err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
        }
				printTodos(todos, "GetTodos", w)
			} else {
				id, err := strconv.Atoi(segments[2])
				if err != nil {
					http.Error(w, "Invalid ID", http.StatusBadRequest)
					return
				}
				todo, err := localRepository.GetTodoById(id)
        if err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
        }
        printTodos([]domain.Todo{todo}, "GetTodoById", w)
			}

		case http.MethodPost:
			var newTodo domain.Todo
			if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
      todo, err := localRepository.AddTodo(newTodo)
			if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
      }
      printTodos([]domain.Todo{todo}, "AddTodo", w)

		case http.MethodPut:
			if len(segments) < 3 || segments[2] == "" {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			} else {
				id, err := strconv.Atoi(segments[2])
				if err != nil {
					http.Error(w, "Invalid ID", http.StatusBadRequest)
					return
				}
				var newTodo domain.Todo
				if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
        newTodo.ID = id
        todo, err := localRepository.UpdateTodo(newTodo)
        if err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
        }
        printTodos([]domain.Todo{todo}, "UpdateTodo", w)
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
				todo, err := localRepository.DeleteTodo(id)
        if err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
        }
        printTodos([]domain.Todo{todo}, "DeleteTodo", w)
			}
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.ListenAndServe(":8080", nil)
}

func printTodos(todos []domain.Todo, title string, w http.ResponseWriter) {
	fmt.Fprintln(w, title)
	for _, todo := range todos {
		fmt.Fprintf(w, "%d. %s - %s - %s\n", todo.ID, todo.Title, todo.Description, todo.Status)
	}
}
