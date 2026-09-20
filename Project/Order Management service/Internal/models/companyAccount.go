package models

type CompanyAccount struct {
	ID            int64  `json:"id,omitempty" db:"id"`
	AccountNumber string `json:"account_number,omitempty" db:"account_number"`
	CompanyName   string `json:"company_name,omitempty" db:"company_name"`

	AddressLine1 string  `json:"address_line_1,omitempty" db:"address_line_1"`
	AddressLine2 *string `json:"address_line_2,omitempty" db:"address_line_2"`
	City         string  `json:"city,omitempty" db:"city"`
	State        *string `json:"state,omitempty" db:"state"`
	PostalCode   string  `json:"postal_code,omitempty" db:"postal_code"`
	Country      string  `json:"country,omitempty" db:"country"`

	Status string `json:"status,omitempty" db:"status"`

	CreatedAt *string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt *string `json:"updated_at,omitempty" db:"updated_at"`
}
