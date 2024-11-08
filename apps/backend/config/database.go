package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Client, error) {
	connpattern := "mongodb+srv://anumadhavan888:jePmah6FIWKIVvjn@cluster0.sym0l.mongodb.net/"

	clientOptions := options.Client().ApplyURI(connpattern)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to mongoDB!!!")
	}

	log.Println("MongoDB connected")
	return client, nil
}
