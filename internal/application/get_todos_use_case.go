package application

import "github.com/RonaldoSetzer/todo-api-go/internal/domain"

type GetTodosUseCase struct {
	repository domain.Repository
}

func NewGetTodosUseCase(repository domain.Repository) *GetTodosUseCase {
	return &GetTodosUseCase{repository: repository}
}

func (u *GetTodosUseCase) Execute() ([]domain.Todo, error) {
	return u.repository.GetTodos()
}
