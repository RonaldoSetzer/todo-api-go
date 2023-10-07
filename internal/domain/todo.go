package domain

import (
	"github.com/RonaldoSetzer/todo-api-go/internal/errors"
	"github.com/google/uuid"
)

type Todo struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TodoStatus `json:"status"`
}

func NewTodo(title string, description string) *Todo {
	status := DO
	id := uuid.New()
	return &Todo{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      status,
	}
}

func (t *Todo) ChangeTitle(title string) {
  t.Title = title
}

func (t *Todo) ChangeDescription(description string) {
  t.Description = description
}

func (t *Todo) ChangeStatus(status string) error {
  if !isValidateStatus(status) {
    return errors.NewInvalidStatusError()
  }
  t.Status = TodoStatus(status)
  return nil
}
