package models

import "time"

type Spec struct {
	ID        int       `json:"id"`
	PhoneID   int       `json:"phone_id"`
	Processor string    `json:"processor"`
	Memory    string    `json:"memory"`
	Storage   string    `json:"storage"`
	Screen    string    `json:"screen"`
	Camera    string    `json:"camera"`
	Price     string    `json:"price"`
	Review    string    `json:"review"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    int       `json:"user_id"`
	User      User      `json:"-"`
	Phone     Phone     `json:"-"`
}

type SpecInput struct {
	PhoneID   int    `json:"phone_id"`
	Processor string `json:"processor"`
	Memory    string `json:"memory"`
	Storage   string `json:"storage"`
	Screen    string `json:"screen"`
	Camera    string `json:"camera"`
	Price     string `json:"price"`
	Review    string `json:"review"`
}
