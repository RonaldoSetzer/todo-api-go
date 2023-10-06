package application

import "github.com/RonaldoSetzer/todo-api-go/internal/domain"

type GetTodoUseCase struct {
	repository domain.Repository
}

func NewGetTodoUseCase(repository domain.Repository) *GetTodoUseCase {
	return &GetTodoUseCase{repository: repository}
}

func (g *GetTodoUseCase) Execute(id int) (domain.Todo, error) {
	return g.repository.GetTodoById(id)
}
