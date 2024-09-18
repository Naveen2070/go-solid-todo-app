package main

import (
	"log"
	"os"

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

	// Get the PORT from the environment
	PORT := os.Getenv("PORT")

	// Define routes
	app.Get("/", todocontrollers.HelloWorld)
	app.Post("/api/add-todos", todocontrollers.AddTodoHandler)
	app.Get("/api/todos", todocontrollers.GetTodosHandler)
	app.Get("/api/todos/:id", todocontrollers.GetTodoByIdHandler)
	app.Put("/api/update-todos/:id", todocontrollers.UpdateTodoHandler)
	app.Put("/api/todos/:id/complete", todocontrollers.MarkTodoCompleteHandler)
	app.Delete("/api/delete-todos/:id", todocontrollers.DeleteTodoHandler)

	log.Fatal(app.Listen(":" + PORT))
}
