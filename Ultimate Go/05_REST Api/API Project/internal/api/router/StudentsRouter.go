package router

import (
	"net/http"
	"restapi/internal/api/handlers"
)

func StudentsRouter() *http.ServeMux {
	mux := http.NewServeMux()
	//Student by ALL
	mux.HandleFunc("GET /students", handlers.GetStudentsHandeler)
	mux.HandleFunc("POST /students", handlers.AddStudentHandler)
	mux.HandleFunc("PATCH /students", handlers.PatchStudentsHandler)
	mux.HandleFunc("DELETE /students", handlers.DeleteStudentsHandler)
	//Student by ID
	mux.HandleFunc("GET /students/{id}", handlers.GetStudentHandeler)
	mux.HandleFunc("PUT /students/{id}", handlers.PutStudentHandler)
	mux.HandleFunc("PATCH /students/{id}", handlers.PatchStudentHandler)
	mux.HandleFunc("DELETE /students/{id}", handlers.DeleteStudentHandler)

	mux.HandleFunc("GET /execs", handlers.ExecutiveHandler)

	return mux
}
