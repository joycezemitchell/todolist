package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database - Mongo DB collection global variable
var Database *mongo.Database

func init() {

	godotenv.Load("/var/www/todo.allyapps.com/todo.env")

	// connect to MongoDB
	fmt.Println("Connecting to MongoDB")
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + os.Getenv("DBUSERNAME") + ":" + os.Getenv("DBPASSWORD") + "@" + os.Getenv("MONGOIP") + ":" + os.Getenv("MONGOPORT")))

	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	Database = client.Database("Todo")
}
