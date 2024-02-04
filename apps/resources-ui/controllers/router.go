package controllers

import (
	"net/http"
	"strings"
)

// Generic Resources API Router
// Doing this because I want to avoid any Golang webframework
// Can simply manage the API endpoints this way for now
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
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
