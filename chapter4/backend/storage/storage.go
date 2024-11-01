package storage

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type Storage struct {
	client *mongo.Client
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Connect() error {
	user := os.Getenv("MONGO_USER")
	pass := os.Getenv("MONGO_PASS")

	uri := fmt.Sprintf("mongodb://%s:%s@mongo:27017/", user, pass)
	clientOptions := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	s.client = client

	fmt.Println("MongoDB connected!")
	return nil
}

func (s *Storage) Disconnect() {
	_ = s.client.Disconnect(context.Background())
}

func (s *Storage) GetAllTodos() ([]Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if s.client == nil {
		return nil, errors.New("storage is not initialized")
	}

	collection := s.client.Database("chapter4").Collection("todos")
	if collection == nil {
		return nil, errors.New("collection todos not found")
	}

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalln(err)
	}
	defer cursor.Close(ctx)

	todos := []Todo{}
	for cursor.Next(ctx) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			// TODO fix
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (s *Storage) InsertTodo(todo *Todo) (*mongo.InsertOneResult, error) {
	if s.client == nil {
		return nil, errors.New("storage is not initialized")
	}

	collection := s.client.Database("chapter4").Collection("todos")
	if collection == nil {
		return nil, errors.New("collection todos not found")
	}

	return collection.InsertOne(context.Background(), todo)
}

func (s *Storage) DeleteById(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	if s.client == nil {
		return nil, errors.New("storage is not initialized")
	}

	collection := s.client.Database("chapter4").Collection("todos")
	if collection == nil {
		return nil, errors.New("collection todos not found")
	}

	return collection.DeleteOne(context.Background(), bson.M{"_id": id})
}
