package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

type student struct {
	Name  string `bson:"name"`
	Grade int    `bson:"Grade"`
}

// connect as setup connection to db
func connect() (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://root:root@localhost:27017")
	// Create new client for connect to mongo
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	// Connect as run goroutine to monitor status connection
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	// Return with set database name, and nil for errors
	return client.Database("go_mongodb"), nil
}

// insert as insert data to db
func insert() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, student{"wick", 2})
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, student{"Ethan", 2})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Insert Success!")
}

func find() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	cursor, err := db.Collection("student").Find(ctx, bson.M{"name": "Aisyah"})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer cursor.Close(ctx)

	result := make([]student, 0)
	for cursor.Next(ctx) {
		var row student
		err := cursor.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}

	if len(result) > 0 {
		fmt.Println("Name  :", result[0].Name)
		fmt.Println("Grade :", result[0].Grade)
	}

}

func update() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	var selector = bson.M{"name": "wick"}
	var changes = student{"Aisyah", 2}
	_, err = db.Collection("student").UpdateOne(ctx, selector, bson.M{"$set": changes})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Update success!")
}

func remove() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	var selector = bson.M{"name": "Aisyah"}
	_, err = db.Collection("student").DeleteOne(ctx, selector)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Remove success!")
}

func main() {
	remove()
}
