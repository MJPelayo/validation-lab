package main

import (
	"encoding/json"
	"net/http"
)

func readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(data)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return err
	}

	return nil
}