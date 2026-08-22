package handlers

import (
	"encoding/json"
	"net/http"
	"order_mgt/Internal/api/middlewares"
	"order_mgt/Internal/models"
	sqlconnect "order_mgt/Internal/repository/sqlConnect"
	utilssql "order_mgt/pkg/utils_sql"
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

	username, ok := r.Context().Value(middlewares.UsernameKey).(string)
	if !ok {
		http.Error(w, "username not found in context", http.StatusUnauthorized)
		return
	}
	product.UpdatedBy = &username

	exists, err := utilssql.ValidateSessionsInDB(r.Context(), sessionID)
	if err != nil {
		http.Error(w, "Unable to find the session", http.StatusInternalServerError)
		return
	}

	if !exists {
		http.Error(w, "Invalid session", http.StatusBadRequest)
		return
	}

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
