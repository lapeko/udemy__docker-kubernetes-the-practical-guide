package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"time"
)

type TodoRequest struct {
	Title string `json:"title"`
}

var todoCollection *mongo.Collection

func getHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := todoCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalln(err)
	}
	defer cursor.Close(ctx)

	results := []bson.M{}
	for cursor.Next(ctx) {
		var document bson.M
		if err := cursor.Decode(&document); err != nil {
			// TODO fix
			log.Fatalln(err)
		}
		results = append(results, document)
	}

	c.JSON(http.StatusOK, gin.H{"OK": true, "payload": results})
}

func postHandler(c *gin.Context) {
	todo := &TodoRequest{}

	if err := c.BindJSON(todo); err != nil {
		// TODO fix
		log.Fatalln(err)
	}

	res, err := todoCollection.InsertOne(context.Background(), todo)
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{"OK": true, "payload": res})
}

func deleteHandler(c *gin.Context) {

}

func connectMongo() (*mongo.Collection, *mongo.Client, error) {
	user := os.Getenv("MONGO_USER")
	pass := os.Getenv("MONGO_PASS")

	uri := fmt.Sprintf("mongodb://%s:%s@chapter4-mongo:27017/", user, pass)
	clientOptions := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	fmt.Println("Connected to MongoDB!")

	db := client.Database("chapter4")
	collection := db.Collection("todos")

	return collection, client, nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln(errors.New("PORT environment variable not set"))
	}

	collection, client, err := connectMongo()
	todoCollection = collection

	if err != nil {
		log.Fatalln(err)
	}

	defer client.Disconnect(context.Background())

	r := gin.Default()

	r.GET("/", getHandler)
	r.POST("/", postHandler)
	r.DELETE("/", deleteHandler)

	log.Fatalln(r.Run(fmt.Sprintf(":%s", port)))
}
