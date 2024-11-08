package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/ANU7MADHAV/algo-arena/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	collection *mongo.Collection
	timeout    time.Duration
}

func NewRepository(db *mongo.Database, timeout time.Duration) *UserRepository {
	return &UserRepository{
		collection: db.Collection("users"),
		timeout:    timeout,
	}
}

func (u UserRepository) InsertData(ctx context.Context) (user *models.User, err error) {

	ctx, canecel := context.WithTimeout(ctx, 5*time.Microsecond)
	defer canecel()

	opts := options.InsertOne().SetComment("user data is inserting")
	result, err := u.collection.InsertOne(ctx, user, opts)

	fmt.Println("result", result)

	if err != nil {
		return nil, fmt.Errorf("failed to insert user %v: %w", user, err)
	}

	if iod, ok := result.InsertedID.(primitive.ObjectID); ok {
		user.ID = iod
	}

	return user, nil

}
