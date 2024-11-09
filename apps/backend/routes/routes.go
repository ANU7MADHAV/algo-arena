package routes

import (
	"github.com/ANU7MADHAV/algo-arena/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")

	{
		v1.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "Hitted"})
		})
		v1.GET("/users", controllers.GetAllUsers)
		v1.POST("/users", controllers.CreateUsers)
		v1.PUT("/users/:id", controllers.UpdateUser)

		v1.POST("/problems", controllers.CreateProblem)
		v1.GET("/problems", controllers.GetAllProblems)
		v1.GET("/problems/:id", controllers.GetProblemById)

		v1.POST("/submission", controllers.CreateSubmission)
	}
	return r
}
