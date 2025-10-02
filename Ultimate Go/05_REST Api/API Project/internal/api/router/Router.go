package router

import (
	"net/http"
	"restapi/internal/api/handlers"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux() //multiplexer (mux)
	mux.HandleFunc("/", handlers.RootHandler)

	// By ALL
	mux.HandleFunc("GET /teachers/", handlers.GetTeachersHandeler)
	mux.HandleFunc("POST /teachers/", handlers.AddTeacherHandler)
	mux.HandleFunc("PATCH /teachers/", handlers.PatchTeachersHandler)
	mux.HandleFunc("DELETE /teachers/", handlers.DeleteTeachersHandler)
	// By ID
	mux.HandleFunc("GET /teachers/{id}", handlers.GetTeacherHandeler)
	mux.HandleFunc("PUT /teachers/{id}", handlers.PutTeacherHandler)
	mux.HandleFunc("PATCH /teachers/{id}", handlers.PatchTeacherHandler)
	mux.HandleFunc("DELETE /teachers/{id}", handlers.DeleteTeacherHandler)

	mux.HandleFunc("/students/", handlers.StudentHandlers)

	mux.HandleFunc("/execs/", handlers.ExecutiveHandler)
	return mux
}
