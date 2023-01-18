package controllers

import (
	"Final-Project-Sanbercode-Go-Batch-41/database"
	"Final-Project-Sanbercode-Go-Batch-41/models"
	"Final-Project-Sanbercode-Go-Batch-41/repository"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateSpec(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var input models.SpecInput
	var data models.Spec
	var rating models.Phone

	// validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get user info
	user, _ := c.Get("user")
	users := user.(models.User)

	// get phone id
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

	// create spec
	spec := models.Spec{
		PhoneID:   input.PhoneID,
		Processor: input.Processor,
		Memory:    input.Memory,
		Storage:   input.Storage,
		Screen:    input.Screen,
		Camera:    input.Camera,
		Price:     input.Price,
		Review:    input.Review,
		UserID:    users.ID,
	}
	fmt.Sprintln(spec)
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Time{}

	stmt, err := db.Prepare("INSERT INTO spec (phone_id, processor, memory, storage, screen, camera, price, review, created_at, updated_at, user_id) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// execute query
	res, err := stmt.Exec(input.PhoneID, input.Processor, input.Memory, input.Storage, input.Screen, input.Camera, input.Price, input.Review, data.CreatedAt, data.UpdatedAt, users.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Sprintln(res)

	c.JSON(http.StatusOK, gin.H{"data": input, "status": "insert success"})

}
func UpdateSpec(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var spec models.Spec
	var input models.SpecInput
	var phone models.Phone

	spec.ID, _ = strconv.Atoi(c.Param("id"))

	// Get model if exist
	row := db.QueryRow("SELECT * FROM spec WHERE id = $1", spec.ID)
	err := row.Scan(&spec.ID, &spec.PhoneID, &spec.Processor, &spec.Memory, &spec.Storage, &spec.Screen, &spec.Camera, &spec.Price, &spec.Review, &spec.CreatedAt, &spec.UpdatedAt, &spec.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user active info
	user, _ := c.Get("user")
	users := user.(models.User)

	spec.UpdatedAt = time.Now()

	// get phone id
	rows := db.QueryRow("SELECT * FROM phone WHERE id = $1", input.PhoneID)
	err = rows.Scan(&phone.ID, &phone.Type, &phone.Year, &phone.BrandID, &phone.CreatedAt, &phone.UpdatedAt, &phone.EditorName, &phone.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": "PhoneID not found!"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Prepare the update statement
	stmt, err := db.Prepare("UPDATE spec SET phone_id = $1, processor = $2, memory = $3, storage = $4, screen = $5, camera = $6, price = $7, review = $8, updated_at = $9, user_id = $10 WHERE id = $11")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Execute the update statement
	_, err = stmt.Exec(input.PhoneID, input.Processor, input.Memory, input.Storage, input.Screen, input.Camera, input.Price, input.Review, spec.UpdatedAt, users.ID, spec.ID)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"data": input, "status": "Update success"})
}

func GetAllSpec(c *gin.Context) {
	var result gin.H

	phone, _, err := repository.GetAllSpec(database.DbConnection)

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

func GetSpecByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
		return
	}

	result, err := repository.GetSpecByID(database.DbConnection, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func GetSpecCommentByPhoneID(c *gin.Context) {
	var spec models.Spec
	var comment models.CommentsPhone
	var result gin.H
	var results gin.H

	id, _ := strconv.Atoi(c.Param("id"))
	spec.PhoneID = id
	comment.PhoneID = id

	comments, err := repository.GetCommentByPhoneID(database.DbConnection, comment)
	if err != nil {
		results = gin.H{
			"result": err,
		}
	} else {
		results = gin.H{
			"result": comments,
		}
	}

	specs, err := repository.GetSpecByPhoneID(database.DbConnection, spec)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": specs,
		}
	}

	c.JSON(http.StatusOK, gin.H{"comment": results, "spec": result})
}
