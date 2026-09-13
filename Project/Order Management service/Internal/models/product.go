package models

type ProductImage struct {
	URL       string `json:"url,omitempty" db:"url,omitempty"`
	IsPrimary bool   `json:"is_primary" db:"is_primary"` // omitempty is not needed because is_primary will have no value if false
}

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

	CreatedAt          *string `json:"created_at,omitempty" db:"created_at,omitempty"`
	SpecsUpdatedAt     *string `json:"spec_updated_at,omitempty" db:"spec_updated_at,omitempty"`
	InventoryUpdatedAt *string `json:"inventory_updated_at,omitempty" db:"inventory_updated_at,omitempty"`

	UpdatedBy *string `json:"updated_by,omitempty" db:"updated_by,omitempty"`

	Images []ProductImage `json:"images,omitempty" db:"images,omitempty"`
}

type Inventory struct {
	Price              float64 `json:"price,omitempty" db:"price,omitempty"`
	Currency           string  `json:"currency,omitempty" db:"currency,omitempty"`
	Stock              int     `json:"stock,omitempty" db:"stock,omitempty"`
	Unit               string  `json:"unit,omitempty" db:"unit,omitempty"`
	UpdatedBy          *string `json:"updated_by,omitempty" db:"updated_by,omitempty"`
	InventoryUpdatedAt *string `json:"inventory_updated_at,omitempty" db:"inventory_updated_at,omitempty"`
	Status             string  `json:"status,omitempty" db:"status,omitempty"`
}
