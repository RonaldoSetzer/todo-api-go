package errors

type InvalidStatusError struct{}

func NewInvalidStatusError() *InvalidStatusError {
  return &InvalidStatusError{}
}

func (e *InvalidStatusError) Error() string {
  return "Invalid status"
}
