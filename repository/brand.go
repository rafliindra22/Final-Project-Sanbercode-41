package repository

import (
	"Final-Project-Sanbercode-Go-Batch-41/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllBrand(db *sql.DB) (result []models.Brand, count int, err error) {
	sql := "SELECT * FROM brand"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var brand = models.Brand{}

		err = rows.Scan(&brand.ID, &brand.Name, &brand.Description, &brand.CreatedAt, &brand.UpdatedAt, &brand.UserID)
		if err != nil {
			panic(err)
		}

		result = append(result, brand)
	}
	count = len(result)

	return
}

func GetBrandByID(db *sql.DB, id int) (result models.Brand, err error) {
	result = models.Brand{}

	row := db.QueryRow("SELECT * FROM brand WHERE id = $1", id)
	err = row.Scan(&result.ID, &result.Name, &result.Description, &result.CreatedAt, &result.UpdatedAt, &result.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Record Not FOund"})
	}
	return
}
