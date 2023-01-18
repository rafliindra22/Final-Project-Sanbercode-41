package repository

import (
	"Final-Project-Sanbercode-Go-Batch-41/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	c *gin.Context
)

func GetAllPhone(db *sql.DB) (result []models.Phone, count int, err error) {
	sql := "SELECT * FROM phone"

	rows, err := db.Query(sql)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Record Not FOund"})
	}
	defer rows.Close()

	for rows.Next() {
		var phone = models.Phone{}

		err = rows.Scan(&phone.ID, &phone.Type, &phone.Year, &phone.BrandID, &phone.CreatedAt, &phone.UpdatedAt, &phone.EditorName, &phone.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Record Not FOund"})
		}

		result = append(result, phone)
	}
	count = len(result)

	return
}

func GetPhoneById(db *sql.DB, id int) (result models.Phone, err error) {
	result = models.Phone{}

	row := db.QueryRow("SELECT * FROM phone WHERE id = $1", id)
	err = row.Scan(&result.ID, &result.Type, &result.Year, &result.BrandID, &result.CreatedAt, &result.UpdatedAt, &result.EditorName, &result.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Record Not FOund"})
	}
	return
}

func GetPhonesByBrandID(db *sql.DB, phone models.Phone) (result []models.Phone, err error) {
	sql := "SELECT * FROM phone WHERE brand_id = $1"

	rows, err := db.Query(sql, phone.BrandID)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var phone = models.Phone{}

		err = rows.Scan(&phone.ID, &phone.Type, &phone.Year, &phone.BrandID, &phone.CreatedAt, &phone.UpdatedAt, &phone.EditorName, &phone.UserID)
		if err != nil {
			panic(err)
		}

		result = append(result, phone)
	}
	return
}
