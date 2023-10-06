package application

import "github.com/RonaldoSetzer/todo-api-go/internal/domain"

type DeleteTodoUseCase struct {
	repository domain.Repository
}

func NewDeleteTodoUseCase(repository domain.Repository) *DeleteTodoUseCase {
	return &DeleteTodoUseCase{repository: repository}
}

func (u *DeleteTodoUseCase) Execute(todoID int) (domain.Todo, error) {
	return u.repository.DeleteTodo(todoID)
}
