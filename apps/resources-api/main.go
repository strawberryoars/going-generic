package main

import (
	"context"
	"net/http"

	"github.com/strawberryoars/going-generic/apps/resources-api/clients"
	"github.com/strawberryoars/going-generic/apps/resources-api/controllers"
)

func init() {
	clients.InitMongoConnection()
	clients.InitSchemas()
}

func main() {

	defer func() {
		if err := clients.Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	http.HandleFunc("/resources/", controllers.ResourcesRouter)
	http.ListenAndServe(":8080", nil)
}
