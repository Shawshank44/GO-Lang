package router

import (
	"fmt"
	"net/http"
	"order_mgt/Internal/api/handlers"
	utilssql "order_mgt/pkg/utils_sql"
)

func ProductRouter(MinioService *utilssql.MinioService) *http.ServeMux {
	mux := http.NewServeMux()

	// Session Routes:
	mux.HandleFunc("POST /session/create", handlers.CreateSession)

	// Image uploader :
	mux.HandleFunc("POST /products/assets", handlers.UploadProductImage(MinioService))

	// GET :
	mux.HandleFunc("GET /getproducts", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to products page")
	})

	mux.HandleFunc("GET /getproduct/detail/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to product detail page %v", r.PathValue("id"))
	})

	mux.HandleFunc("GET /getproduct/search", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to product search page")
	})

	// POST :
	mux.HandleFunc("POST /admins/product/registery/create", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to productregistery page")
	})

	// PATCH :
	mux.HandleFunc("PATCH /admins/product/registery/update/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to product update page. %v", r.PathValue("id"))
	})

	mux.HandleFunc("PATCH /admins/product/inventory/update/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to inventory update page %v", r.PathValue("id"))
	})

	// DELETE :
	mux.HandleFunc("DELETE /admins/product/registery/delete", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to Products DELETE Page.")
	})

	return mux
}
