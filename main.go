package main

import (
	"log"
	"net/http"

	"github.com/anggunpermata/custom-webhook/config"
	"github.com/anggunpermata/custom-webhook/controller"
)

func main() {
	log.Println("Starting the HTTP server on port 8080")

	mux := http.NewServeMux()
	mux.HandleFunc("/caa", controller.InitiateChat)

	config.InitPort()
	err := http.ListenAndServe(config.PORT, mux)
	log.Fatal(err)
}
