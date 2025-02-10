package controllers

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

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

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Panic("Error while verifying the connection:", err)
	}

	db.client = client
	return nil
}

func (db *DB) CreateJobListing(input model.CreateJobListingInput) (*model.JobListing, error) {
	var jobListing model.JobListing
	return &jobListing, nil
}

func (db *DB) UpdateJobListing(
	id string,
	input model.UpdateJobListingInput,
) (*model.JobListing, error) {
	var updatedJoblisting model.JobListing
	return &updatedJoblisting, nil
}

func (db *DB) DeleteJobListing(id string) (*model.DeleteJobListingResponse, error) {
	var response model.DeleteJobListingResponse
	return &response, nil
}

func (db *DB) GetJobs() ([]*model.JobListing, error) {
	var jobListings []*model.JobListing
	return jobListings, nil
}

func (db *DB) GetJob(id string) (*model.JobListing, error) {
	var jobListing model.JobListing
	return &jobListing, nil
}
