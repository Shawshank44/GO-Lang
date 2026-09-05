package router

import (
	"net/http"
	"order_mgt/pkg/storage"
)

func MainRouter(MinioService *storage.MinioService) *http.ServeMux {
	Urouter := UserRouter()
	Prouter := ProductRouter(MinioService)
	Arouter := AdminRouter()

	Urouter.Handle("/", Prouter)
	Prouter.Handle("/", Arouter)

	return Urouter
}
