package router

import (
	"net/http"
	"restapi/internal/api/handlers"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux() //multiplexer (mux)
	mux.HandleFunc("GET /", handlers.RootHandler)

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
