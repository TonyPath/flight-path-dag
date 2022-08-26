package service

import "github.com/TonyPath/flight-path-dag/internal/models"

type FlightPathService struct{}

func (svc *FlightPathService) FindStartingAndEnding(flights []models.Flight) (models.TotalFlightPath, error) {
	// will be used to build the graph
	mRoutes := make(map[string]string)

	// will be used to find the starting point.
	//
	destinations := make(map[string]struct{}, 0)

	for _, route := range flights {
		mRoutes[route[0]] = route[1]
		destinations[route[1]] = struct{}{}
	}

	startingAirport := findStartingAirport(mRoutes, destinations)
	endingAirport := findEndingAirport(mRoutes, startingAirport)

	return models.TotalFlightPath{
		Starting: startingAirport,
		Ending:   endingAirport,
	}, nil

}

func findStartingAirport(mRoutes map[string]string, destinations map[string]struct{}) string {
	var start string
	// traverse the graph to find out the airport not listed in destinations.
	for from := range mRoutes {
		if _, found := destinations[from]; !found {
			start = from
			break
		}
	}

	return start
}

func findEndingAirport(mRoutes map[string]string, src string) string {
	// traverse the graph (DFS) to find out the last point
	for {
		dest, ok := mRoutes[src]
		if !ok {
			return src
		}
		src = dest
	}
}
