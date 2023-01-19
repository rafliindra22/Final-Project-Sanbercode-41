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

func CreatePhone(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var input models.PhoneInput
	var data models.Phone
	var rating models.Brand

	// validate input
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get brand id
	row := db.QueryRow("SELECT * FROM brand WHERE id = $1", input.BrandID)
	err = row.Scan(&rating.ID, &rating.Name, &rating.Description, &rating.CreatedAt, &rating.UpdatedAt, &rating.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": "BrandID not found!"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Get user active info
	user, _ := c.Get("user")
	users := user.(models.User)

	// Create Phone
	phone := models.Phone{UserID: users.ID, EditorName: users.Username, Type: input.Type, Year: input.Year, BrandID: input.BrandID}
	fmt.Sprintln(phone)
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Time{}

	stmt, err := db.Prepare("INSERT INTO phone (type, year, brand_id, created_at, updated_at, editor_name, user_id) VALUES ($1,$2,$3,$4,$5,$6,$7)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// execute query
	res, err := stmt.Exec(input.Type, input.Year, input.BrandID, data.CreatedAt, data.UpdatedAt, users.Username, users.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Sprintln(res)

	c.JSON(http.StatusCreated, gin.H{"data": input, "status": "insert success"})
}

func GetAllPhone(c *gin.Context) {
	var result gin.H

	phone, _, err := repository.GetAllPhone(database.DbConnection)

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

func GetPhoneById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
		return
	}

	result, err := repository.GetPhoneById(database.DbConnection, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func UpdatePhone(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	var phone models.Phone
	var input models.PhoneInput
	var merk models.Brand

	phone.ID, _ = strconv.Atoi(c.Param("id"))

	// Get model if exist
	row := db.QueryRow("SELECT * FROM phone WHERE id = $1", phone.ID)
	err := row.Scan(&phone.ID, &phone.Type, &phone.Year, &phone.BrandID, &phone.CreatedAt, &phone.UpdatedAt, &phone.EditorName, &phone.UserID)
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

	phone.UpdatedAt = time.Now()

	// get brand id
	rows := db.QueryRow("SELECT * FROM brand WHERE id = $1", input.BrandID)
	err = rows.Scan(&merk.ID, &merk.Name, &merk.Description, &merk.CreatedAt, &merk.UpdatedAt, &merk.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": "BrandID not found!"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Prepare the update statement
	stmt, err := db.Prepare("UPDATE phone SET type = $1, year = $2, brand_id = $3, updated_at = $4, editor_name = $5, user_id = $6 WHERE id = $7")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Execute the update statement
	_, err = stmt.Exec(input.Type, input.Year, input.BrandID, phone.UpdatedAt, users.Username, users.ID, phone.ID)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"data": input, "status": "Update success"})
}

func GetPhonesByBrandID(c *gin.Context) {
	var phone models.Phone
	var result gin.H

	id, _ := strconv.Atoi(c.Param("id"))
	phone.BrandID = id

	phones, err := repository.GetPhonesByBrandID(database.DbConnection, phone)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": phones,
		}
	}

	c.JSON(http.StatusOK, result)
}

func DeletePhone(c *gin.Context) {
	var phone models.Phone

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
		return
	}

	phone.ID = id

	errs := repository.DeletePhone(database.DbConnection, phone)
	if errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone with that ID doesn't exist"})
	}
	c.JSON(http.StatusOK, gin.H{

		"result": "Success Delete Phone",
	})
}
