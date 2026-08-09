package router

import (
	"net/http"
	utilssql "order_mgt/pkg/utils_sql"
)

func MainRouter(MinioService *utilssql.MinioService) *http.ServeMux {
	Urouter := UserRouter()
	Prouter := ProductRouter(MinioService)
	Arouter := AdminRouter()

	Urouter.Handle("/", Prouter)
	Prouter.Handle("/", Arouter)

	return Urouter
}
