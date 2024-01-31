package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/santhosh-tekuri/jsonschema/v5"
	"github.com/strawberryoars/going-generic/apps/resources-api/clients"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Generic Update Resource Endpoint - /resources/:resourceName/:resourceId
//
// TODO:
// validation on patch updates
//
// Examples:
// curl -X PUT -H "Content-Type: application/json" -d '{"type": "dingo", "hello": "DAWG"}' http://localhost:8080/resources/test/65a5da5a7633ad2102896d2f
// curl -X PUT -H "Content-Type: application/json" -d '{"name":"gauge","description":"my gauge","unit":"Celsius","attributes":{},"value":72,"time_unix_nano":170520618653603}'http://localhost:8080/resources/metric/65a5d4493ae6a896bdeda729
func UpdateHandler(w http.ResponseWriter, r *http.Request, resourceName string, resourceId string) {
	logMessage := fmt.Sprintf("UpdateHandler - PUT request /resources/%s/%s", resourceName, resourceId)
	log.Println(logMessage)

	collection := clients.Client.Database("resources").Collection(resourceName)

	var updatedResource map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&updatedResource)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id, err := primitive.ObjectIDFromHex(resourceId)
	if err != nil {
		log.Println("Failed to parse resourceId:", err)
		http.Error(w, "Invalid resourceId", http.StatusBadRequest)
		return
	}

	sch, err := jsonschema.CompileString(resourceName, clients.Schemas[resourceName])
	if err != nil {
		http.Error(w, fmt.Sprintf("Resource schema error: %v", err), http.StatusBadRequest)
		return
	}

	if err = sch.Validate(updatedResource); err != nil {
		http.Error(w, fmt.Sprintf("Invalid resource: %v", err), http.StatusBadRequest)
		return
	}

	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: updatedResource}}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println("Failed to update resource:", err)
		http.Error(w, "Failed to update resource", http.StatusInternalServerError)
		return
	}

	if updateResult.MatchedCount == 0 {
		http.Error(w, "No resources found with the given resourceId", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"matchedCount": %d}`, updateResult.MatchedCount)))
}
