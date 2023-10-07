package application

import "github.com/RonaldoSetzer/todo-api-go/internal/domain"

type AddTodoUseCase struct {
	repository domain.Repository
}

func NewAddTodoUseCase(repository domain.Repository) *AddTodoUseCase {
	return &AddTodoUseCase{repository: repository}
}

func (u *AddTodoUseCase) Execute(title string, description string) (domain.Todo, error) {
  todo := domain.NewTodo(title, description)
	return u.repository.AddTodo(*todo)
}
