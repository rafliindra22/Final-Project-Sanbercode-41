package models

import "time"

type Phone struct {
	ID         int             `json:"id"`
	Type       string          `json:"type"`
	Year       int             `json:"year"`
	BrandID    int             `json:"brand_id"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
	EditorName string          `json:"editor_name"`
	UserID     int             `json:"user_id"`
	User       User            `json:"-"`
	Brand      Brand           `json:"-"`
	Comments   []CommentsPhone `json:"-"`
}

type PhoneInput struct {
	Type    string `json:"type"`
	Year    int    `json:"year"`
	BrandID int    `json:"brand_id"`
}
