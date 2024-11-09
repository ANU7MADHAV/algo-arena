package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ANU7MADHAV/algo-arena/services"
	"github.com/gin-gonic/gin"
)

func CreateSubmission(c *gin.Context) {
	var entry services.Submission
	submissionService := &services.Submission{}

	if err := c.ShouldBindJSON(&entry); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"Error": err.Error()})
		return
	}

	submission, err := submissionService.CreateSubmission(entry)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("submission", submission)

	c.JSON(http.StatusOK, submission)
}

func GetAllSubmissions(c *gin.Context) {
	submissionsService := &services.Submission{}

	submissions, err := submissionsService.GetAllSubmissions()

	fmt.Println("submissions", submissions)

	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, submissions)
}

func GetSubmissionById(c *gin.Context) {

	submissionService := &services.Submission{}
	id := c.Param("id")

	submission, err := submissionService.GetSubmissionById(id)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, submission)
}
