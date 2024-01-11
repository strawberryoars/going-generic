package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/strawberryoars/going-generic/apps/resources-api/clients"
	"go.mongodb.org/mongo-driver/bson"
)

func ListHandler(w http.ResponseWriter, r *http.Request, resourceName string) {
	logMessage := fmt.Sprintf("ListHandler - GET request /resources/%s", resourceName)
	log.Println(logMessage)

	collection := clients.Client.Database("resources").Collection(resourceName)

	filter := bson.D{{Key: "type", Value: "article"}}

	var results []map[string]interface{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	res, _ := json.Marshal(results)

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
