package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

var err error
var BlogCol *mongo.Collection

func ConnectToDB() {
	//load the environment variables
	LoadEnvVariables()

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("DB_URL")).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	Client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	BlogCol = Client.Database(os.Getenv("DB_NAME")).Collection("Posts")

	//set context with timeout to disconnect after 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Hour)
	if err = Client.Disconnect(ctx); err != nil {
		panic(err)
	}
	fmt.Println("Disconnected from the database")
	defer cancel()

	// Send a ping to confirm a successful connection
	if err := Client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

}
