package models

import "database/sql"

type Admin struct {
	ID                int            `json:"id,omitempty" db:"id,omitempty"`
	Username          string         `json:"username,omitempty" db:"username,omitempty"`
	Email             string         `json:"email,omitempty" db:"email,omitempty"`
	Password          string         `json:"password,omitempty" db:"password,omitempty"`
	UserCreatedAt     *string        `json:"user_created_at,omitempty" db:"user_created_at,omitempty"`
	PasswordChangedAt *string        `json:"password_changed_at,omitempty" db:"password_changed_at,omitempty"`
	PasswordOtp       sql.NullString `json:"password_otp,omitempty" db:"password_otp,omitempty"`
	OtpExpires        sql.NullString `json:"otp_expires,omitempty" db:"otp_expires,omitempty"`
	InactiveStatus    bool           `json:"inactive_status,omitempty" db:"inactive_status,omitempty"`
}

type UpdatePasswordRequestAdmins struct {
	Otp         string `json:"otp,omitempty" db:"otp,omitempty"`
	NewPassword string `json:"new_password,omitempty" db:"new_password,omitempty"`
}

type AdminResponse struct {
	ID                int     `json:"id,omitempty" db:"id,omitempty"`
	Username          string  `json:"username,omitempty" db:"username,omitempty"`
	Email             string  `json:"email,omitempty" db:"email,omitempty"`
	Password          string  `json:"password,omitempty" db:"password,omitempty"`
	UserCreatedAt     *string `json:"user_created_at,omitempty" db:"user_created_at,omitempty"`
	PasswordChangedAt *string `json:"password_changed_at,omitempty" db:"password_changed_at,omitempty"`
	InactiveStatus    bool    `json:"inactive_status,omitempty" db:"inactive_status,omitempty"`
}

type AdminUpdateDetail struct {
	FirstName string `json:"first_name,omitempty" db:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty" db:"last_name,omitempty"`
	Email     string `json:"email,omitempty" db:"email,omitempty"`
}

type ConfirmDetailAdmins struct {
	Otp   string `json:"otp,omitempty" db:"otp,omitempty"`
	Email string `json:"email,omitempty" db:"email,omitempty"`
}
