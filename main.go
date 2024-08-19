package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/adminsynoeun/go_02/blob/main/utils/math.go"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/person", Add).Methods("GET")

	// Start the HTTP server on port 8090
	log.Println("Server listening on :8090")
	log.Fatal(http.ListenAndServe(":8090", r))
}