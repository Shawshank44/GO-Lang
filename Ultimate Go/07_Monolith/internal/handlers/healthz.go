package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := struct {
		Success bool
		Message string
	}{
		Success: true,
		Message: "Okay, API is Up and Running",
	}

	err := json.NewEncoder(w).Encode(&res)
	if err != nil {
		log.Fatal("invalid payload")
	}
}
