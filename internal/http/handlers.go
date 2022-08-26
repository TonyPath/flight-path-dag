package http

import (
	"net/http"

	"github.com/TonyPath/flight-path-dag/internal/models"
)

type FlightPathService interface {
	FindStartingAndEnding(flights []models.Flight) (models.TotalFlightPath, error)
}

type FlightHandlers struct {
	FlightPathService FlightPathService
}

func (h FlightHandlers) FindStartingAndEnding(w http.ResponseWriter, r *http.Request) {
	var in []models.Flight

	err := readJSON(r, &in)
	if err != nil {
		respondError(w, err, http.StatusUnprocessableEntity)
		return
	}

	res, err := h.FlightPathService.FindStartingAndEnding(in)
	if err != nil {
		respondError(w, err, http.StatusInternalServerError)
		return
	}

	out := struct {
		Starting string `json:"starting"`
		Ending   string `json:"ending"`
	}{
		Starting: res.Starting,
		Ending:   res.Ending,
	}

	respond(w, out, http.StatusOK)
}
