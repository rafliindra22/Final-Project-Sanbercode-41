package models

import "time"

type Brand struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserID      int       `json:"user_id"`
	User        User      `json:"-"`
	Phones      []Phone   `json:"-"`
}

type BrandInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
