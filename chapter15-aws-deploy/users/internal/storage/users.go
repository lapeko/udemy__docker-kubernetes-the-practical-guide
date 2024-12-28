package storage

import "go.mongodb.org/mongo-driver/mongo"

type Users struct {
	client *mongo.Client
}

var users *Users

func NewUsers(mc *mongo.Client) *Users {
	if users == nil {
		users = &Users{
			client: mc,
		}
	}
	return users
}

func (u *Users) GetByEmail(email string) error {
	return nil
}

func (u *Users) Create(email string, hashedPassword string) error {
	return nil
}
