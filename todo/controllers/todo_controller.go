package todocontrollers

import (
	"strconv"

	"github.com/Naveen2070/go-rest-api/todo/models"
	todoservices "github.com/Naveen2070/go-rest-api/todo/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func HelloWorld(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "Hello, World!",
	})
}

func AddTodoHandler(c *fiber.Ctx, collection *mongo.Collection) error {
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
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Todo added successfully",
	})
}

func GetTodosHandler(c *fiber.Ctx, collection *mongo.Collection) error {
	todos, err := todoservices.GetTodos(collection)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to get todos",
		})
	}
	if len(todos) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "No todos found",
		})
	}

	return c.Status(200).JSON(todos)
}

func GetTodoByIdHandler(c *fiber.Ctx, collection *mongo.Collection) error {
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

func UpdateTodoHandler(c *fiber.Ctx, collection *mongo.Collection) error {
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

	todos, err := todoservices.UpdateTodoBody(collection, id, todo.Body)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Todo not found",
		})
	}

	return c.Status(200).JSON(todos)
}

func MarkTodoCompleteHandler(c *fiber.Ctx, collection *mongo.Collection) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	todos, err := todoservices.MarkTodoComplete(collection, id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Todo not found",
		})
	}

	return c.Status(200).JSON(todos)
}

func DeleteTodoHandler(c *fiber.Ctx, collection *mongo.Collection) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	todos, err := todoservices.DeleteTodoById(collection, id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Todo not found",
		})
	}

	return c.Status(200).JSON(todos)
}
