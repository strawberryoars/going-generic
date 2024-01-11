package controllers

import (
	"net/http"
	"strings"
)

func ResourcesRouter(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 || parts[1] != "resources" {
		http.NotFound(w, r)
		return
	}

	resourceName := parts[2]

	switch r.Method {
	case http.MethodGet:
		ListHandler(w, r, resourceName)
	case http.MethodPost:
		CreateHandler(w, r, resourceName)
	case http.MethodPut:
		UpdateHandler(w, r, resourceName)
	case http.MethodDelete:
		DeleteHandler(w, r, resourceName)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
