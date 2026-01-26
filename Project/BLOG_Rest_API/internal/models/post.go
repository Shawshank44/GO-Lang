package models

type Post struct {
	ID        int      `json:"id,omitempty" db:"id,omitempty"`
	Username  string   `json:"username,omitempty" db:"username,omitempty"`
	Title     string   `json:"title,omitempty" db:"title,omitempty"`
	Content   string   `json:"content,omitempty" db:"content,omitempty"`
	Tags      []string `json:"tags,omitempty" db:"tags,omitempty"`
	CreatedAt string   `json:"created_at,omitempty" db:"created_at,omitempty"`
}
