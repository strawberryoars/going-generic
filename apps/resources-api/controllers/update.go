package controllers

import (
	"fmt"
	"log"
	"net/http"
)

func UpdateHandler(w http.ResponseWriter, r *http.Request, resourceName string) {
	logMessage := fmt.Sprintf("TODOOO logic for handling PUT request to /resources/%s", resourceName)
	log.Println(logMessage)
}
