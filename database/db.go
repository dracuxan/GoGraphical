package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/dracuxan/job-listing-api/graph/model"
)

var url string = "mongodb://127.0.0.1:27017/"

type DB struct {
	client *mongo.Client
}

func (db *DB) Connect() error {
	clientOptions := options.Client().ApplyURI(url)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Panic("Error connecting to mongodb:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Panic("Error while verifying the connection:", err)
	}

	db.client = client
	return nil
}

func (db *DB) CreateJobListing(input model.CreateJobListingInput) (*model.JobListing, error) {
	return nil, nil
}
