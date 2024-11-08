package controllers

import (
	"log"
	"net/http"

	"github.com/ANU7MADHAV/algo-arena/services"
	"github.com/gin-gonic/gin"
)

var user services.User

func GetAllUsers(c *gin.Context) {
	users, err := user.GetAllUsers()

	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(200, gin.H{"message": users})
}

func CreateUsers(c *gin.Context) {
	var entry services.User
	if err := c.ShouldBindJSON(&entry); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	users, err := user.CreateUser(entry)

	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, users)
}
