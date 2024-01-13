package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"

	"github.com/strawberryoars/going-generic/apps/resources-api/clients"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Generic List Resources endpoint - /resources/:resourceName?filter={}&sort={}
// URI components for filtering and sorting should be encoded like so:
//
//	let jsonString = '{"type":"article"}';
//	let encodedJson = encodeURIComponent(jsonString);
//	let url = `http://localhost:8080/resources/blog?filter=${encodedJson}`;
//
// Sorting:
// sort order: 1 for ascendeing and -1 for descending
// sort = { key: order }
//
// For Example:
//
//	curl -X GET 'http://localhost:8080/resources/blogs?page=1&pageSize=1
//	curl -X GET 'http://localhost:8080/resources/blogs?filter=%7B%22type%22%3A%22article%22%7D&page=1&pageSize=10'
//	curl -X GET 'http://localhost:8080/resources/blogs?sort=%7B%22type%22%3A1%7D&page=1&pageSize=10
func ListHandler(w http.ResponseWriter, r *http.Request, resourceName string) {
	logMessage := fmt.Sprintf("ListHandler - GET request /resources/%s", resourceName)
	log.Println(logMessage)

	collection := clients.Client.Database("resources").Collection(resourceName)

	// Parse Filter Query Parameter
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

	// Parse Sort Query Parameter
	var sort bson.D
	sortParam := r.URL.Query().Get("sort")

	if sortParam != "" {
		decodedSortParam, err := url.QueryUnescape(sortParam)
		if err != nil {
			log.Println("Failed to decode sort query parameter: ", err)
			http.Error(w, "Invalid sort query parameter", http.StatusBadRequest)
			return
		}

		err = bson.UnmarshalExtJSON([]byte(decodedSortParam), true, &sort)
		if err != nil {
			log.Println("Failed to parse sort query parameter:", err)
			http.Error(w, "Invalid sort query parameter", http.StatusBadRequest)
			return
		}
	} else {
		sort = bson.D{}
	}

	// Parse Pagination Query Parameter
	pageParam := r.URL.Query().Get("page")
	pageSizeParam := r.URL.Query().Get("pageSize")

	pageNumber, err := strconv.Atoi(pageParam)
	if err != nil || pageNumber < 1 {
		log.Println("Invalid or missing page number")
		http.Error(w, "Invalid or missing page number", http.StatusBadRequest)
		return
	}

	pageSize, err := strconv.Atoi(pageSizeParam)
	if err != nil || pageSize < 1 {
		log.Println("Invalid or missing page size")
		http.Error(w, "Invalid or missing page size", http.StatusBadRequest)
		return
	}

	totalCount, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	totalNumberOfPages := int(math.Ceil(float64(totalCount) / float64(pageSize)))
	skip := (pageNumber - 1) * pageSize
	opts := options.Find().SetSort(sort).SetSkip(int64(skip)).SetLimit(int64(pageSize))

	// opts := options.Find().SetSort(sort)
	var results []map[string]interface{}
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		panic(err)
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	response := map[string]interface{}{
		"results": results,
		"pagination": map[string]interface{}{
			"total": totalCount,
			"page":  pageNumber,
			"pages": totalNumberOfPages,
		},
	}

	res, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
