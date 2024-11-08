package domain

import (
	"context"

	"github.com/ANU7MADHAV/algo-arena/models"
)

type UserRepository interface {
	GetAllData(ctx context.Context) (user []models.User, err error)
	InsertData(ctx context.Context, req *models.User) error
	UpdateData(ctx context.Context, req *models.User) error
	DeleteUser(ctx context.Context, req *models.User) error
	GetData(ctx context.Context, username *string) (user *models.User, err error)
}
