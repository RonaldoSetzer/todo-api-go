package infrastructure

import (
	"github.com/RonaldoSetzer/todo-api-go/internal/domain"
	"github.com/RonaldoSetzer/todo-api-go/internal/errors"
	"github.com/google/uuid"
)

type LocalRepository struct {
	todos  map[uuid.UUID]domain.Todo
}

func NewLocalRepository() *LocalRepository {
	return &LocalRepository{
		todos:  make(map[uuid.UUID]domain.Todo),
	}
}

func (r *LocalRepository) GetTodos() ([]domain.Todo, error) {
	todos := make([]domain.Todo, 0, len(r.todos))
	for _, todo := range r.todos {
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *LocalRepository) GetTodoById(id uuid.UUID) (domain.Todo, error) {
	todo, ok := r.todos[id]
	if !ok {
		return domain.Todo{}, errors.TodoNotFoundError{ID: id}
	}
	return todo, nil
}

func (r *LocalRepository) AddTodo(newTodo domain.Todo) (domain.Todo, error) {
  r.todos[newTodo.ID] = newTodo
  return newTodo, nil
}

func (r *LocalRepository) UpdateTodo(updatedTodo domain.Todo) (domain.Todo, error) {
	_, ok := r.todos[updatedTodo.ID]
	if !ok {
		return domain.Todo{}, errors.TodoNotFoundError{ID: updatedTodo.ID}
	}
	r.todos[updatedTodo.ID] = updatedTodo
	return updatedTodo, nil
}

func (r *LocalRepository) DeleteTodo(id uuid.UUID) (domain.Todo, error) {
	todo, ok := r.todos[id]
	if !ok {
		return domain.Todo{}, errors.TodoNotFoundError{ID: id}
	}
  delete(r.todos, id)
	return todo, nil
}
