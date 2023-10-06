package application

import "github.com/RonaldoSetzer/todo-api-go/internal/domain"

type UpdateTodoUseCase struct {
	repository domain.Repository
}

func NewUpdateTodoUseCase(repository domain.Repository) *UpdateTodoUseCase {
	return &UpdateTodoUseCase{repository: repository}
}

func (u *UpdateTodoUseCase) Execute(todo domain.Todo) (domain.Todo, error) {
	return u.repository.UpdateTodo(todo)
}
