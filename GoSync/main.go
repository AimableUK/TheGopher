package main

import (
	"context"
	"log"

	"github.com/AimableUK/TheGopher/GoSync/controllers"
	"github.com/AimableUK/TheGopher/GoSync/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	// 1. Setup Context
	ctx := context.TODO()

	// 2. Setup DB connection (Local variable, NOT global)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer mongoclient.Disconnect(ctx)

	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	// 3. Dependency Injection (Wiring everything together)
	userCollection := mongoclient.Database("userdb").Collection("users")

	// Inject the collection into the service
	userService := services.NewUserService(userCollection, ctx)

	// Inject the service into the controller
	userController := controllers.New(userService)

	// 4. Setup Server
	server := gin.Default()
	basePath := server.Group("/v1")
	userController.RegisterUserRoutes(basePath)

	log.Fatal(server.Run(":9090"))
}
