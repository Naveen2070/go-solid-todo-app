package todocontrollers

import (
	"strconv"

	"github.com/Naveen2070/go-rest-api/todo/models"
	todoservices "github.com/Naveen2070/go-rest-api/todo/services"
	"github.com/gofiber/fiber/v2"
)

// HelloWorld controller
func HelloWorld(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "Hello, World!",
	})
}

// AddTodoHandler handles adding a new todo to the collection (supports MongoDB and GORM)
func AddTodoHandler(c *fiber.Ctx, collection interface{}) error {
	var todo models.Todo

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	err := todoservices.AddTodo(collection, todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to add todo",
			"error":   err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Todo added successfully",
	})
}

// GetTodosHandler handles getting all todos from the collection (supports MongoDB and GORM)
func GetTodosHandler(c *fiber.Ctx, collection interface{}) error {
	todos, err := todoservices.GetTodos(collection)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to get todos",
		})
	}
	if len(todos) == 0 {
		return c.Status(200).JSON([]models.Todo{})
	}

	return c.Status(200).JSON(todos)
}

// GetTodoByIdHandler handles getting a todo by ID (supports MongoDB and GORM)
func GetTodoByIdHandler(c *fiber.Ctx, collection interface{}) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	todo, err := todoservices.GetTodoById(collection, id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Todo not found",
		})
	}

	return c.Status(200).JSON(todo)
}

// UpdateTodoHandler handles updating a todo's body (supports MongoDB and GORM)
func UpdateTodoHandler(c *fiber.Ctx, collection interface{}) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	updatedTodo, err := todoservices.UpdateTodoBody(collection, id, todo.Body)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Todo not found",
		})
	}

	return c.Status(200).JSON(updatedTodo)
}

// MarkTodoCompleteHandler handles marking a todo as complete (supports MongoDB and GORM)
func MarkTodoCompleteHandler(c *fiber.Ctx, collection interface{}) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	updatedTodo, err := todoservices.MarkTodoComplete(collection, id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Todo not found",
		})
	}

	return c.Status(200).JSON(updatedTodo)
}

// DeleteTodoHandler handles deleting a todo by ID (supports MongoDB and GORM)
func DeleteTodoHandler(c *fiber.Ctx, collection interface{}) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	deletedTodo, err := todoservices.DeleteTodoById(collection, id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Todo not found",
		})
	}

	return c.Status(200).JSON(deletedTodo)
}
