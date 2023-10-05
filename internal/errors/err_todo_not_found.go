package errors

import (
  "fmt"
)

type ErrTodoNotFound struct {
  ID int
}

func (e ErrTodoNotFound) Error() string {
  return fmt.Sprintf("Todo with ID %d not found", e.ID)
}
