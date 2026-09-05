package handlers

import (
	"encoding/json"
	"net/http"
	"order_mgt/Internal/api/middlewares"
	"order_mgt/Internal/models"
	sqlconnect "order_mgt/Internal/repository/sqlConnect"
	"order_mgt/pkg/storage"
	"order_mgt/pkg/utils"
	utilssql "order_mgt/pkg/utils_sql"
	"strconv"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	sessionID := r.URL.Query().Get("session_id")

	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid json Payload", http.StatusBadRequest)
		return
	}

	exists, err := utilssql.ValidateSessionsInDB(r.Context(), sessionID)
	if err != nil {
		http.Error(w, "Unable to find the session", http.StatusInternalServerError)
		return
	}

	if !exists {
		http.Error(w, "Invalid session", http.StatusBadRequest)
		return
	}

	username, ok := r.Context().Value(middlewares.UsernameKey).(string)
	if !ok {
		http.Error(w, "username not found in context", http.StatusUnauthorized)
		return
	}
	product.UpdatedBy = &username

	id, err := sqlconnect.CreateProductInDB(r.Context(), &product, sessionID)
	if err != nil {
		http.Error(w, "Unable to create product", http.StatusInternalServerError)
		return
	}

	res := struct {
		Success   bool
		ProductID int64
		Message   string
	}{
		Success:   true,
		ProductID: id,
		Message:   "Product has been created successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&res)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	page, limit := utils.GetPaginationParams(r)

	productlist, totalproducts, err := sqlconnect.GetProductsFromDB(r.Context(), r, limit, page)
	if err != nil {
		http.Error(w, "unable to fetch products", http.StatusBadRequest)
		return
	}

	totalPages := (totalproducts + limit - 1) / limit

	res := struct {
		Status     string
		Count      int
		TotalPages int
		PageNo     int
		PageSize   int
		Data       []models.Product
	}{
		Status:     "Success",
		Count:      totalproducts,
		TotalPages: totalPages,
		PageNo:     page,
		PageSize:   limit,
		Data:       productlist,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&res)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idstr := r.PathValue("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		http.Error(w, "Invalid ID type", http.StatusBadRequest)
		return
	}

	product, err := sqlconnect.GetProductFromDB(r.Context(), id)
	if err != nil {
		http.Error(w, "unable to fetch the product", http.StatusInternalServerError)
		return
	}

	res := struct {
		Status  string
		Product models.Product
	}{
		Status:  "Success",
		Product: *product,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func UpdateProduct(minioService *storage.MinioService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		productID := r.PathValue("id")
		sessionID := r.URL.Query().Get("session_id")

		pid, err := strconv.Atoi(productID)
		if err != nil {
			http.Error(w, "unable to convert id", http.StatusForbidden)
			return
		}

		var product models.Product
		err = json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}

		exists, err := utilssql.ValidateSessionsInDB(r.Context(), sessionID)
		if err != nil {
			http.Error(w, "Unable to find the session", http.StatusInternalServerError)
			return
		}

		if !exists {
			http.Error(w, "Invalid session", http.StatusBadRequest)
			return
		}

		username, ok := r.Context().Value(middlewares.UsernameKey).(string)
		if !ok {
			http.Error(w, "username not found in context", http.StatusUnauthorized)
			return
		}
		product.UpdatedBy = &username

		err = sqlconnect.UpdateProductInDB(r.Context(), minioService, &product, pid, sessionID)
		if err != nil {
			http.Error(w, "Invalid session", http.StatusForbidden)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		res := struct {
			Success string
			Product models.Product
		}{
			Success: "Product updated successfully",
			Product: product,
		}
		json.NewEncoder(w).Encode(&res)
	}
}
