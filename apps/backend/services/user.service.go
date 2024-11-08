package services

import (
	"context"
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

func New(mongo *mongo.Client) User {
	client = mongo
	return User{}
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
