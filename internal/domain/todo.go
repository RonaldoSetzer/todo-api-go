package domain

type Todo struct {
  ID          int    `json:"id"`
  Title       string `json:"title"`
  Description string `json:"description"`
  Status      string `json:"status"`
}

const (
  DO    = "do"
  DONE  = "done"
  DOING = "doing"
) 
