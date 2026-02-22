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
	mux.HandleFunc("POST /users/login", users.Login)
	mux.HandleFunc("POST /users/logout", users.Logout)

	//Update email :
	mux.HandleFunc("POST /users/updatedetail/{id}", users.UpdateDetail)
	mux.HandleFunc("POST /users/confirmdetail/{id}", users.Confirmdetail)

	// Delete
	mux.HandleFunc("DELETE /users/deactivate/{id}", users.DeactivateUser)

	// MISCs:
	mux.HandleFunc("POST /users/forgotpassword", users.ForgotPassword)

	mux.HandleFunc("POST /users/resetpassword", users.ResetPassword)

	return mux
}
