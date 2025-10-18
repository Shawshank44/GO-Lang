package router

import (
	"net/http"
	"restapi/internal/api/handlers"
)

func ExecsRouter() *http.ServeMux {
	mux := http.NewServeMux()

	//Execs by ALL
	mux.HandleFunc("GET /execs", handlers.ExecutiveHandler)
	mux.HandleFunc("POST /execs", handlers.ExecutiveHandler)
	mux.HandleFunc("PATCH /execs", handlers.ExecutiveHandler)

	//Execs by ID
	mux.HandleFunc("GET /execs/{id}", handlers.ExecutiveHandler)
	mux.HandleFunc("PATCH /execs/{id}", handlers.ExecutiveHandler)
	mux.HandleFunc("DELETE /execs/{id}", handlers.ExecutiveHandler)
	mux.HandleFunc("POST /execs/{id}/updatepassword", handlers.ExecutiveHandler)
	//Execs by Auth
	mux.HandleFunc("POST /execs/login", handlers.ExecutiveHandler)
	mux.HandleFunc("POST /execs/logout", handlers.ExecutiveHandler)
	mux.HandleFunc("POST /execs/forgotpassword", handlers.ExecutiveHandler)
	mux.HandleFunc("POST /execs/resetpassword/reset/{resetcode}", handlers.ExecutiveHandler)
	return mux
}
