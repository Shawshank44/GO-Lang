package router

import (
	"net/http"
	"restapi/internal/api/handlers"
)

func TeachersRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Teacher by ALL
	mux.HandleFunc("GET /teachers", handlers.GetTeachersHandeler)
	mux.HandleFunc("POST /teachers", handlers.AddTeacherHandler)
	mux.HandleFunc("PATCH /teachers", handlers.PatchTeachersHandler)
	mux.HandleFunc("DELETE /teachers", handlers.DeleteTeachersHandler)
	// Teacher by ID
	mux.HandleFunc("GET /teachers/{id}", handlers.GetTeacherHandeler)
	mux.HandleFunc("PUT /teachers/{id}", handlers.PutTeacherHandler)
	mux.HandleFunc("PATCH /teachers/{id}", handlers.PatchTeacherHandler)
	mux.HandleFunc("DELETE /teachers/{id}", handlers.DeleteTeacherHandler)

	// Subroutes :
	mux.HandleFunc("GET /teachers/{id}/students", handlers.GetStudentsByTeacherID)
	mux.HandleFunc("GET /teachers/{id}/studentcount", handlers.GetStudentsCountByTeacherID)

	return mux
}
