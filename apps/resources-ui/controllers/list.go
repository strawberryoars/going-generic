package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

type ColumnConfig struct {
	ColumnName  string // Name of the column as displayed in the table
	JsonPointer string // JSON Pointer to the value in the JSON object
}

func ListHandler(w http.ResponseWriter, r *http.Request, resourceName string) {
	parts := strings.Split(r.URL.Path, "/")
	baseUrl := "http://localhost:8080/resources/" + parts[2]
	fullUrl := baseUrl + "?" + r.URL.RawQuery
	resp, err := http.Get(fullUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var resources map[string]interface{}
	err = json.Unmarshal(body, &resources)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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

	pagination := resources["pagination"].(map[string]interface{})
	total := pagination["total"].(float64)
	totalNumberOfDocuements := int(total)

	// hardcoded for now
	columns := []ColumnConfig{
		{
			ColumnName:  "Type",
			JsonPointer: "type",
		},
		{
			ColumnName:  "Hello",
			JsonPointer: "hello",
		},
	}

	// Process the resources and generate the HTML
	html := ""
	if resources["results"] != nil {
		results := resources["results"].([]interface{})
		for _, resource := range results {
			resourceMap := resource.(map[string]interface{})
			// typeField := resourceMap["type"].(string)
			if pageNumber*pageSize < totalNumberOfDocuements {
				html += `<tr hx-get="/resources/test?page=` + strconv.Itoa(pageNumber+1) + `&pageSize=` + strconv.Itoa(pageSize) + `" hx-trigger="revealed" hx-swap="afterend">`
			} else {
				html += `<tr>`
			}
			// html += `<td>`
			// html += typeField
			// html += `</td>`
			for _, col := range columns {
				// value := gjson.Get(resourceMap, col.JsonPointer) // Use gjson package to parse JSON pointers
				// html += fmt.Sprintf("<td>%s</td>", value.String())
				var resourceBytes []byte
				if resourceBytes, err = json.Marshal(resourceMap); err != nil {
					log.Println("Error marshaling resourceMap:", err)
					continue
				}

				// Use gjson.GetBytes to get the value based on the JSON pointer
				value := gjson.GetBytes(resourceBytes, col.JsonPointer).String()
				html += fmt.Sprintf("<td>%s</td>", value)
			}

			html += `</tr>`
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}
