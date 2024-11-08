package usecase

import (
	"context"
	"log"

	"github.com/ANU7MADHAV/algo-arena/models"
	domain2 "github.com/ANU7MADHAV/algo-arena/repository"
)

type UserSeviceImpl struct {
	userRepo domain2.UserRepository
	ctx      context.Context
}

func (u UserSeviceImpl) CreateUser(ctx context.Context, req *models.User) error {
	if ctx == nil {
		ctx = context.Background()
	}

	user, err := u.userRepo.InsertData(ctx)

	if err != nil {
		return err
	}

	log.Println("Successfully Insert Data of User")

	return user, nil
}
