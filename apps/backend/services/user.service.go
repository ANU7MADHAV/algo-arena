package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Username  string    `json:"username,omitempty" bson:"username,omitempty"`
	Email     string    `json:"email,omitempty" bson:"email,omitempty"`
	Password  string    `json:"password,omitempty" bson:"password,omitempty"`
	Role      string    `json:"role,omitempty" bson:"role,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

var client *mongo.Client

func New(mongo *mongo.Client) {
	client = mongo
}

func returnCollectPointer(collection string) *mongo.Collection {

	return client.Database("algo_arena").Collection(collection)
}

func (u *User) GetAllUsers() ([]User, error) {
	collection := returnCollectPointer("users")

	var users []User

	cursor, err := collection.Find(context.Background(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user User

		cursor.Decode(&user)
		users = append(users, user)
	}

	return users, nil
}

func (u *User) ChecKUser(entry User) (User, error) {
	collection := returnCollectPointer("users")

	// Check if user already exists
	var existingUser User
	err := collection.FindOne(context.Background(), bson.D{{Key: "email", Value: entry.Email}}).Decode(&existingUser)

	// If user is found, return error
	if err == nil {
		return User{}, fmt.Errorf("user with email %s already exists", entry.Email)
	}

	return existingUser, nil

}

func (u *User) CreateUser(entry User) (User, error) {
	collection := returnCollectPointer("users")

	entry.Role = "user"
	entry.CreatedAt = time.Now()
	entry.UpdatedAt = time.Now()

	insertResult, err := collection.InsertOne(context.TODO(), entry)

	if err != nil {
		log.Println("Error:", err)
		return User{}, err
	}

	entry.ID = insertResult.InsertedID.(primitive.ObjectID).Hex()
	return entry, nil
}

func (u *User) UpdateUser(id string, entry User) (User, error) {
	collection := returnCollectPointer("users")

	mongoId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.D{{Key: "_id", Value: mongoId}}

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "username", Value: entry.Username},
		{Key: "email", Value: entry.Email},
		{Key: "password", Value: entry.Password},
	}}}

	result, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		panic(err)
	}

	if result.MatchedCount == 0 {
		return User{}, errors.New("user not found")
	} else if result.ModifiedCount == 0 {
		log.Println("No modifications were made to the user data")

	}
	return entry, nil
}
