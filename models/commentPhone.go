package models

import "time"

type CommentsPhone struct {
	ID        int       `json:"id"`
	PhoneID   int       `json:"phone_id"`
	Name      string    `json:"name"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UserID    int       `json:"user_id"`
	User      User      `json:"-"`
	Phone     Phone     `json:"-"`
}

type CommentPhoneInput struct {
	PhoneID int    `json:"phone_id"`
	Comment string `json:"comment"`
}
