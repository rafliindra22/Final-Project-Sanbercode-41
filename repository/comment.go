package repository

import (
	"Final-Project-Sanbercode-Go-Batch-41/models"
	"database/sql"
)

func GetAllComment(db *sql.DB) (result []models.CommentsPhone, count int, err error) {
	sql := "SELECT * FROM commentphone"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var comment = models.CommentsPhone{}

		err = rows.Scan(&comment.ID, &comment.PhoneID, &comment.Name, &comment.Comment, &comment.CreatedAt, &comment.UserID)
		if err != nil {
			panic(err)
		}

		result = append(result, comment)
	}
	count = len(result)

	return
}

func DeleteComment(db *sql.DB, comment models.CommentsPhone) (err error) {
	sql := "DELETE FROM commentphone WHERE id = $1"

	errs := db.QueryRow(sql, comment.ID)

	return errs.Err()
}
