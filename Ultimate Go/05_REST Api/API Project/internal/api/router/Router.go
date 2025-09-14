package router

import (
	"net/http"
	"restapi/internal/api/handlers"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux() //multiplexer (mux)
	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/teachers/", handlers.TeachersHandler)
	mux.HandleFunc("/students/", handlers.StudentHandlers)
	mux.HandleFunc("/execs/", handlers.ExecutiveHandler)
	return mux
}
