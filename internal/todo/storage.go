package todo

import "go.mongodb.org/mongo-driver/mongo"

type TodoStorage struct {
	db *mongo.Database
}

func NewTodoStorage(db *mongo.Database) *TodoStorage {
	return &TodoStorage{
		db: db,
	}
}
