package router

import (
	users "blog_rest_api/internal/handlers/Users"
	"net/http"
)

func UsersRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Read
	mux.HandleFunc("GET /getusers", users.GetUsers)
	mux.HandleFunc("GET /getusers/{id}", users.GetUserByID)
	// Create
	mux.HandleFunc("POST /users/register", users.RegisterUser)
	mux.HandleFunc("POST /users/login", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to login page"))
	})
	mux.HandleFunc("POST /users/logout", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to logout page"))
	})

	//Update email :
	mux.HandleFunc("POST /users/updatedetail/{id}", users.UpdateDetail)
	mux.HandleFunc("POST /users/confirmdetail/{id}", users.Confirmdetail)

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
