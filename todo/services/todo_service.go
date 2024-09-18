package todoservices

import (
	"context"
	"errors"
	"time"

	"github.com/Naveen2070/go-rest-api/todo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Add a new todo to the MongoDB collection
func AddTodo(collection *mongo.Collection, todo models.Todo) error {
	// Assign a new ID (you can use an auto-increment strategy if required)
	todo.Id = int(time.Now().Unix()) // For simplicity, use the Unix timestamp as the ID
	_, err := collection.InsertOne(context.TODO(), todo)
	return err
}

// Get all todos from the MongoDB collection
func GetTodos(collection *mongo.Collection) ([]models.Todo, error) {
	var todos []models.Todo

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	if err := cursor.All(context.TODO(), &todos); err != nil {
		return nil, err
	}

	return todos, nil
}

// Get a todo by ID from the MongoDB collection
func GetTodoById(collection *mongo.Collection, id int) (models.Todo, error) {
	var todo models.Todo

	err := collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&todo)
	if err == mongo.ErrNoDocuments {
		return models.Todo{}, errors.New("todo not found")
	}

	return todo, err
}

// Update a todo body in the MongoDB collection
func UpdateTodoBody(collection *mongo.Collection, id int, body string) (models.Todo, error) {
	filter := bson.M{"id": id}
	update := bson.M{"$set": bson.M{"body": body}}

	var todos []models.Todo

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err == mongo.ErrNoDocuments {
		return models.Todo{}, errors.New("todo not found")
	}

	err = collection.FindOne(context.TODO(), filter).Decode(&todos)

	if result.MatchedCount == 0 {
		return models.Todo{}, errors.New("todo not found")
	}

	return todos[0], err
}

// Mark a todo as complete in the MongoDB collection
func MarkTodoComplete(collection *mongo.Collection, id int) (models.Todo, error) {
	filter := bson.M{"id": id}
	update := bson.M{"$set": bson.M{"isCompleted": true}}

	var todos []models.Todo

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err == mongo.ErrNoDocuments {
		return models.Todo{}, errors.New("todo not found")
	}

	err = collection.FindOne(context.TODO(), filter).Decode(&todos)

	if result.MatchedCount == 0 {
		return models.Todo{}, errors.New("todo not found")
	}

	return todos[0], err
}

// Delete a todo by ID from the MongoDB collection
func DeleteTodoById(collection *mongo.Collection, id int) (models.Todo, error) {
	filter := bson.M{"id": id}
	var todos []models.Todo

	err := collection.FindOne(context.TODO(), filter).Decode(&todos)

	if err == mongo.ErrNoDocuments {
		return models.Todo{}, errors.New("todo not found")
	}

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err == mongo.ErrNoDocuments {
		return models.Todo{}, errors.New("todo not found")
	}

	if result.DeletedCount == 0 {
		return models.Todo{}, errors.New("todo not found")
	}

	return todos[0], err
}
