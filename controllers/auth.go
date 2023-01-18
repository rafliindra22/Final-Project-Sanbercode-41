package controllers

import (
	"Final-Project-Sanbercode-Go-Batch-41/database"
	"Final-Project-Sanbercode-Go-Batch-41/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var input models.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{
		Username: input.Username,
		Password: input.Password,
	}

	token, err := LoginCheck(u.Username, u.Password, database.DbConnection)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	user := map[string]string{
		"username": u.Username,
	}

	c.JSON(http.StatusOK, gin.H{"message": "login success", "username": user, "token": token})
}

func Register(c *gin.Context) {
	var input models.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}
	u.Name = input.Name
	u.Username = input.Username
	u.Password = input.Password
	u.FullAccess = false

	_, err := u.SaveUser(database.DbConnection)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	FA := strconv.FormatBool(u.FullAccess)

	user := map[string]string{
		"name":        input.Name,
		"username":    input.Username,
		"full_access": FA,
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success", "username": user})
}

func RegisterAdmin(c *gin.Context) {
	var input models.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Name = input.Name
	u.Username = input.Username
	u.Password = input.Password
	u.FullAccess = true
	u.Created_at = time.Now()
	u.Updated_at = time.Time{}

	_, err := u.SaveUser(database.DbConnection)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var FA string

	if u.FullAccess {
		FA = "true"
	}

	user := map[string]string{
		"name":        input.Name,
		"username":    input.Username,
		"full_access": FA,
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success", "username": user})

}
