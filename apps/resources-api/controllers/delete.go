package controllers

import (
	"fmt"
	"log"
	"net/http"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request, resourceName string) {
	logMessage := fmt.Sprintf("TODOOO logic for handling DELETE request to /resources/%s", resourceName)
	log.Println(logMessage)
}
