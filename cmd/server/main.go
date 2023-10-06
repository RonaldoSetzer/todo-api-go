package main

import (
	"net/http"

	"github.com/RonaldoSetzer/todo-api-go/internal/application"
	"github.com/gorilla/mux"
)

func main() {
  builder := application.NewApplicationBuilder()
  app := builder.Build()

	router := mux.NewRouter()

	router.HandleFunc("/todos", app.TodoHandler.HandleAddTodoRequest).Methods("POST")
  router.HandleFunc("/todos", app.TodoHandler.HandleGetTodosRequest).Methods("GET")
  router.HandleFunc("/todos/{id}", app.TodoHandler.HandleGetTodoRequest).Methods("GET")
  router.HandleFunc("/todos/{id}", app.TodoHandler.HandleUpdateTodoRequest).Methods("PUT")
  router.HandleFunc("/todos/{id}", app.TodoHandler.HandleDeleteTodoRequest).Methods("DELETE")

	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
}
