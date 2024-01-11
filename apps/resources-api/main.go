package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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

	uri := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(uri)
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	// Get the 'collection' query parameter
	collectionParam := r.URL.Query().Get("collection")
	if collectionParam == "" {
		http.Error(w, "Missing 'collection' query parameter", http.StatusBadRequest)
		return
	}

	// Get a handle for your collection
	collection := client.Database("resources").Collection(collectionParam)

	// Define the filter
	filter := bson.D{{Key: "type", Value: "article"}}

	// Execute the query
	var results []map[string]interface{}
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

func listHandler(w http.ResponseWriter, r *http.Request, resourceName string) {
	logMessage := fmt.Sprintf("TODO logic for handling GET request to /resources/%s", resourceName)
	log.Println(logMessage)
}

func createHandler(w http.ResponseWriter, r *http.Request, resourceName string) {
	logMessage := fmt.Sprintf("TODO logic for handling POST request to /resources/%s", resourceName)
	log.Println(logMessage)
}

func updateHandler(w http.ResponseWriter, r *http.Request, resourceName string) {
	logMessage := fmt.Sprintf("TODO logic for handling PUT request to /resources/%s", resourceName)
	log.Println(logMessage)
}

func deleteHandler(w http.ResponseWriter, r *http.Request, resourceName string) {
	logMessage := fmt.Sprintf("TODO logic for handling DELETE request to /resources/%s", resourceName)
	log.Println(logMessage)
}

func resourceHandler(w http.ResponseWriter, r *http.Request) {
	// Split the URL path into parts
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 || parts[1] != "resources" {
		http.NotFound(w, r)
		return
	}

	// The resource name is the third part of the URL
	resourceName := parts[2]

	// Now you can use resourceName in your handler logic
	switch r.Method {
	case http.MethodGet:
		listHandler(w, r, resourceName)
	case http.MethodPost:
		createHandler(w, r, resourceName)
	case http.MethodPut:
		updateHandler(w, r, resourceName)
	case http.MethodDelete:
		deleteHandler(w, r, resourceName)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
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

	http.HandleFunc("/resources/", resourceHandler)
	http.ListenAndServe(":8080", nil)
}
