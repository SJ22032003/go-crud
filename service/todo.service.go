package service

import (
	"context"

	"github.com/SJ22032003/go-crud/db"
	"github.com/SJ22032003/go-crud/model"
	"go.mongodb.org/mongo-driver/bson"
)

type TodoService struct {
	ctx context.Context
}

const (
	TODO = "todos"
)

func (t *TodoService) GetAllTodos() []model.Todo {

	var todos []model.Todo

	collection := db.GetCollection(TODO)

	cursor, err := collection.Find(t.ctx, bson.M{})
	if err != nil {
		return todos
	}

	for cursor.Next(t.ctx) {
		var todo model.Todo
		cursor.Decode(&todo)
		todos = append(todos, todo)
	}

	return todos

}