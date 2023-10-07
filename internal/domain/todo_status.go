package domain

type TodoStatus string

const (
	DO    TodoStatus = "DO"
	DONE  TodoStatus = "DONE"
	DOING TodoStatus = "DOING"
)

func isValidateStatus(status string) bool {
  return status == string(DO) || status == string(DONE) || status == string(DOING)
}
  
