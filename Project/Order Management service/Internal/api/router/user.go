package router

import (
	"fmt"
	"net/http"
)

func UserRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /getusers", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to users page")
	})

	return mux
}
