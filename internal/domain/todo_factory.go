package domain

import "github.com/google/uuid"

func NewTodoFactory() *TodoFactory {
	return &TodoFactory{}
}

type TodoFactory struct{}

func (f *TodoFactory) NewTodoDTO(todo Todo) TodoDTO {
	return TodoDTO{
		ID:          todo.ID.String(),
		Title:       todo.Title,
		Description: todo.Description,
		Status:      string(todo.Status),
	}
}

func (f *TodoFactory) Create(todoDTO TodoDTO) Todo {
	if todoDTO.ID != "" {
		return f.MapTodoFromDto(todoDTO)
	}

	todo := *NewTodo(todoDTO.Title, todoDTO.Description)
	if todoDTO.Status != "" {
		todo.ChangeStatus(todoDTO.Status)
	}
	return todo
}

func (f *TodoFactory) MapTodoFromDto(todoDTO TodoDTO) Todo {
	ID := uuid.MustParse(todoDTO.ID)
	Status := TodoStatus(todoDTO.Status)
	return Todo{
		ID:          ID,
		Title:       todoDTO.Title,
		Description: todoDTO.Description,
		Status:      Status,
	}
}

func (f *TodoFactory) MapTodoToDto(todo Todo) TodoDTO {
  return TodoDTO{
    ID:          todo.ID.String(),
    Title:       todo.Title,
    Description: todo.Description,
    Status:      string(todo.Status),
  }
}
