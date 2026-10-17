package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/Shawshank44/olx-api/internal/httpx"
	"github.com/Shawshank44/olx-api/internal/middlewares"
)

type listing struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int64     `json:"price"`
	City        string    `json:"city"`
	CreatedAt   time.Time `json:"created_at"`
}

type ListingHandler struct { // go constructor pattern
	db     *sql.DB
	logger *slog.Logger
}

func NewListingHandler(db *sql.DB, logger *slog.Logger) *ListingHandler {
	return &ListingHandler{
		db:     db,
		logger: logger,
	}
}

func (lh ListingHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	rows, err := lh.db.QueryContext(ctx, `SELECT id, title, description, price, city, created_at FROM listings ORDER BY created_at DESC LIMIT 100`)
	if err != nil {
		lh.logger.Error("Listing query error", "Source", "List.db.QueryContext", "err", err)
		log.Printf("GET - db.Query : %v", err)
		httpx.Error(w, http.StatusInternalServerError, "internal error", httpx.CodeInternalError)
		return
	}
	defer rows.Close()

	listings := make([]listing, 0)

	for rows.Next() {
		var l listing
		err := rows.Scan(&l.ID, &l.Title, &l.Description, &l.Price, &l.City, &l.CreatedAt)
		if err != nil {
			lh.logger.Error("Listing rows scan error", "Source", "List.rows.Scan", "err", err)
			httpx.Error(w, http.StatusInternalServerError, "internal error", httpx.CodeInternalError)
			return
		}
		listings = append(listings, l)
	}
	if err := rows.Err(); err != nil {
		log.Printf("GET - rows.Err : %v", err)
		httpx.Error(w, http.StatusInternalServerError, "internal error", httpx.CodeInternalError)
		return
	}
	lh.logger.Info("Listings fetched", "total", len(listings))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := struct {
		Success  bool      `json:"success"`
		Status   int       `json:"status"`
		Listings []listing `json:"listings"`
	}{
		Success:  true,
		Status:   http.StatusOK,
		Listings: listings,
	}

	_ = json.NewEncoder(w).Encode(&res)

}

func (lh ListingHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestID := middlewares.RequestIDFromContext(ctx)
	var req CreateListingRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		lh.logger.Error("failed to decode", "request_id", requestID, "err", err)
		httpx.Error(w, http.StatusBadRequest, "invalid body", httpx.CodeMalformedJSON)
		return
	}

	row := lh.db.QueryRowContext(ctx, `INSERT INTO listings (title, description, price, city) VALUES ($1, $2, $3, $4) RETURNING id, created_at`, req.Title, req.Description, req.Price, req.City)
	var out CreateListingResponse
	if err := row.Scan(&out.ID, &out.CreatedAt); err != nil {
		lh.logger.Error("failed to insert", "request_id", requestID, "err", err)
		httpx.Error(w, http.StatusInternalServerError, "invalid body", httpx.CodeInternalError)
		return
	}

	lh.logger.Info("listing created", "request_id", requestID, "listing_id", out.ID, "created_at", out.CreatedAt)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := struct {
		Success bool                  `json:"success"`
		Status  int                   `json:"status"`
		List    CreateListingResponse `json:"list"`
	}{
		Success: true,
		Status:  http.StatusOK,
		List:    out,
	}

	_ = json.NewEncoder(w).Encode(&res)
}

func (lh ListingHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	ctx := r.Context()
	requestID := middlewares.RequestIDFromContext(ctx)
	_, err := lh.db.ExecContext(ctx, `DELETE FROM listings WHERE id = $1`, id)
	if err != nil {
		// log.Printf("DELETE - db.Exec : %v", err) // recommeded to use Slog package and mention the source "db.Exec"
		lh.logger.Error("delete failed", "listing_id", id, "request_id", requestID, "err", err)
		httpx.Error(w, http.StatusInternalServerError, "something went wrong", httpx.CodeInternalError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
