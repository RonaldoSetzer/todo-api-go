package application

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/RonaldoSetzer/todo-api-go/internal/domain"
)

type TodoHandler struct {
	addTodoUseCase    *AddTodoUseCase
	getTodosUseCase   *GetTodosUseCase
	getTodoUseCase    *GetTodoUseCase
	updateTodoUseCase *UpdateTodoUseCase
	deleteTodoUseCase *DeleteTodoUseCase
}

func NewTodoHandler(
	addTodoUseCase *AddTodoUseCase,
	getTodosUseCase *GetTodosUseCase,
	getTodoUseCase *GetTodoUseCase,
	updateTodoUseCase *UpdateTodoUseCase,
	deleteTodoUseCase *DeleteTodoUseCase,
) *TodoHandler {
	return &TodoHandler{
		addTodoUseCase:    addTodoUseCase,
		getTodosUseCase:   getTodosUseCase,
		getTodoUseCase:    getTodoUseCase,
		updateTodoUseCase: updateTodoUseCase,
		deleteTodoUseCase: deleteTodoUseCase,
	}
}

func (h *TodoHandler) HandleAddTodoRequest(w http.ResponseWriter, r *http.Request) {
	var newTodo domain.Todo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newTodo); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdTodo, err := h.addTodoUseCase.Execute(newTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdTodo)
}

func (h *TodoHandler) HandleGetTodosRequest(w http.ResponseWriter, r *http.Request) {
	todos, err := h.getTodosUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func (h *TodoHandler) HandleUpdateTodoRequest(w http.ResponseWriter, r *http.Request) {
	var updatedTodo domain.Todo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedTodo); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	todo, err := h.updateTodoUseCase.Execute(updatedTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) HandleDeleteTodoRequest(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	segments := strings.Split(path, "/")
	todoID, err := strconv.Atoi(segments[len(segments)-1])

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	todo, err := h.deleteTodoUseCase.Execute(todoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) HandleGetTodoRequest(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	segments := strings.Split(path, "/")
	todoID, err := strconv.Atoi(segments[len(segments)-1])

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	todo, err := h.getTodoUseCase.Execute(todoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
