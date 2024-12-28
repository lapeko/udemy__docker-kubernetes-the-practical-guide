package storage

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Users struct {
	collection *mongo.Collection
}

var users *Users

func NewUsers(mc *mongo.Client) *Users {
	if users == nil {
		users = &Users{
			collection: mc.Database("ch15").Collection("users"),
		}
	}
	return users
}

func (u *Users) GetByEmail(email string) (*User, error) {
	var user User
	err := u.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (u *Users) Create(email string, hashedPassword string) (*User, error) {
	user := &User{
		Email:    email,
		Password: hashedPassword,
	}
	_, err := u.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
