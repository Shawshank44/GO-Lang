package router

import "net/http"

func PostsRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /posts", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Posts page"))
	})

	return mux
}
