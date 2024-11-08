package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (uc UserController) RegisterRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	fmt.Println("Routes registered")
	userroute.POST("/create", uc.createUser)

}
