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

func GetAllProblems(c *gin.Context) {
	var entry services.Problem

	problemService := &services.Problem{}

	if err := c.ShouldBindJSON(&entry); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	problems, err := problemService.ListProblems()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("problems", problems)
	c.JSON(http.StatusOK, problems)
}

func GetProblemById(c *gin.Context) {
	id := c.Param("id")

	problemService := &services.Problem{}

	problem, err := problemService.GetProblemById(id)

	if err != nil {
		log.Fatal(err)
		c.JSON(404, gin.H{"message": "Problem not found"})
	}

	fmt.Println("problem", problem)

	c.JSON(http.StatusOK, problem)
}
