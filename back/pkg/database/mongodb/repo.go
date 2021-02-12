package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// the uri for connecting to mongoDB
const mongoURI = "mongodb://localhost:27017"

// time outs for queries
const ShortTimeOut = 5000
const LongTimeOut = 100000

// this client is shared for every one who wants to access to mongo
var client *mongo.Client

// connecting to mongo and retrieving a mongoClient
func init() {
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = mongoClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	client = mongoClient

}

func GetClient() *mongo.Client {
	return client
}