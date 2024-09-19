package todoservices

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Naveen2070/go-rest-api/todo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	_ "modernc.org/sqlite"
)

// Add a new todo to the collection (either MongoDB or SQLite)
func AddTodo(collection interface{}, todo models.Todo) error {
	// Assign a new ID (use the Unix timestamp for simplicity)

	switch col := collection.(type) {
	case *mongo.Collection:
		// MongoDB logic
		todo.Id = int(time.Now().Unix())
		_, err := col.InsertOne(context.TODO(), todo)
		return err
	case *sql.DB:
		// SQLite logic
		_, err := col.Exec(`INSERT INTO todos (body, is_completed) VALUES (?, ?)`,
			todo.Body, todo.IsCompleted)
		return err
	default:
		return errors.New("unsupported collection type")
	}
}

// Get all todos from the collection (either MongoDB or SQLite)
func GetTodos(collection interface{}) ([]models.Todo, error) {
	var todos []models.Todo

	switch col := collection.(type) {
	case *mongo.Collection:
		// MongoDB logic
		cursor, err := col.Find(context.TODO(), bson.M{})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(context.TODO())

		if err := cursor.All(context.TODO(), &todos); err != nil {
			return nil, err
		}
	case *sql.DB:
		// SQLite logic
		rows, err := col.Query(`SELECT id, body, is_completed FROM todos`)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var todo models.Todo
			if err := rows.Scan(&todo.Id, &todo.Body, &todo.IsCompleted); err != nil {
				return nil, err
			}
			todos = append(todos, todo)
		}
		if err := rows.Err(); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("unsupported collection type")
	}

	return todos, nil
}

// Get a todo by ID from the collection (either MongoDB or SQLite)
func GetTodoById(collection interface{}, id int) (models.Todo, error) {
	var todo models.Todo

	switch col := collection.(type) {
	case *mongo.Collection:
		// MongoDB logic
		err := col.FindOne(context.TODO(), bson.M{"id": id}).Decode(&todo)
		if err == mongo.ErrNoDocuments {
			return models.Todo{}, errors.New("todo not found")
		}
		return todo, err
	case *sql.DB:
		// SQLite logic
		row := col.QueryRow(`SELECT id, body, is_completed FROM todos WHERE id = ?`, id)
		err := row.Scan(&todo.Id, &todo.Body, &todo.IsCompleted)
		if err == sql.ErrNoRows {
			return models.Todo{}, errors.New("todo not found")
		}
		return todo, err
	default:
		return models.Todo{}, errors.New("unsupported collection type")
	}
}

// Update a todo's body in the collection (either MongoDB or SQLite)
func UpdateTodoBody(collection interface{}, id int, body string) (models.Todo, error) {
	var todo models.Todo

	switch col := collection.(type) {
	case *mongo.Collection:
		// MongoDB logic
		filter := bson.M{"id": id}
		update := bson.M{"$set": bson.M{"body": body}}

		result, err := col.UpdateOne(context.TODO(), filter, update)
		if err != nil || result.MatchedCount == 0 {
			return models.Todo{}, errors.New("todo not found")
		}

		err = col.FindOne(context.TODO(), filter).Decode(&todo)
		return todo, err
	case *sql.DB:
		// SQLite logic
		_, err := col.Exec(`UPDATE todos SET body = ? WHERE id = ?`, body, id)
		if err != nil {
			return models.Todo{}, err
		}

		return GetTodoById(col, id)
	default:
		return models.Todo{}, errors.New("unsupported collection type")
	}
}

// Mark a todo as complete in the collection (either MongoDB or SQLite)
func MarkTodoComplete(collection interface{}, id int) (models.Todo, error) {
	var todo models.Todo

	switch col := collection.(type) {
	case *mongo.Collection:
		// MongoDB logic
		filter := bson.M{"id": id}
		update := bson.M{"$set": bson.M{"isCompleted": true}}

		result, err := col.UpdateOne(context.TODO(), filter, update)
		if err != nil || result.MatchedCount == 0 {
			return models.Todo{}, errors.New("todo not found")
		}

		err = col.FindOne(context.TODO(), filter).Decode(&todo)
		return todo, err
	case *sql.DB:
		// SQLite logic
		_, err := col.Exec(`UPDATE todos SET is_completed = 1 WHERE id = ?`, id)
		if err != nil {
			return models.Todo{}, err
		}

		return GetTodoById(col, id)
	default:
		return models.Todo{}, errors.New("unsupported collection type")
	}
}

// Delete a todo by ID from the collection (either MongoDB or SQLite)
func DeleteTodoById(collection interface{}, id int) (models.Todo, error) {
	var todo models.Todo

	switch col := collection.(type) {
	case *mongo.Collection:
		// MongoDB logic
		filter := bson.M{"id": id}
		err := col.FindOne(context.TODO(), filter).Decode(&todo)
		if err == mongo.ErrNoDocuments {
			return models.Todo{}, errors.New("todo not found")
		}

		result, err := col.DeleteOne(context.TODO(), filter)
		if err != nil || result.DeletedCount == 0 {
			return models.Todo{}, errors.New("todo not found")
		}
		return todo, nil
	case *sql.DB:
		// SQLite logic
		row := col.QueryRow(`SELECT id, body, is_completed FROM todos WHERE id = ?`, id)
		err := row.Scan(&todo.Id, &todo.Body, &todo.IsCompleted)
		if err == sql.ErrNoRows {
			return models.Todo{}, errors.New("todo not found")
		}

		_, err = col.Exec(`DELETE FROM todos WHERE id = ?`, id)
		if err != nil {
			return models.Todo{}, err
		}

		return todo, nil
	default:
		return models.Todo{}, errors.New("unsupported collection type")
	}
}
