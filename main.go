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
	mux.HandleFunc("/caa", controller.CreateRequest)

	// e := echo.New()
	config.InitPort()
	// Routes(e)

	// if err := e.Start(config.PORT); err != nil {
	// 	e.Logger.Fatal(err)
	// }

	err := http.ListenAndServe(config.PORT, mux)
	log.Fatal(err)
}

// func Routes(e *echo.Echo) {
// 	e.POST("/caa", controller.InitiatedChat)
// }
