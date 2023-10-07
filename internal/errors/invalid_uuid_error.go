package errors

import "fmt"

type InvalidUUIDError struct{}

func NewInvalidUUIDError() *InvalidUUIDError {
	return &InvalidUUIDError{}
}

func (e *InvalidUUIDError) Error() string {
	return fmt.Sprintf("Invalid UUID")
}
