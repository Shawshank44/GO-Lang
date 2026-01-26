package router

import "net/http"

func MainRouter() *http.ServeMux {
	Urouter := UsersRouter()
	Prouter := PostsRouter()
	Arouter := AdminRouter()

	Urouter.Handle("/", Prouter)
	Prouter.Handle("/", Arouter)

	return Urouter
}
