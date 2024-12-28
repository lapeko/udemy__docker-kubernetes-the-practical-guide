package storage

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type Storage struct {
	client *mongo.Client
	Users  *Users
}

var storage *Storage

func NewStorage() *Storage {
	if storage == nil {
		storage = &Storage{}
	}
	return storage
}

func (s *Storage) Connect() error {
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		return errors.New("environment variable MONGODB_URI is not set")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return err
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		return errors.New(fmt.Sprintf("MongoDB is not available: %v", err))
	}

	s.Users = NewUsers(client)

	return nil
}

func (s *Storage) Disconnect() error {
	if s.client == nil {
		return errors.New("mongo client is empty")
	}
	if err := s.client.Disconnect(context.TODO()); err != nil {
		return err
	}
	return nil
}
