package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	dbConnect "github.com/Naveen2070/go-rest-api/db"
	todocontrollers "github.com/Naveen2070/go-rest-api/todo/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
)

var sqliteDB *sql.DB

func main() {
	app := fiber.New()

	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the DB_TYPE from environment variables
	dbType := os.Getenv("DB_TYPE")

	var todoCollection interface{}

	if dbType == "SQLITE" {
		// Initialize SQLite
		sqliteDB, err = sql.Open("sqlite", "todos.db")
		if err != nil {
			log.Fatal("Failed to connect to SQLite:", err)
		}

		// Auto-migrate your models (Todo) - You need to manually create the table
		_, err = sqliteDB.Exec(`CREATE TABLE IF NOT EXISTS todos (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			body TEXT NOT NULL,
			is_completed BOOLEAN NOT NULL DEFAULT FALSE
		)`)
		if err != nil {
			log.Fatal("Failed to create tables:", err)
		}

		todoCollection = sqliteDB
		fmt.Println("Using SQLite")
	} else if dbType == "MONGO" {
		// Initialize MongoDB connection
		client := dbConnect.ConnectDB()
		defer dbConnect.DisconnectDB(client)
		todoCollection = client.Database("todoApp").Collection("todos")
		fmt.Println("Using MongoDB")
	} else {
		log.Fatal("Unsupported DB_TYPE. Must be either 'SQLITE' or 'MONGO'.")
	}

	// Get the PORT from the environment
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}

	// Define routes
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

	// Start the server
	log.Fatal(app.Listen(fmt.Sprintf(":%s", PORT)))
}
