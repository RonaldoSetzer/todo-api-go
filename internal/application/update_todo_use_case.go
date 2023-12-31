package application

import (
	"github.com/RonaldoSetzer/todo-api-go/internal/domain"
	"github.com/google/uuid"
)

type UpdateTodoUseCase struct {
	repository domain.Repository
}

func NewUpdateTodoUseCase(repository domain.Repository) *UpdateTodoUseCase {
	return &UpdateTodoUseCase{repository: repository}
}

func (u *UpdateTodoUseCase) Execute(id uuid.UUID, todoDTO domain.TodoDTO) (domain.TodoDTO, error) {
	dto, err := u.repository.GetTodoById(id)
  todo:= domain.NewTodoFactory().MapTodoFromDto(dto)
	if err != nil {
		return domain.TodoDTO{}, err
	}
	if todoDTO.Title != "" {
    todo.ChangeTitle(todoDTO.Title)
	}
	if todoDTO.Description != "" {
		todo.ChangeDescription(todoDTO.Description)
	}
	if todoDTO.Status != "" {
		err := todo.ChangeStatus(todoDTO.Status)
		if err != nil {
			return domain.TodoDTO{}, err
		}
	}
	return u.repository.UpdateTodo(todo)
}
