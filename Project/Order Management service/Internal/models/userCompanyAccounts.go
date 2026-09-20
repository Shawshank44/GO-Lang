package models

type UserCompanyAccount struct {
	UserID           int64 `json:"user_id" db:"user_id"`
	CompanyAccountID int64 `json:"company_account_id" db:"company_account_id"`

	Role   string `json:"role,omitempty" db:"role"`
	Status string `json:"status,omitempty" db:"status"`

	LinkedAt *string `json:"linked_at,omitempty" db:"linked_at"`
}

type UserAccountResponse struct {
	ID            int64  `json:"id"`
	AccountNumber string `json:"account_number"`
	CompanyName   string `json:"company_name"`
	Role          string `json:"role"`
}

type UserProfileResponse struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`

	Accounts []UserAccountResponse `json:"accounts"`
}
