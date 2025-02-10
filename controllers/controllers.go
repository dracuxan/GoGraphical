package controllers

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/dracuxan/job-listing-api/graph/model"
)

var url string = "mongodb://127.0.0.1:27017/"

type DB struct {
	client   *mongo.Client
	database *mongo.Database
}

func Connect() *DB {
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

	database := client.Database("GoGraphical")
	return &DB{
		client:   client,
		database: database,
	}
}

func (db *DB) CreateJobListing(jobInfo model.CreateJobListingInput) (*model.JobListing, error) {
	jobCollec := db.database.Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	insert, err := jobCollec.InsertOne(
		ctx,
		bson.M{
			"title":       jobInfo.Title,
			"description": jobInfo.Description,
			"url":         jobInfo.URL,
			"company":     jobInfo.Company,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	insertedID := insert.InsertedID.(primitive.ObjectID).Hex()
	returnJobListing := model.JobListing{
		ID:          insertedID,
		Title:       jobInfo.Title,
		Company:     jobInfo.Company,
		Description: jobInfo.Description,
		URL:         jobInfo.URL,
	}
	return &returnJobListing, nil
}

func (db *DB) UpdateJobListing(
	jobId string,
	jobInfo model.UpdateJobListingInput,
) (*model.JobListing, error) {
	jobCollec := db.database.Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updateJobInfo := bson.M{}

	if jobInfo.Title != "" {
		updateJobInfo["title"] = jobInfo.Title
	}
	if jobInfo.Description != "" {
		updateJobInfo["description"] = jobInfo.Description
	}
	if jobInfo.URL != "" {
		updateJobInfo["url"] = jobInfo.URL
	}

	_id, _ := primitive.ObjectIDFromHex(jobId)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateJobInfo}

	results := jobCollec.FindOneAndUpdate(
		ctx,
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(1),
	)

	var jobListing model.JobListing

	if err := results.Decode(&jobListing); err != nil {
		log.Fatal(err)
	}

	return &jobListing, nil
}

func (db *DB) DeleteJobListing(jobId string) (*model.DeleteJobListingResponse, error) {
	jobCollec := db.database.Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(jobId)
	filter := bson.M{"_id": _id}
	_, err := jobCollec.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	return &model.DeleteJobListingResponse{DeleteJobID: jobId}, nil
}

func (db *DB) GetJobs() ([]*model.JobListing, error) {
	jobColl := db.database.Collection("jobs")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var jobListings []*model.JobListing
	cursor, err := jobColl.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err = cursor.All(context.TODO(), &jobListings); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return jobListings, nil
}

func (db *DB) GetJob(jobID string) (*model.JobListing, error) {
	jobColl := db.database.Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(jobID)
	if err != nil {
		return nil, err
	}

	var jobListing model.JobListing
	err = jobColl.FindOne(ctx, bson.M{"_id": objID}).Decode(&jobListing)
	if err != nil {
		return nil, err
	}

	return &jobListing, nil
}
