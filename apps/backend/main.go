package main

import (
	"context"
	"log"
	"time"

	"github.com/ANU7MADHAV/algo-arena/db"
	"github.com/ANU7MADHAV/algo-arena/routes"
	"github.com/ANU7MADHAV/algo-arena/services"
)

func main() {
	mongoClient, err := db.ConnectMongo()

	if err != nil {
		log.Panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()

	r := routes.SetupRoutes()

	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	services.New(mongoClient)

	r.Run()
}
