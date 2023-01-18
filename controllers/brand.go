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

func GetAllBrand(c *gin.Context) {
	var result gin.H

	// get query
	categories, _, err := repository.GetAllBrand(database.DbConnection)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": categories,
		}
	}

	c.JSON(http.StatusOK, result)
}

func CreateBrand(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)
	var input models.BrandInput
	var data models.Brand

	// Validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user active info
	user, _ := c.Get("user")
	users := user.(models.User)

	// Create Brand
	merk := models.Brand{UserID: users.ID, Name: input.Name, Description: input.Description}
	fmt.Sprintln(merk)
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Time{}

	stmt, err := db.Prepare("INSERT INTO brand (user_id, name, description) VALUES ($1,$2,$3)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// execute query
	res, err := stmt.Exec(users.ID, input.Name, input.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Sprintln(res)

	c.JSON(http.StatusOK, gin.H{"data": input, "status": "insert success"})

}

func UpdateBrand(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var input models.BrandInput
	var merk models.Brand

	merk.ID, _ = strconv.Atoi(c.Param("id"))

	// Get model if exist
	row := db.QueryRow("SELECT * FROM brand WHERE id = $1", merk.ID)
	err := row.Scan(&merk.ID, &merk.Name, &merk.Description, &merk.CreatedAt, &merk.UpdatedAt, &merk.UserID)
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

	merk.UpdatedAt = time.Now()

	// Prepare the update statement
	stmt, err := db.Prepare("UPDATE brand SET name = $1, description = $2, updated_at = $3, user_id = $4 WHERE id = $5")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Execute the update statement
	_, err = stmt.Exec(input.Name, input.Description, merk.UpdatedAt, users.ID, merk.ID)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"data": input, "status": "Update success"})
}

func GetBrandByID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
		return
	}

	result, err := repository.GetBrandByID(database.DbConnection, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}
