package controllers

import (
	"Final-Project-Sanbercode-Go-Batch-41/database"
	"Final-Project-Sanbercode-Go-Batch-41/models"
	"Final-Project-Sanbercode-Go-Batch-41/repository"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var input models.CommentPhoneInput
	var data models.CommentsPhone
	var rating models.Phone

	// validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get user info
	user, _ := c.Get("user")
	users := user.(models.User)

	row := db.QueryRow("SELECT * FROM phone WHERE id = $1", input.PhoneID)
	err := row.Scan(&rating.ID, &rating.Type, &rating.Year, &rating.BrandID, &rating.CreatedAt, &rating.UpdatedAt, &rating.EditorName, &rating.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": "PhoneID not found!"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	data.CreatedAt = time.Now()

	// create spec
	comment := models.CommentsPhone{
		PhoneID: input.PhoneID,
		Name:    users.Name,
		Comment: input.Comment,
		UserID:  users.ID,
	}
	fmt.Sprintln(comment)

	stmt, err := db.Prepare("INSERT INTO commentphone (phone_id, name, comment, created_at, user_id) VALUES ($1,$2,$3,$4,$5)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res, err := stmt.Exec(rating.ID, users.Name, input.Comment, data.CreatedAt, users.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Sprintln(res)

	c.JSON(http.StatusOK, gin.H{"data": input, "status": "comment success"})
}

func GetAllComment(c *gin.Context) {
	var result gin.H

	phone, _, err := repository.GetAllComment(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": phone,
		}
	}

	c.JSON(http.StatusOK, result)
}
