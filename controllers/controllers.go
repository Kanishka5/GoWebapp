package controllers

import (
	"net/http"

	"../models"
	"github.com/gin-gonic/gin"
)

// Home view render
func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}

// UserForm : renders a user form template
func UserForm(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", nil)
}

// AllUser : retuns all users
func AllUser(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// CreateUser : create new user
func CreateUser(c *gin.Context) {
	// Validate input
	var input models.User
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create new user
	user := models.User{Name: input.Name, Email: input.Email}
	models.DB.Create(&user)

	// return the user json
	c.JSON(http.StatusOK, gin.H{"users": user})
}

// FindUser : search user by primary key
func FindUser(c *gin.Context) {
	id := c.Param("id")
	var users []models.User
	models.DB.First(&users, id)
	c.JSON(http.StatusOK, gin.H{"users": users})
}
