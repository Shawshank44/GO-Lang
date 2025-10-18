package router

import (
	"net/http"
)

func MainRouter() *http.ServeMux {
	Trouter := TeachersRouter()
	Srouter := StudentsRouter()
	Erouter := ExecsRouter()

	Srouter.Handle("/", Erouter)
	Trouter.Handle("/", Srouter)
	return Trouter
}
