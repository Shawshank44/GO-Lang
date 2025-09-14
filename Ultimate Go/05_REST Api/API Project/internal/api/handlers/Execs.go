package handlers

import (
	"fmt"
	"net/http"
)

func ExecutiveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Welcome to Executive GET"))
		return
	case http.MethodPost:
		w.Write([]byte("Welcome to Executive POST"))
		return
	case http.MethodPut:
		w.Write([]byte("Welcome to Executive PUT"))
		return
	case http.MethodPatch:
		w.Write([]byte("Welcome to Executive PATCH"))
		return
	case http.MethodDelete:
		w.Write([]byte("Welcome to Executive DELETE"))
		return
	}

	w.Write([]byte("Hello Welcome Executives"))
}
