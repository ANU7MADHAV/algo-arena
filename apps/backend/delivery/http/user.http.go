package http

import (
	"fmt"
	"net/http"

	"github.com/ANU7MADHAV/algo-arena/delivery/response"
	domain "github.com/ANU7MADHAV/algo-arena/domain/usecase"
	"github.com/ANU7MADHAV/algo-arena/models"

	// "github.com/ANU7MADHAV/algo-arena/models"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUseCase domain.UserUseCase
}

func New(userSevice domain.UserUseCase) UserController {
	return UserController{
		UserUseCase: userSevice,
	}
}

func (uc *UserController) createUser(ctx *gin.Context) {
	var user models.User

	fmt.Println("user")

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}
	err := uc.UserUseCase.CreateUser(ctx, &user)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response.UserResponse{Status: http.StatusOK, Message: "successfull", Data: ""})
}
