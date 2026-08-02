package router

import (
	"fmt"
	"net/http"
)

func ProductRouter() *http.ServeMux {
	mux := http.NewServeMux()

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
