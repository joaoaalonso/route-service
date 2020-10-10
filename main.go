package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.elastic.co/apm/module/apmgorilla"
)

// Point interface
type Point struct {
	Latitude  float64
	Longitude float64
}

// RouteInput interface
type RouteInput struct {
	ReturnToOrigin bool
	Origin         Point
	Points         []Point
}

func routing(w http.ResponseWriter, r *http.Request) {
	var input RouteInput
	json.NewDecoder(r.Body).Decode(&input)

	result, err := router(input.ReturnToOrigin, input.Origin, input.Points)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		errorResponse(w, err, 404)
		return
	}

	json.NewEncoder(w).Encode(result)

}

func errorResponse(w http.ResponseWriter, err error, statusCode int) {
	http.Error(w, err.Error(), statusCode)
}

func loggingHandler(h http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, h)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/route", routing).Methods("POST")

	port := "8000"

	router.Use(loggingHandler)
	router.Use(handlers.CompressHandler)
	router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	router.Use(apmgorilla.Middleware())

	log.Println("listening on " + port)
	http.ListenAndServe(":"+port, router)
}
