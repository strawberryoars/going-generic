package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/santhosh-tekuri/jsonschema/v5"
	"github.com/strawberryoars/going-generic/apps/resources-api/clients"
)

// Generic Create Resource Endpoint - /resources/:resourceName
//
// Examples:
// curl -X POST -H "Content-Type: application/json" -d '{"type": "apple", "hello": "world"}' http://localhost:8080/resources/test
// curl -X POST -H "Content-Type: application/json" -d '{"type": "banjo", "hello": "music"}' http://localhost:8080/resources/test
// curl -X POST -H "Content-Type: application/json" -d '{"type": "cajun", "hello": "stirfry"}' http://localhost:8080/resources/test
// curl -X POST -H "Content-Type: application/json" -d '{"type": "dingo", "hello": "dawg"}' http://localhost:8080/resources/test
// curl -X POST -H "Content-Type: application/json" -d '{"name":"gauge","description":"my gauge","unit":"Celsius","attributes":{},"value":44,"time_unix_nano":170520618653603}' http://localhost:8080/resources/metric
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

	sch, err := jsonschema.CompileString(resourceName, clients.Schemas[resourceName])
	if err != nil {
		http.Error(w, fmt.Sprintf("Resource schema error: %v", err), http.StatusBadRequest)
		return
	}

	if err = sch.Validate(newResource); err != nil {
		http.Error(w, fmt.Sprintf("Invalid resource: %v", err), http.StatusBadRequest)
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
