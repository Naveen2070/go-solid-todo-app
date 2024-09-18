package todoservices

import (
	"errors"

	"github.com/Naveen2070/go-rest-api/todo/models"
)

var todos []models.Todo

// Add a new todo
func AddTodo(todo models.Todo) []models.Todo {
	todo.Id = len(todos) + 1
	todos = append(todos, todo)
	return todos
}

// Get all todos
func GetTodos() []models.Todo {
	return todos
}

// Get a todo by ID
func GetTodoById(id int) (models.Todo, error) {
	for _, t := range todos {
		if t.Id == id {
			return t, nil
		}
	}
	return models.Todo{}, errors.New("Todo not found")
}

// Update a todo body
func UpdateTodoBody(id int, body string) ([]models.Todo, error) {
	for i, t := range todos {
		if t.Id == id {
			todos[i].Body = body
			return todos, nil
		}
	}
	return nil, errors.New("Todo not found")
}

// Update a todo status to completed
func MarkTodoComplete(id int) ([]models.Todo, error) {
	for i, t := range todos {
		if t.Id == id {
			todos[i].IsCompleted = true
			return todos, nil
		}
	}
	return nil, errors.New("Todo not found")
}

// Delete a todo
func DeleteTodoById(id int) ([]models.Todo, error) {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return todos, nil
		}
	}
	return nil, errors.New("Todo not found")
}
