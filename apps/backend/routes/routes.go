package routes

import (
	"github.com/gin-gonic/gin"
)


func SetupRoutes (r *gin.Engine) {
 auth := r.Group("/auth")

 {
	auth.POST("/register")
	auth.POST("/login")
 }

 problems := r.Group("/problems")

 {
	problems.GET("/")
	problems.GET("/:id")
	problems.POST("/")
	problems.PUT("/:id")
	problems.DELETE("/:id")
 }

 submissions := r.Group("/submissions")

 {
	submissions.GET("/")
	submissions.GET("/:id")
	submissions.POST("/")
 }
}