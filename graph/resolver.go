package graph

import (
	"github.com/dracuxan/job-listing-api/controllers"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *controllers.DB
}
