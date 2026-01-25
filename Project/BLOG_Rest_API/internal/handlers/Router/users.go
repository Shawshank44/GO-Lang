package router

import "net/http"

func UsersRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Users page"))
	})

	return mux
}
