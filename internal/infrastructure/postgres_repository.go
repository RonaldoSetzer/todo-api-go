package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/RonaldoSetzer/todo-api-go/internal/domain"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository() *PostgresRepository {
  dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", "db", "5432", "todouser", "todopassword", "todo_db")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) GetTodos() ([]domain.TodoDTO, error) {
	todos := make([]domain.TodoDTO, 0)
	rows, err := r.db.Query("SELECT id, title, description, status FROM todos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id string
		var title string
		var description string
		var status string
		err = rows.Scan(&id, &title, &description, &status)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, domain.TodoDTO{ID: id, Title: title, Description: description, Status: status})
	}
	return todos, nil
}

func (r *PostgresRepository) GetTodoById(id uuid.UUID) (domain.TodoDTO, error) {
	row, err := r.db.Query("SELECT id, title, description, status FROM todos WHERE id = $1", id.String())
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var todo domain.TodoDTO
	for row.Next() {
		var id string
		var title string
		var description string
		var status string
		err = row.Scan(&id, &title, &description, &status)
		if err != nil {
			log.Fatal(err)
		}
		todo = domain.TodoDTO{ID: id, Title: title, Description: description, Status: status}
	}
	return todo, nil
}

func (r *PostgresRepository) AddTodo(todo domain.Todo) (domain.TodoDTO, error) {
	todoDTO := domain.NewTodoFactory().MapTodoToDto(todo)
	row, err := r.db.Query("INSERT INTO todos (id, title, description, status) VALUES ($1, $2, $3, $4)", todoDTO.ID, todoDTO.Title, todoDTO.Description, todoDTO.Status)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	return todoDTO, nil
}

func (r *PostgresRepository) UpdateTodo(todo domain.Todo) (domain.TodoDTO, error) {
	todoDTO := domain.NewTodoFactory().MapTodoToDto(todo)
	row, err := r.db.Query("UPDATE todos SET title = $1, description = $2, status = $3 WHERE id = $4", todoDTO.Title, todoDTO.Description, todoDTO.Status, todoDTO.ID)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	return todoDTO, nil
}

func (r *PostgresRepository) DeleteTodo(id uuid.UUID) (domain.TodoDTO, error) {
	todo, err := r.GetTodoById(id)
	if err != nil {
		log.Fatal(err)
	}
	row, err := r.db.Query("DELETE FROM todos WHERE id = $1", todo.ID)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	return todo, nil
}
