package models

import (
	"database/sql"
	"html"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	FullAccess bool      `json:"full_access"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

func (u *User) SaveUser(db *sql.DB) (*User, error) {
	//turn password into hash
	hashedPassword, errPassword := u.GeneratePassword(u.Password)
	if errPassword != nil {
		return &User{}, errPassword
	}
	u.Password = string(hashedPassword)
	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	_, err := db.Exec("INSERT INTO users (name, username, password, full_acces, created_at, updated_at ) VALUES ($1, $2, $3, $4, $5, $6)", u.Name, u.Username, u.Password, u.FullAccess, u.Created_at, u.Updated_at)
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) GeneratePassword(pass string) (string, error) {
	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if errPassword != nil {
		return "", errPassword
	}
	return string(hashedPassword), nil
}
