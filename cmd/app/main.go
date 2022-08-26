package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	httpAPI "github.com/TonyPath/flight-path-dag/internal/http"
	"github.com/TonyPath/flight-path-dag/internal/service"
)

func main() {
	_ = run()
}

func run() error {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	flightPathService := service.FlightPathService{}

	apiConfig := httpAPI.APIConfig{
		FlightPathService: &flightPathService,
	}
	routes := httpAPI.Routes(&apiConfig)
	api := http.Server{
		Handler: routes,
		Addr:    ":50000",
	}

	serverError := make(chan error, 1)
	go func() {
		serverError <- api.ListenAndServe()
	}()

	//---------------------------
	//
	select {
	case err := <-serverError:
		return fmt.Errorf("api error: %w", err)
	case sig := <-shutdown:
		log.Println("shutdown", sig)
		return nil
	}
}
