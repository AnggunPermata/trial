package main

import (
	"log"
	"net/http"

	"github.com/anggunpermata/custom-webhook/config"
	"github.com/anggunpermata/custom-webhook/controller"
)

func main() {
	log.Println("Starting the HTTP server on port 8090")

	mux := http.NewServeMux()
	mux.HandleFunc("/caa", controller.CreateRequest)

	config.InitPort()
	err := http.ListenAndServe(config.PORT, mux)
	log.Fatal(err)
}
