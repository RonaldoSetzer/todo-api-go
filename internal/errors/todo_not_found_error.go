package errors

import (
	"fmt"

	"github.com/google/uuid"
)

type TodoNotFoundError struct {
  ID uuid.UUID
}

func (e TodoNotFoundError) Error() string {
  return fmt.Sprintf("Todo with ID %d not found", e.ID)
}
