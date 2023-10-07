package domain

import "github.com/google/uuid"

type Repository interface {
	GetTodos() ([]Todo, error)
	GetTodoById(id uuid.UUID) (Todo, error)
	AddTodo(todo Todo) (Todo, error)
	UpdateTodo(todo Todo) (Todo, error)
	DeleteTodo(id uuid.UUID) (Todo, error)
}
