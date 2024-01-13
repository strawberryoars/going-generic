package clients

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client

// Initializes MongoDB Client
// Reads dotenv file to load the database URI
func InitMongoConnection() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	uri := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(uri)
	Client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	err = Client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
}
