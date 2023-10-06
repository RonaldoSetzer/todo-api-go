package application

import "github.com/RonaldoSetzer/todo-api-go/internal/domain"

type AddTodoUseCase struct {
	repository domain.Repository
}

func NewAddTodoUseCase(repository domain.Repository) *AddTodoUseCase {
	return &AddTodoUseCase{repository: repository}
}

func (u *AddTodoUseCase) Execute(todo domain.Todo) (domain.Todo, error) {
	return u.repository.AddTodo(todo)
}
