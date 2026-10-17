package handlers

import (
	"fmt"
	"strings"
	"time"
)

type CreateListingRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	City        string `json:"city"`
}

type CreateListingResponse struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

type ValidationError struct {
	Field string
	MSG   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s : %s", e.Field, e.MSG)
}

func (req CreateListingRequest) Validate() error {
	if strings.TrimSpace(req.Title) == "" {
		return &ValidationError{
			Field: "title",
			MSG:   "Must not be empty",
		}
	} else if strings.TrimSpace(req.Description) == "" {
		return &ValidationError{
			Field: "description",
			MSG:   "Must not be empty",
		}
	} else if req.Price <= 0 {
		return &ValidationError{
			Field: "price",
			MSG:   "Must be greater than zero",
		}
	} else if strings.TrimSpace(req.City) == "" {
		return &ValidationError{
			Field: "city",
			MSG:   "Must not be empty",
		}
	}

	return nil
}
