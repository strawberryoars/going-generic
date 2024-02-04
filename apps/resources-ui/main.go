package main

import (
	"net/http"

	"github.com/strawberryoars/going-generic/apps/resources-ui/controllers"
)

func main() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)
	http.Handle("/resources/", http.HandlerFunc(controllers.ResourcesRouter))
	http.ListenAndServe(":5000", nil)
}
