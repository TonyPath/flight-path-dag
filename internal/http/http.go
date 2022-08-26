package http

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type APIConfig struct {
	FlightPathService FlightPathService
}

func Routes(cfg *APIConfig) *chi.Mux {
	mux := chi.NewRouter()

	flightHandlers := FlightHandlers{
		FlightPathService: cfg.FlightPathService,
	}

	mux.Post("/api/process-flight-list", flightHandlers.FindStartingAndEnding)

	return mux
}

func readJSON(r *http.Request, dst any) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	_ = r.Body.Close()

	return json.Unmarshal(body, dst)
}

func respond(w http.ResponseWriter, v any, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	b, _ := json.Marshal(v)
	_, _ = w.Write(b)
}

func respondError(w http.ResponseWriter, err error, statusCode int) {
	errorResponse := ErrorResponse{
		ErrorMessage: err.Error(),
	}

	respond(w, errorResponse, statusCode)
}

type ErrorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}
