package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/person", debugJSON).Methods("GET")

	log.Println("Server listening on :8090")
	log.Fatal(http.ListenAndServe(":8090", r))
}
type Person struct {
	Name string
	Age  int
  }
func debugJSON(w http.ResponseWriter, r *http.Request) {
	person := Person{"John Doe", 30}
	json, err := json.Marshal(person)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
  }