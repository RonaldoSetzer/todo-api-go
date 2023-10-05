package domain

type Repository interface {
	GetTodos() ([]Todo, error)
	GetTodoById(id int) (Todo, error)
	AddTodo() (Todo, error)
	UpdateTodo() (Todo, error)
	DeleteTodo() (Todo, error)
}
