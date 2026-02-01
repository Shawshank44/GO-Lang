package router

import (
	users "blog_rest_api/internal/handlers/Users"
	"fmt"
	"net/http"
)

func UsersRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Read
	mux.HandleFunc("GET /getusers", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Users page"))
	})
	mux.HandleFunc("GET /getusers/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to users by id page.", r.PathValue("id"))
	})
	// Create
	mux.HandleFunc("POST /users/register", users.RegisterUser)
	mux.HandleFunc("POST /users/login", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to login page"))
	})
	mux.HandleFunc("POST /users/logout", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to logout page"))
	})

	//Update
	mux.HandleFunc("PATCH /updateusers/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to update users by id", r.PathValue("id"))
	})

	// Delete
	mux.HandleFunc("DELETE /users/deactivate", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to delete page"))
	})

	// MISCs:
	mux.HandleFunc("POST /users/forgotpassword", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to forgot password page"))
	})

	mux.HandleFunc("POST /users/resetpassword", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to resetpassword page"))
	})

	return mux
}
