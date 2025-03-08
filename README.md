[![wakatime](https://wakatime.com/badge/github/dracuxan/GoGraphical.svg)](https://wakatime.com/badge/github/dracuxan/GoGraphical)

# GoGraphical

**GoGraphical** is a CRUD API built using GraphQL, Go, and MongoDB. This project showcases a framework for managing job listings using GraphQL.

## Features

- [x] Schema File
- [x] Resolvers
- [x] Controllers
- [ ] Test Files
- [x] Workflows
- [ ] Planned Features:
    - [ ] Authentication (JWT-based)
    - [ ] Pagination for job listings
    - [ ] Filtering and sorting

## Project Setup & Installation

### Prerequisites

- Go (version 1.16 or later)
- MongoDB

### Installation Steps

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-username/GoGraphical.git
   cd GoGraphical
   ```

2. **Install dependencies:**

   ```bash
   go mod tidy
   ```

3. **Run the server:**

   ```bash
   go run server.go
   ```

## Directory Structure

```
GoGraphical/
├── controllers/                # Database interactions
│   └── controllers.go
├── go.mod
├── go.sum
├── gqlgen.yml
├── graph/
│   ├── generated.go
│   ├── model/                  # Structs for GraphQL types
│   │   └── models_gen.go
│   ├── resolver.go
│   ├── schema.graphqls         # GraphQL schema
│   └── schema.resolvers.go     # GraphQL schema resolvers
├── LICENSE
├── README.md
├── server.go                   # Entry point of the application
└── tools.go
```

## Running the GraphQL Playground

- Once the server is running, open: [http://localhost:8080/](http://localhost:8080/)
- Use the UI to test queries and mutations.

## GraphQL Queries

### Query: Get All Jobs

```graphql
query GetAllJob {
  jobs {
    _id
    description
    company
    url
  }
}
```

### Query: Get Job By ID

```graphql
query GetJobById($id: ID!) {
  job(id: $id) {
    _id
    description
    company
    url
  }
}
```

**Input:**

```json
{
  "id": "job-id"
}
```

## GraphQL Mutations

### Mutation: Create Job Listing

```graphql
mutation CreateJobListing($input: CreateJobListingInput!) {
  createJobListing(input: $input) {
    _id
    title
    description
    company
    url
  }
}
```

**Input:**

```json
{
  "input": {
    "title": "SDE - I",
    "description": "Work with us.",
    "company": "Google",
    "url": "www.google.com/"
  }
}
```

### Mutation: Update Job Listing

```graphql
mutation UpdateJob($id: ID!, $input: UpdateJobListingInput!) {
  updateJobListing(id: $id, input: $input) {
    title
    description
    _id
    company
    url
  }
}
```

**Input:**

```json
{
  "id": "job-id",
  "input": {
    "title": "SDE - II"
  }
}
```

### Mutation: Delete Job Listing

```graphql
mutation DeleteJobListing($id: ID!) {
  deleteJobListing(id: $id) {
    deleteJobId
  }
}
```

**Input:**

```json
{
  "id": "job-id"
}
```
