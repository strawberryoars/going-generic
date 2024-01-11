package controllers

import (
	"fmt"
	"log"
	"net/http"
)

func CreateHandler(w http.ResponseWriter, r *http.Request, resourceName string) {
	logMessage := fmt.Sprintf("TODOOO logic for handling POST request to /resources/%s", resourceName)
	log.Println(logMessage)
}
