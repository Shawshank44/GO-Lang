package router

import (
	"net/http"
)

func MainRouter() *http.ServeMux {
	Trouter := TeachersRouter()
	Srouter := StudentsRouter()

	Trouter.Handle("/", Srouter)
	return Trouter
}
