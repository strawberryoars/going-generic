package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/strawberryoars/going-generic/apps/resources-api/clients"
	"go.mongodb.org/mongo-driver/bson"
)

// Generic List Resources endpoint - /resources/:resourceName?filter={}
// URI components should be encoded like so:
//
//	let jsonString = '{"type":"article"}';
//	let encodedJson = encodeURIComponent(jsonString);
//	let url = `http://localhost:8080/resources/blog?filter=${encodedJson}`;
//
// For Example:
//
//	http://localhost:8080/resources/blogs?filter=%7B%22type%22%3A%22article%22%7D
func ListHandler(w http.ResponseWriter, r *http.Request, resourceName string) {
	logMessage := fmt.Sprintf("ListHandler - GET request /resources/%s", resourceName)
	log.Println(logMessage)

	collection := clients.Client.Database("resources").Collection(resourceName)

	var filter bson.D
	filterParam := r.URL.Query().Get("filter")

	if filterParam != "" {
		decodedFilterParam, err := url.QueryUnescape(filterParam)
		if err != nil {
			log.Println("Failed to decode filter query parameter: ", err)
			http.Error(w, "Invalid filter query parameter", http.StatusBadRequest)
			return
		}

		err = bson.UnmarshalExtJSON([]byte(decodedFilterParam), true, &filter)
		if err != nil {
			log.Println("Failed to parse filter query parameter:", err)
			http.Error(w, "Invalid filter query parameter", http.StatusBadRequest)
			return
		}
	} else {
		filter = bson.D{}
	}

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
