package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ListHandler(w http.ResponseWriter, r *http.Request, resourceName string) {
	resp, err := http.Get("http://localhost:8080/resources/test?page=1&pageSize=2")
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

	// Process the resources and generate the HTML
	html := ""
	results := resources["results"].([]interface{})
	for _, resource := range results {
		resourceMap := resource.(map[string]interface{})
		typeField := resourceMap["type"].(string)
		html += "<tr><td>" + typeField + "</td></tr>"
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}
