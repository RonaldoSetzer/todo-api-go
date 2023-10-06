package application

import "github.com/RonaldoSetzer/todo-api-go/internal/infrastructure"

type ApplicationBuilder struct{}

func NewApplicationBuilder() *ApplicationBuilder {
	return &ApplicationBuilder{}
}

func (builder *ApplicationBuilder) Build() *Application {
	localRepository := infrastructure.NewLocalRepository()

	addTodoUseCase := NewAddTodoUseCase(localRepository)
	getTodoUseCase := NewGetTodoUseCase(localRepository)
	getTodosUseCase := NewGetTodosUseCase(localRepository)
	updateTodoUseCase := NewUpdateTodoUseCase(localRepository)
	deleteTodoUseCase := NewDeleteTodoUseCase(localRepository)

	todoHandler := NewTodoHandler(addTodoUseCase, getTodosUseCase, getTodoUseCase, updateTodoUseCase, deleteTodoUseCase)

	return &Application{
		TodoHandler: todoHandler,
	}
}
