package models
import (
	"encoding/json"
	"net/http"
)

type Person struct {
	Name string
	Age  int
  }

func testing_api(w http.ResponseWriter, r *http.Request) {
	person := Person{"Heng Synoeun", 2024}
	json, err := json.Marshal(person)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
  }