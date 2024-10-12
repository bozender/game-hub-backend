package main

import (
	"context"
	"fmt"
	"log"

	"game-hub-backend/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var gameCollection *mongo.Collection
var userCollection *mongo.Collection

// MongoDB connection setup
func connectMongo() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://bozender59:KudiKudi59**@gamehub.8xoho.mongodb.net/?retryWrites=true&w=majority&appName=gamehub"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	// Get the database
	database := client.Database("gamehub")

	// Get collections for games and users
	gameCollection = database.Collection("games")
	userCollection = database.Collection("users")
}

// Close MongoDB connection on exit
func disconnectMongo() {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Disconnected from MongoDB!")
}

func main() {
	// Connect to MongoDB
	connectMongo()
	defer disconnectMongo()

	// Setup Gin router
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r)

	// Start the server
	r.Run(":8080") // Default port is 8080
}
