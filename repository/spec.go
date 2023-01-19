package repository

import (
	"Final-Project-Sanbercode-Go-Batch-41/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllSpec(db *sql.DB) (result []models.Spec, count int, err error) {
	sql := "SELECT * FROM spec"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var spec = models.Spec{}

		err = rows.Scan(&spec.ID, &spec.PhoneID, &spec.Processor, &spec.Memory, &spec.Storage, &spec.Screen, &spec.Camera, &spec.Price, &spec.Review, &spec.CreatedAt, &spec.UpdatedAt, &spec.UserID)
		if err != nil {
			panic(err)
		}

		result = append(result, spec)
	}
	count = len(result)

	return
}

func GetSpecByID(db *sql.DB, id int) (result models.Spec, err error) {
	result = models.Spec{}

	row := db.QueryRow("SELECT * FROM spec WHERE id = $1", id)
	err = row.Scan(&result.ID, &result.PhoneID, &result.Processor, &result.Memory, &result.Storage, &result.Screen, &result.Camera, &result.Price, &result.Review, &result.CreatedAt, &result.UpdatedAt, &result.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Record Not FOund"})
	}
	return
}

func GetSpecByPhoneID(db *sql.DB, spec models.Spec) (result []models.Spec, err error) {
	sql := "SELECT * FROM spec WHERE phone_id = $1"

	row, err := db.Query(sql, spec.PhoneID)
	if err != nil {
		panic(err)
	}

	defer row.Close()

	for row.Next() {
		var spec = models.Spec{}

		err = row.Scan(&spec.ID, &spec.PhoneID, &spec.Processor, &spec.Memory, &spec.Storage, &spec.Screen, &spec.Camera, &spec.Price, &spec.Review, &spec.CreatedAt, &spec.UpdatedAt, &spec.UserID)
		if err != nil {
			panic(err)
		}

		result = append(result, spec)
	}

	return
}

func GetCommentByPhoneID(db *sql.DB, comment models.CommentsPhone) (result []models.CommentsPhone, err error) {
	sql := "SELECT * FROM commentphone WHERE phone_id = $1"

	row, err := db.Query(sql, comment.PhoneID)
	if err != nil {
		panic(err)
	}

	defer row.Close()

	for row.Next() {
		var comment = models.CommentsPhone{}

		err = row.Scan(&comment.ID, &comment.PhoneID, &comment.Name, &comment.Comment, &comment.CreatedAt, &comment.UserID)
		if err != nil {
			panic(err)
		}

		result = append(result, comment)
	}
	return
}

func DeleteSpec(db *sql.DB, spec models.Spec) (err error) {
	sql := "DELETE FROM spec WHERE id = $1"

	errs := db.QueryRow(sql, spec.ID)

	return errs.Err()
}
