// internal/domain/local_repository.go
package infrastructure

import (
	"github.com/RonaldoSetzer/todo-api-go/internal/domain"
	"github.com/RonaldoSetzer/todo-api-go/internal/errors"
)

type LocalRepository struct {
	todos  map[int]domain.Todo
	nextID int
}

func NewLocalRepository() *LocalRepository {
	return &LocalRepository{
		todos:  make(map[int]domain.Todo),
		nextID: 1,
	}
}

func (r *LocalRepository) GetTodos() ([]domain.Todo, error) {
	todos := make([]domain.Todo, 0, len(r.todos))
	for _, todo := range r.todos {
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *LocalRepository) GetTodoById(id int) (domain.Todo, error) {
	todo, ok := r.todos[id]
	if !ok {
		return domain.Todo{}, errors.ErrTodoNotFound{ID: id}
	}
	return todo, nil
}

func (r *LocalRepository) AddTodo(newTodo domain.Todo) (domain.Todo, error) {
	newID := r.nextID
	newTodo.ID = newID
	r.todos[newID] = newTodo
	r.nextID++
	return newTodo, nil
}

func (r *LocalRepository) UpdateTodo(updatedTodo domain.Todo) (domain.Todo, error) {
	_, ok := r.todos[updatedTodo.ID]
	if !ok {
		return domain.Todo{}, errors.ErrTodoNotFound{ID: updatedTodo.ID}
	}
	r.todos[updatedTodo.ID] = updatedTodo
	return updatedTodo, nil
}

func (r *LocalRepository) DeleteTodo(id int) (domain.Todo, error) {
	todo, ok := r.todos[id]
	if !ok {
		return domain.Todo{}, errors.ErrTodoNotFound{ID: id}
	}
	delete(r.todos, id)
	return todo, nil
}
