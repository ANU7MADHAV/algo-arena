package services

import (
	"context"
	"errors"
	"time"

	"github.com/ANU7MADHAV/algo-arena/models"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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
	ctx,cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()

	existingUser := &models.User{}
	err := s.collection.FindOne(ctx,bson.M{
		"$or" : []bson.M {
			{"email" : user.Email},
			{"username" : user.Username},
		},
	}).Decode(existingUser)

	if err != nil && err != mongo.ErrNoDocuments {
		return nil,err
	}
	if existingUser.ID != primitive.NilObjectID {
return nil, errors.New("user with this email or username already exists")
	}

	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)

	if err != nil {
		return nil,err
	}

	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Role = "user"

	result,err := s.collection.InsertOne(ctx,user)

	if err != nil {
		return nil,err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)


	user.Password = ""

	return user,nil
}


func (s *AuthService) LoginUser(credentials *models.User) (*models.AuthResponse,error) {
	ctx,cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)

	defer cancel()

	user := &models.User{}

	err := s.collection.FindOne(ctx,bson.M{
		"$or" : []bson.M {
			{"email" : user.Email},
			{"username" : user.Username},
		},
	}).Decode(user)

	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil,err
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(credentials.Password))
	if err != nil {
		return nil,errors.New("invalid credentials")
	}

	token,err := s.generateJWT(user)

	if err != nil {
		return nil,err
	}

	user.Password = ""

	return &models.AuthResponse {
		User: user,
		Token: token,
	},nil
}

func (s *AuthService) generateJWT (user *models.User) (string,error){
	claims := jwt.MapClaims{
"id" : user.ID,
"username" : user.Username,
"email" : user.Email,
"role" : user.Role,
"exp" : time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString([]byte(s.jwtSecret))
}