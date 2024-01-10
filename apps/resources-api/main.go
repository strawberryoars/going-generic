package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// var err error
	uri := os.Getenv("MONGODB_URI")
	log.Println("uri: ", uri)
	clientOptions := options.Client().ApplyURI(uri)

	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

}

type BlogPost struct {
	// ID    primitive.ObjectID `bson:"_id"`
	// Title string             `bson:"title"`
	// Body  string             `bson:"body"`
	Type  string `bson:"type"`
	Hello string `bson:"hello"`
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Get a handle for your collection
	collection := client.Database("resources").Collection("blogs")

	// Define the filter
	filter := bson.D{{Key: "type", Value: "article"}}

	// Execute the query
	var results []BlogPost
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	// Convert results to JSON and write to the r esponse
	res, _ := json.Marshal(results)

	log.Println(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	err := client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/query", handleRequest)
	http.ListenAndServe(":8080", nil)
}
