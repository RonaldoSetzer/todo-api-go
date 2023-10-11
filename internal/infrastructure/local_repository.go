package infrastructure

import (
	"github.com/RonaldoSetzer/todo-api-go/internal/domain"
	"github.com/RonaldoSetzer/todo-api-go/internal/errors"
	"github.com/google/uuid"
)

type LocalRepository struct {
	todos  map[string]domain.TodoDTO
}

func NewLocalRepository() *LocalRepository {
	return &LocalRepository{
		todos:  make(map[string]domain.TodoDTO),
	}
}

func (r *LocalRepository) GetTodos() ([]domain.TodoDTO, error) {
	todos := make([]domain.TodoDTO, 0, len(r.todos))
	for _, todo := range r.todos {
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *LocalRepository) GetTodoById(id uuid.UUID) (domain.TodoDTO, error) {
	todo, ok := r.todos[id.String()]
	if !ok {
		return domain.TodoDTO{}, errors.TodoNotFoundError{ID: id}
	}
	return todo, nil
}

func (r *LocalRepository) AddTodo(newTodo domain.Todo) (domain.TodoDTO, error) {
  dto := domain.NewTodoFactory().MapTodoToDto(newTodo)
  r.todos[dto.ID] = dto
  return dto, nil
}

func (r *LocalRepository) UpdateTodo(updatedTodo domain.Todo) (domain.TodoDTO, error) {
  dto := domain.NewTodoFactory().MapTodoToDto(updatedTodo)
	_, ok := r.todos[dto.ID]
	if !ok {
		return domain.TodoDTO{}, errors.TodoNotFoundError{ID: updatedTodo.ID}
	}
	r.todos[dto.ID] = dto
	return dto, nil
}

func (r *LocalRepository) DeleteTodo(id uuid.UUID) (domain.TodoDTO, error) {
	todo, ok := r.todos[id.String()]
	if !ok {
		return domain.TodoDTO{}, errors.TodoNotFoundError{ID: id}
	}
  delete(r.todos, id.String())
	return todo, nil
}
