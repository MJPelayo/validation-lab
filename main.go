package main

import (
	"encoding/json"
	"net/http"
)

type Student struct {
	Name      string `json:"name"`
	Programme string `json:"programme"`
	Year      int    `json:"year"`
}

func createStudent(w http.ResponseWriter, r *http.Request) {

	var input Student

	err := readJSON(w, r, &input)
	if err != nil {
		return
	}

	v := newValidator()

	// Validation rules
	v.Check(input.Name != "", "name", "must be provided")
	v.Check(len(input.Name) <= 100, "name", "must not exceed 100 characters")

	v.Check(input.Programme != "", "programme", "must be provided")

	v.Check(input.Year >= 1 && input.Year <= 4, "year", "must be between 1 and 4")

	if !v.Valid() {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(v.Errors)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(input)
}

func main() {

	http.HandleFunc("/students", createStudent)

	http.ListenAndServe(":4000", nil)
}