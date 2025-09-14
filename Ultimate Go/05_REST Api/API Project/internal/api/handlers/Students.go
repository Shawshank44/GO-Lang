package handlers

import (
	"fmt"
	"net/http"
)

func StudentHandlers(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Welcome to student GET"))
		return
	case http.MethodPost:
		w.Write([]byte("Welcome to student POST"))
		return
	case http.MethodPut:
		w.Write([]byte("Welcome to student PUT"))
		return
	case http.MethodPatch:
		w.Write([]byte("Welcome to student PATCH"))
		return
	case http.MethodDelete:
		w.Write([]byte("Welcome to student DELETE"))
		return
	}

	w.Write([]byte("Hello Welcome Student"))
}
