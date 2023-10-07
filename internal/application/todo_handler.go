package application

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/RonaldoSetzer/todo-api-go/internal/domain"
	"github.com/google/uuid"
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
	var todo domain.TodoDTO
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&todo); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdTodo, err := h.addTodoUseCase.Execute(todo.Title, todo.Description)
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
	path := r.URL.Path
	segments := strings.Split(path, "/")
	todoID, err := uuid.Parse(segments[len(segments)-1])
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	var todoDto domain.TodoDTO
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&todoDto); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updateTodo, err := h.updateTodoUseCase.Execute(todoID,todoDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateTodo)
}

func (h *TodoHandler) HandleDeleteTodoRequest(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	segments := strings.Split(path, "/")
	todoID, err := uuid.Parse(segments[len(segments)-1])
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
	todoID, err := uuid.Parse(segments[len(segments)-1])
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
