package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/Shawshank44/olx-api/internal/middlewares"
)

type listing struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       string    `json:"price"`
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
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	listings := make([]listing, 0)

	for rows.Next() {
		var l listing
		err := rows.Scan(&l.ID, &l.Title, &l.Description, &l.Price, &l.City, &l.CreatedAt)
		if err != nil {
			lh.logger.Error("Listing rows scan error", "Source", "List.rows.Scan", "err", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
		listings = append(listings, l)
	}
	if err := rows.Err(); err != nil {
		log.Printf("GET - rows.Err : %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
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

	err = json.NewEncoder(w).Encode(&res)
	if err != nil {
		log.Printf("GET - json.NewEncoder : %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

}

func (lh ListingHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	ctx := r.Context()
	requestID := middlewares.RequestIDFromContext(ctx)
	_, err := lh.db.ExecContext(ctx, `DELETE FROM listing WHERE id = $1`, id)
	if err != nil {
		// log.Printf("DELETE - db.Exec : %v", err) // recommeded to use Slog package and mention the source "db.Exec"
		lh.logger.Error("delete failed", "listing_id", id, "request_id", requestID, "err", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
