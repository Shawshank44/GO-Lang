package router

import (
	"net/http"
	"order_mgt/Internal/api/handlers"
	"order_mgt/pkg/storage"
)

func ProductRouter(MinioService *storage.MinioService) *http.ServeMux { // Minio Service Called
	mux := http.NewServeMux()

	// Session Routes:
	mux.HandleFunc("POST /session/create", handlers.CreateSession)

	// Image uploader :
	mux.HandleFunc("POST /products/assets", handlers.UploadProductImage(MinioService))

	// GET :
	mux.HandleFunc("GET /getproducts", handlers.GetProducts)

	mux.HandleFunc("GET /getproduct/detail/{id}", handlers.GetProduct)

	mux.HandleFunc("GET /getproduct/search", handlers.SearchProducts)

	// POST :
	mux.HandleFunc("POST /admins/product/registery/create", handlers.CreateProduct)

	// PATCH :
	mux.HandleFunc("PATCH /admins/product/registery/{id}/update", handlers.UpdateProduct(MinioService))

	mux.HandleFunc("PATCH /admins/product/inventory/update/{id}", handlers.UpdateInventory)

	return mux
}
