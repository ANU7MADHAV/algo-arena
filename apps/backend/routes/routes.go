package routes

import (
	"github.com/ANU7MADHAV/algo-arena/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"Ping": "Pong"})
	})

	v1 := r.Group("/v1")

	{
		v1.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "Hitted"})
		})

		v1.GET("/users", controllers.GetAllUsers)
		v1.GET("/users/:id")
		v1.POST("/users/create", controllers.CreateUsers)
		v1.PUT("/users/:id", controllers.UpdateUser)

		v1.POST("/problems/create", controllers.CreateProblem)
		v1.GET("/problems", controllers.GetAllProblems)
		v1.GET("/problems/:id", controllers.GetProblemById)

		v1.POST("/submissions/create", controllers.CreateSubmission)
		v1.GET("/submissions", controllers.GetAllSubmissions)
		v1.GET("/submissions/:id", controllers.GetSubmissionById)

		v1.GET("/reset", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "Done"})
		})
	}
	return r
}
