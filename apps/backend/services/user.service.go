package services

import (
	"github.com/ANU7MADHAV/algo-arena/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthService struct {
	collection *mongo.Collection
	jwtSecret string
}

func NewAuthService (db *mongo.Database, jwtSecret string) *AuthService {
return &AuthService{
	collection: db.Collection("users"),
	jwtSecret: jwtSecret,
}
}

func (s *AuthService) CreateUser(user *models.User) (*models.User,error) {
	
}