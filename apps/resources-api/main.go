package main

import (
	"context"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/strawberryoars/going-generic/apps/resources-api/clients"
	"github.com/strawberryoars/going-generic/apps/resources-api/controllers"
)

func init() {
	clients.InitMongoConnection()
}

func main() {

	defer func() {
		if err := clients.Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	err := clients.Client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/resources/", controllers.ResourcesRouter)
	http.ListenAndServe(":8080", nil)
}
