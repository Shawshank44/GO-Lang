package router

import (
	"net/http"
	"restapi/internal/api/handlers"
)

func ExecsRouter() *http.ServeMux {
	mux := http.NewServeMux()

	//Execs by ALL
	mux.HandleFunc("GET /execs", handlers.GetExecutivesHandler)
	mux.HandleFunc("POST /execs", handlers.AddExecutivesHandler)
	mux.HandleFunc("PATCH /execs", handlers.PatchExecutivesHandler)

	//Execs by ID
	mux.HandleFunc("GET /execs/{id}", handlers.GetExecutiveHandler)
	mux.HandleFunc("PATCH /execs/{id}", handlers.PatchExecutiveHandler)
	mux.HandleFunc("DELETE /execs/{id}", handlers.DeleteExecutiveHandler)
	mux.HandleFunc("POST /execs/{id}/updatepassword", handlers.UpdatePasswordHandler)
	// //Execs by Auth
	mux.HandleFunc("POST /execs/login", handlers.LoginHandler)
	mux.HandleFunc("POST /execs/logout", handlers.LogoutHandler)
	mux.HandleFunc("POST /execs/forgotpassword", handlers.ForgotPassword)
	mux.HandleFunc("POST /execs/resetpassword/reset/{resetcode}", handlers.ResetPasswordHandler)
	return mux
}
