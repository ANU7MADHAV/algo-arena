package controllers

import (
	"net/http"

	"github.com/ANU7MADHAV/algo-arena/models"
	"github.com/gin-gonic/gin"
)

func Register (c *gin.Context){
var  user models.User
if err := c.ShouldBindJSON(&user); err != nil {
	c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
	return
}
}