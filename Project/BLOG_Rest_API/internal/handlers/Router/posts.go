package router

import (
	"fmt"
	"net/http"
)

func PostsRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /getposts", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Posts page"))
	})
	mux.HandleFunc("GET /getpost/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to get post by id", r.PathValue("id"))
	})

	mux.HandleFunc("POST /createpost", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to create post page"))
	})

	mux.HandleFunc("PATCH /updateposts", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to update posts page"))
	})
	mux.HandleFunc("PATCH /updatepost/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to Update post by ID page", r.PathValue("id"))
	})

	mux.HandleFunc("DELETE /deleteposts", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to delete post page"))
	})

	mux.HandleFunc("DELETE /deletepost/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to delete post by ID page"))
	})

	return mux
}
