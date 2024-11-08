package domain

import (
	"context"

	"github.com/ANU7MADHAV/algo-arena/delivery/http"
	"github.com/ANU7MADHAV/algo-arena/models"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, req *models.User) error
	GetUser(ctx context.Context, req *string) (*models.User, error)
	GetAll(ctx context.Context) ([]models.User, error)
	UpdateUser(ctx context.Context, req *models.User) error
	DeleteUser(ctx context.Context, req *string) error
}

func NewUserController(us UserUseCase) http.UserController {
	return http.UserController{
		UserUseCase: us,
	}
}
