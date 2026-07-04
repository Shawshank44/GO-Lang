package router

import (
	"fmt"
	"net/http"
)

func ProductRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /getproducts", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to products page")
	})

	return mux
}
