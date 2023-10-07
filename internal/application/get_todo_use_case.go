package application

import (
	"github.com/RonaldoSetzer/todo-api-go/internal/domain"
	"github.com/google/uuid"
)

type GetTodoUseCase struct {
	repository domain.Repository
}

func NewGetTodoUseCase(repository domain.Repository) *GetTodoUseCase {
	return &GetTodoUseCase{repository: repository}
}

func (g *GetTodoUseCase) Execute(id uuid.UUID) (domain.Todo, error) {
	return g.repository.GetTodoById(id)
}
