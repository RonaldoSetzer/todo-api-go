package application

import "github.com/RonaldoSetzer/todo-api-go/internal/infrastructure"

type ApplicationBuilder struct{}

func NewApplicationBuilder() *ApplicationBuilder {
	return &ApplicationBuilder{}
}

func (builder *ApplicationBuilder) Build() *Application {
	repository := infrastructure.NewPostgresRepository()

	addTodoUseCase := NewAddTodoUseCase(repository)
	getTodoUseCase := NewGetTodoUseCase(repository)
	getTodosUseCase := NewGetTodosUseCase(repository)
	updateTodoUseCase := NewUpdateTodoUseCase(repository)
	deleteTodoUseCase := NewDeleteTodoUseCase(repository)

	todoHandler := NewTodoHandler(addTodoUseCase, getTodosUseCase, getTodoUseCase, updateTodoUseCase, deleteTodoUseCase)

	return &Application{
		TodoHandler: todoHandler,
	}
}
