package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	Id          int    `json:"id"`
	IsCompleted bool   `json:"isCompleted"`
	Body        string `json:"body"`
}

func main() {
	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	todos := []Todo{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	//Create a todo
	app.Post("/api/add-todos", func(c *fiber.Ctx) error {
		var todo Todo

		if err := c.BodyParser(&todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{
				"message": "Todo cannot be empty",
			})
		}

		todo.Id = len(todos) + 1

		todos = append(todos, todo)

		return c.Status(201).JSON(todos)
	})

	//Get all todos
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	//Get a todo by id
	app.Get("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		val, err := strconv.Atoi(id) // Convert the id to an int
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid ID",
			})
		}

		for _, t := range todos {
			if t.Id == val {
				return c.Status(200).JSON(t)
			}
		}

		return c.Status(404).JSON(fiber.Map{
			"message": "Todo not found",
		})
	})

	//Update a todo
	app.Put("/api/update-todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		var todo Todo

		if err := c.BodyParser(&todo); err != nil {
			return err
		}

		val, err := strconv.Atoi(id) // Convert the id to an int
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid ID",
			})
		}

		for i, t := range todos {
			if t.Id == val {
				todos[i].Body = todo.Body
				return c.Status(200).JSON(todos)
			}
		}

		return c.Status(404).JSON(fiber.Map{
			"message": "Todo not found",
		})
	})

	//Update a todo status
	app.Put("/api/todos/:id/complete", func(c *fiber.Ctx) error {
		id := c.Params("id")

		val, err := strconv.Atoi(id) // Convert the id to an int
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid ID",
			})
		}

		for i, t := range todos {
			if t.Id == val {
				todos[i].IsCompleted = true
				return c.Status(200).JSON(todos)
			}
		}

		return c.Status(404).JSON(fiber.Map{
			"message": "Todo not found",
		})
	})

	//Delete a todo
	app.Delete("/api/delete-todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		val, err := strconv.Atoi(id) // Convert the id to an int
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid ID",
			})
		}

		for i, t := range todos {
			if t.Id == val {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(todos)
			}
		}

		return c.Status(404).JSON(fiber.Map{
			"message": "Todo not found",
		})
	})

	log.Fatal(
		app.Listen(":" + PORT))
}
