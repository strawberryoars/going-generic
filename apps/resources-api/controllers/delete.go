package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/strawberryoars/going-generic/apps/resources-api/clients"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// func DeleteHandler(w http.ResponseWriter, r *http.Request, resourceName string) {
// 	logMessage := fmt.Sprintf("TODOOO logic for handling DELETE request to /resources/%s", resourceName)
// 	log.Println(logMessage)
// }

// Generic Delete Resource Endpoint - /resources/:resourceName/:resourceId
//
// Examples:
// curl -X DELETE http://localhost:8080/resources/blogs/65a23e48e4db4a6d59746fcc
func DeleteHandler(w http.ResponseWriter, r *http.Request, resourceName string, resourceId string) {
	logMessage := fmt.Sprintf("DeleteHandler - DELETE request /resources/%s/%s", resourceName, resourceId)
	log.Println(logMessage)

	collection := clients.Client.Database("resources").Collection(resourceName)

	id, err := primitive.ObjectIDFromHex(resourceId)
	if err != nil {
		log.Println("Failed to parse resourceId:", err)
		http.Error(w, "Invalid resourceId", http.StatusBadRequest)
		return
	}

	filter := bson.D{{Key: "_id", Value: id}}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println("Failed to delete resource:", err)
		http.Error(w, "Failed to delete resource", http.StatusInternalServerError)
		return
	}

	if deleteResult.DeletedCount == 0 {
		http.Error(w, "No resources found with the given resourceId", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"deletedCount": %d}`, deleteResult.DeletedCount)))
}
