package storage

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	Id    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title string             `json:"title" bson:"title"`
}
