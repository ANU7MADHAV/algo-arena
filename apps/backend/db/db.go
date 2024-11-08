package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://anumadhavan888:jePmah6FIWKIVvjn@cluster0.sym0l.mongodb.net/"))
	if err != nil {
		return client, err
	}
	log.Println("Connected to mongo...")

	return client, nil
}
