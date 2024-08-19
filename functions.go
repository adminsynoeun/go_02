package main

import (
	"encoding/json"
	"net/http"
)

func testing_api(w http.ResponseWriter, r *http.Request) {
	person := Person{"John Doe-02", 46}
	json, err := json.Marshal(person)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
  }