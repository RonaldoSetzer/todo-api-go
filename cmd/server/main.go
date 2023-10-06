package main

import (
	"net/http"

	"github.com/RonaldoSetzer/todo-api-go/internal/application"
	"github.com/RonaldoSetzer/todo-api-go/internal/domain"
	"github.com/RonaldoSetzer/todo-api-go/internal/infrastructure"
	"github.com/gorilla/mux"
)

func main() {
	localRepository := infrastructure.NewLocalRepository()
	localRepository.AddTodo(domain.Todo{ID: 1, Title: "Todo 1", Description: "Description 1", Status: "DO"})
	localRepository.AddTodo(domain.Todo{ID: 2, Title: "Todo 2", Description: "Description 2", Status: "DO"})
	localRepository.AddTodo(domain.Todo{ID: 3, Title: "Todo 3", Description: "Description 3", Status: "DO"})

	addTodoUseCase := application.NewAddTodoUseCase(localRepository)
	getTodosUseCase := application.NewGetTodosUseCase(localRepository)
  getTodoUseCase := application.NewGetTodoUseCase(localRepository)
	updateTodoUseCase := application.NewUpdateTodoUseCase(localRepository)
	deleteTodoUseCase := application.NewDeleteTodoUseCase(localRepository)

	router := mux.NewRouter()

	todoHandler := application.NewTodoHandler(addTodoUseCase, getTodosUseCase, getTodoUseCase, updateTodoUseCase, deleteTodoUseCase)
	router.HandleFunc("/todos", todoHandler.HandleAddTodoRequest).Methods("POST")
	router.HandleFunc("/todos", todoHandler.HandleGetTodosRequest).Methods("GET")
  router.HandleFunc("/todos/{id}", todoHandler.HandleGetTodoRequest).Methods("GET")
	router.HandleFunc("/todos/{id}", todoHandler.HandleUpdateTodoRequest).Methods("PUT")
	router.HandleFunc("/todos/{id}", todoHandler.HandleDeleteTodoRequest).Methods("DELETE")

	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
}
