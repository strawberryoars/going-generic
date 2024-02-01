package main

import (
	"context"
	"net/http"

	"github.com/strawberryoars/going-generic/apps/resources-api/clients"
	"github.com/strawberryoars/going-generic/apps/resources-api/controllers"
)

// TODO: revisit
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		next.ServeHTTP(w, r)
	})
}

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

	// http.HandleFunc("/resources/", controllers.ResourcesRouter)
	http.Handle("/resources/", enableCORS(http.HandlerFunc(controllers.ResourcesRouter)))
	http.ListenAndServe(":8080", nil)
}
