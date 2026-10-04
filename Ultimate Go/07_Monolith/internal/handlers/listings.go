package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type listing struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       string    `json:"price"`
	City        string    `json:"city"`
	CreatedAt   time.Time `json:"created_at"`
}

func List(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`SELECT id, title, description, price, city, created_at FROM listings ORDER BY created_at DESC LIMIT 100`)
		if err != nil {
			log.Printf("db.Query : %v", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		listings := make([]listing, 0)

		for rows.Next() {
			var l listing
			err := rows.Scan(&l.ID, &l.Title, &l.Description, &l.Price, &l.City, &l.CreatedAt)
			if err != nil {
				log.Printf("rows.Scan : %v", err)
				http.Error(w, "internal error", http.StatusInternalServerError)
				return
			}
			listings = append(listings, l)
		}
		if err := rows.Err(); err != nil {
			log.Printf("rows.Err : %v", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

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
			log.Printf("json.NewEncoder : %v", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

	}
}
