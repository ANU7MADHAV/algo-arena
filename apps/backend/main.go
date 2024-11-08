package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ANU7MADHAV/algo-arena/config"
	"github.com/ANU7MADHAV/algo-arena/delivery/http"
	domain "github.com/ANU7MADHAV/algo-arena/domain/usecase"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

var (
	server      *gin.Engine
	us          domain.UserUseCase
	uc          http.UserController
	ctx         context.Context
	mongoClient *mongo.Client
)

func init() {
	ctx = context.TODO()

	// mongo

	mongoClient, err := config.Connect()

	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Mongo connected")
	us := domain.NewUserUseCase(mongoClient)
	uc = http.NewUserController(us)

	server = gin.Default()

}

func main() {
	defer func(mongoClient *mongo.Client, ctx context.Context) {
		err := mongoClient.Disconnect(ctx)
		if err != nil {
			fmt.Println("error")
			return
		}
	}(mongoClient, ctx)
	basepath := server.Group("/v1")
	uc.RegisterRoutes(basepath)

	log.Fatal(server.Run(":3000"))
}
