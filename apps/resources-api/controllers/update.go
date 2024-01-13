package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/strawberryoars/going-generic/apps/resources-api/clients"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Generic Update Resource Endpoint - /resources/:resourceName/:resourceId
//
// TODO:
// validation with JSON schemas
//
// Examples:
// curl -X PUT -H "Content-Type: application/json" -d '{"hello": "taro"}' http://localhost:8080/resources/blogs/65a095882a40f07a96eb176e
// curl -X PUT -H "Content-Type: application/json" -d '{"hello": "DAWG"}' http://localhost:8080/resources/blogs/65a0967a76f3cb4d9891cdcb
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
