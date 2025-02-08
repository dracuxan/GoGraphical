// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateJobListingInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
	URL         string `json:"url"`
}

type DeleteJobListingResponse struct {
	DeleteJobID string `json:"deleteJobId"`
}

type JobListing struct {
	ID          string `json:"_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
	URL         string `json:"url"`
}

type Mutation struct {
}

type Query struct {
}

type UpdateJobListingInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}
