package domain

type Repository interface {
	GetTodos() ([]Todo, error)
	GetTodoById(id int) (Todo, error)
	AddTodo(todo Todo) (Todo, error)
	UpdateTodo(todo Todo) (Todo, error)
	DeleteTodo(id int) (Todo, error)
}
