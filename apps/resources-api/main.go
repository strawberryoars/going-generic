package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

clientOptions := options.Client().ApplyURI("mongodb+srv://username:password@clusterurl/test?retryWrites=true&w=majority")
client, err := mongo.Connect(context.TODO(), clientOptions)
if err != nil {
   panic(err)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Get a handle for your collection
	collection := client.Database("test").Collection("test")

	// Define the filter
	filter := bson.D{{"type", "Oolong"}}

	// Execute the query
	var results []Tea
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	// Convert results to JSON and write to the response
	res, _ := json.Marshal(results)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {
http.HandleFunc("/query", handleRequest)
http.ListenAndServe(":8080", nil)
}
 
// func main() {
// 	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, "Hello World!")
// 	})

// 	http.ListenAndServe(":8080", nil) 
// }
