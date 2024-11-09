package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ANU7MADHAV/algo-arena/services"
	"github.com/gin-gonic/gin"
)

func CreateProblem(c *gin.Context) {
	var entry services.Problem
	problemService := &services.Problem{}

	if err := c.ShouldBindJSON(&entry); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	problem, err := problemService.CreateProblem(entry)

	if err != nil {
		log.Fatalf("error")
	}

	fmt.Println("Problem", problem)
	c.JSON(http.StatusOK, problem)
}
