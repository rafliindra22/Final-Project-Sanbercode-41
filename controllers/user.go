package controllers

import (
	"Final-Project-Sanbercode-Go-Batch-41/models"
	"Final-Project-Sanbercode-Go-Batch-41/utils/token"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string, db *sql.DB) (string, error) {

	var err error

	u := models.User{}

	rows := db.QueryRow("SELECT * FROM users WHERE username = $1", username)
	err = rows.Scan(&u.ID, &u.Name, &u.Username, &u.Password, &u.FullAccess, &u.Created_at, &u.Updated_at)

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	token, err := token.GenerateToken(u.ID, u.FullAccess)

	if err != nil {
		return "", err
	}

	return token, nil

}
