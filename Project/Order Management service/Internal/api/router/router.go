package router

import "net/http"

func MainRouter() *http.ServeMux {
	Urouter := UserRouter()
	Prouter := ProductRouter()
	Arouter := AdminRouter()

	Urouter.Handle("/", Prouter)
	Prouter.Handle("/", Arouter)

	return Urouter
}
