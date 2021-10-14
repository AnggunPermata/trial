package main

import (
	"log"

	"github.com/anggunpermata/custom-webhook/config"
	"github.com/anggunpermata/custom-webhook/controller"
	"github.com/labstack/echo"
)

func main() {
	log.Println("Starting the HTTP server on port 8080")

	// mux := http.NewServeMux()
	// mux.HandleFunc("/caa", controller.InitiatedChat)

	e := echo.New()
	config.InitPort()
	Routes(e)

	if err := e.Start(config.PORT); err != nil {
		e.Logger.Fatal(err)
	}

	// err := http.ListenAndServe(config.PORT, mux)
	// log.Fatal(err)
}

func Routes(e *echo.Echo) {
	e.POST("/caa", controller.AssignAgentWebhook)
}
