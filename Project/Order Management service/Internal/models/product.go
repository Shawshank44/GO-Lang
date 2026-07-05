package models

type Product struct {
	ID          int64  `json:"id,omitempty" db:"id,omitempty"`
	SKU         string `json:"sku,omitempty" db:"sku,omitempty"`
	Name        string `json:"name,omitempty" db:"name,omitempty"`
	Description string `json:"description,omitempty" db:"description,omitempty"`

	Category     string `json:"category,omitempty" db:"category,omitempty"`
	Brand        string `json:"brand,omitempty" db:"brand,omitempty"`
	Manufacturer string `json:"manufacturer,omitempty" db:"manufacturer,omitempty"`

	Price    float64 `json:"price,omitempty" db:"price,omitempty"`
	Currency string  `json:"currency,omitempty" db:"currency,omitempty"`

	Stock int    `json:"stock,omitempty" db:"stock,omitempty"`
	Unit  string `json:"unit,omitempty" db:"unit,omitempty"`

	Status string `json:"status,omitempty" db:"status,omitempty"`

	CreatedAt *string `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty" db:"updated_at,omitempty"`
}

type Inventory struct {
	ProductID int64   `json:"product_id,omitempty" db:"product_id,omitempty"`
	Price     float64 `json:"price,omitempty" db:"price,omitempty"`
	Currency  string  `json:"currency,omitempty" db:"currency,omitempty"`
	Stock     int     `json:"stock,omitempty" db:"stock,omitempty"`
	Unit      string  `json:"unit,omitempty" db:"unit,omitempty"`
	Status    string  `json:"status,omitempty" db:"status,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty" db:"updated_at,omitempty"`
}
