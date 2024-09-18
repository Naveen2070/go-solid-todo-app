package main

import (
	"fmt"
	"log"
	"os"

	dbConnect "github.com/Naveen2070/go-rest-api/db"
	todocontrollers "github.com/Naveen2070/go-rest-api/todo/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to MongoDB
	client := dbConnect.ConnectDB()

	// Use the client (for example, to interact with a "todos" collection)
	todoCollection := client.Database("todoApp").Collection("todos")

	// Get the PORT from the environment
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}

	// Define routes, passing the collection to each controller handler
	app.Get("/", todocontrollers.HelloWorld)
	app.Post("/api/add-todos", func(c *fiber.Ctx) error {
		return todocontrollers.AddTodoHandler(c, todoCollection)
	})
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return todocontrollers.GetTodosHandler(c, todoCollection)
	})
	app.Get("/api/todos/:id", func(c *fiber.Ctx) error {
		return todocontrollers.GetTodoByIdHandler(c, todoCollection)
	})
	app.Put("/api/update-todos/:id", func(c *fiber.Ctx) error {
		return todocontrollers.UpdateTodoHandler(c, todoCollection)
	})
	app.Put("/api/todos/:id/complete", func(c *fiber.Ctx) error {
		return todocontrollers.MarkTodoCompleteHandler(c, todoCollection)
	})
	app.Delete("/api/delete-todos/:id", func(c *fiber.Ctx) error {
		return todocontrollers.DeleteTodoHandler(c, todoCollection)
	})

	// Gracefully disconnect MongoDB when the app shuts down
	defer dbConnect.DisconnectDB(client)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", PORT)))
}
