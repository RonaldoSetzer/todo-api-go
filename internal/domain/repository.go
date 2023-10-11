package domain

import "github.com/google/uuid"

type Repository interface {
	GetTodos() ([]TodoDTO, error)
	GetTodoById(id uuid.UUID) (TodoDTO, error)
	AddTodo(todo Todo) (TodoDTO, error)
	UpdateTodo(todo Todo) (TodoDTO, error)
	DeleteTodo(id uuid.UUID) (TodoDTO, error)
}
