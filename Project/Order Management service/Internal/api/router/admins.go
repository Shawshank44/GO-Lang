package router

import (
	"fmt"
	"net/http"
)

func AdminRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /getadmins", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to admins page")
	})

	return mux
}
