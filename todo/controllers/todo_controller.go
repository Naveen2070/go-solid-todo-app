package todocontrollers

import (
	"strconv"

	"github.com/Naveen2070/go-rest-api/todo/models"
	todoservices "github.com/Naveen2070/go-rest-api/todo/services"
	"github.com/gofiber/fiber/v2"
)

func HelloWorld(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "Hello, World!",
	})
}

func AddTodoHandler(c *fiber.Ctx) error {
	var todo models.Todo

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Todo body cannot be empty",
		})
	}

	todos := todoservices.AddTodo(todo)
	return c.Status(201).JSON(todos)
}

func GetTodosHandler(c *fiber.Ctx) error {
	todos := todoservices.GetTodos()
	return c.Status(200).JSON(todos)
}

func GetTodoByIdHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	todo, err := todoservices.GetTodoById(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Todo not found",
		})
	}

	return c.Status(200).JSON(todo)
}

func UpdateTodoHandler(c *fiber.Ctx) error {
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

	todos, err := todoservices.UpdateTodoBody(id, todo.Body)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Todo not found",
		})
	}

	return c.Status(200).JSON(todos)
}

func MarkTodoCompleteHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	todos, err := todoservices.MarkTodoComplete(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Todo not found",
		})
	}

	return c.Status(200).JSON(todos)
}

func DeleteTodoHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	todos, err := todoservices.DeleteTodoById(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Todo not found",
		})
	}

	return c.Status(200).JSON(todos)
}
