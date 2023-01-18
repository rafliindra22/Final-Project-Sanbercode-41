package middlewares

import (
	"Final-Project-Sanbercode-Go-Batch-41/database"
	"Final-Project-Sanbercode-Go-Batch-41/models"
	"database/sql"
	"log"
	"net/http"

	"Final-Project-Sanbercode-Go-Batch-41/utils/token"

	"github.com/gin-gonic/gin"
)

func PublicMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		id, _ := token.ExtractTokenID(c)
		db := c.MustGet("db").(*sql.DB)
		row := db.QueryRow("SELECT * FROM users WHERE id = $1", id)
		var user models.User
		err = row.Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.FullAccess, &user.Created_at, &user.Updated_at)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Fatal("User not found")
			} else {
				log.Fatal(err)
			}
		}
		c.Set("user", user)

		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		id, _ := token.ExtractTokenID(c)
		db := database.DbConnection
		row := db.QueryRow("SELECT * FROM users WHERE id = $1", id)
		var user models.User
		err = row.Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.FullAccess, &user.Created_at, &user.Updated_at)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Fatal("User not found")
			} else {
				log.Fatal(err)
			}
		}
		c.Set("user", user)

		c.Next()

		// Validate access
		if !user.FullAccess {
			c.JSON(http.StatusBadRequest, gin.H{"error": "you dont have an access!"})
			return
		}
		c.Set("user", user)

		c.Next()
	}
}
