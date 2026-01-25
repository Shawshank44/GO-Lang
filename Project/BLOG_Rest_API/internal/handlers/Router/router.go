package router

import "net/http"

func MainRouter() *http.ServeMux {
	Urouter := UsersRouter()
	Prouter := PostsRouter()

	Urouter.Handle("/", Prouter)

	return Urouter
}
