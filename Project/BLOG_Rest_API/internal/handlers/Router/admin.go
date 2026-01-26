package router

import "net/http"

func AdminRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/admin/super", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to admin page"))
	})

	return mux
}
