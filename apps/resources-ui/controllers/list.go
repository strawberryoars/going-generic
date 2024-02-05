package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

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

	// Process the resources and generate the HTML
	html := ""
	if resources["results"] != nil {
		results := resources["results"].([]interface{})
		for _, resource := range results {
			resourceMap := resource.(map[string]interface{})
			typeField := resourceMap["type"].(string)
			if pageNumber*pageSize < totalNumberOfDocuements {
				html += `<tr hx-get="/resources/test?page=` + strconv.Itoa(pageNumber+1) + `&pageSize=` + strconv.Itoa(pageSize) + `" hx-trigger="revealed" hx-swap="afterend">`
			} else {
				html += `<tr>`
			}
			html += `<td>`
			html += typeField
			html += `</td>`
			html += `</tr>`
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}
