package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/strawberryoars/going-generic/apps/resources-api/clients"
)

// Generic Create Resource Endpoint - /resources/:resourceName
//
// TODO:
// validation with JSON schemas
//
// Examples:
// curl -X POST -H "Content-Type: application/json" -d '{"type": "cajun", "hello": "stirfry"}' http://localhost:8080/resources/blogs
func CreateHandler(w http.ResponseWriter, r *http.Request, resourceName string) {
	logMessage := fmt.Sprintf("CreateHandler - POST request /resources/%s", resourceName)
	log.Println(logMessage)

	collection := clients.Client.Database("resources").Collection(resourceName)

	var newResource map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&newResource)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	insertResult, err := collection.InsertOne(context.TODO(), newResource)
	if err != nil {
		log.Println("Failed to insert new resource:", err)
		http.Error(w, "Failed to create resource", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"id": "%s"}`, insertResult.InsertedID)))
}
