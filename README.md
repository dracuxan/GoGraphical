[![wakatime](https://wakatime.com/badge/github/dracuxan/GoGraphical.svg)](https://wakatime.com/badge/github/dracuxan/GoGraphical)

# GoGraphical - Using GraphQL to create a CRUD API using Go and MongoDB

- [x] Schema File
- [x] Resolvers
- [x] Controllers
- [ ] Test Files
- [x] Workflows
- [ ] New features:
    - [ ] Authentication (JWT-based)
    - [ ] Pagination for job listings
    - [ ] Filtering and sorting

## Project Setup & Installation

1. Clone the repository:

```
git clone https://github.com/your-username/GoGraphical.git

cd GoGraphical
```

2. Install dependencies

```
go mod tidy
```

3. Run the server

```
go run main.go
```

## Directory Structure

```
GoGraphical/
├── controllers/                # Database interactions
│   └── controllers.go
├── go.mod
├── go.sum
├── gqlgen.yml
├── graph/
│   ├── generated.go
│   ├── model/                  # Structs for GraphQL types
│   │   └── models_gen.go
│   ├── resolver.go
│   ├── schema.graphqls         # GraphQL schema
│   └── schema.resolvers.go     # GraphQL schema resolvers
├── LICENSE
├── README.md
├── server.go                   # Entry point of the application
└── tools.go
```

## Running the GraphQL Playground

- Once the server is running, open: `http://localhost:8080/`

- Use the UI to test queries and mutations.


## Queries
### Query:
GettAllJobs:
```
query GetAllJob{
  jobs{
    _id,
    description,
    company,
    url
  }
}
```

---

GetJobById:
```
query GetJobById($id:ID!){
  job(id:$id){
    _id,
    description,
    company,
    url
  }
}
```
input:
```
{
  "id": "job-id"
}
```

---

### Mutations
CreateJobListing:
```
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
input:
```
{
  "input": {
    "title": "SDE - I",
    "description": "Work with us.",
    "company": "Google",
    "url": "www.google.com/"
  }
}
```

---

UpdateJobListing:
```
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
input:
```
{
  "id": "job-id",
  "input": {
    "title": "SDE - II"
  }
}
```

---

DeleteJobListing:
```
mutation DeleteJobListing($id: ID!) {
  deleteJobListing(id: $id) {
    deleteJobId
  }
}
```
input:
```
{
  "id" : "job id"
}
```

---

## Contributing

    1. Fork the repository
    2. Create a feature branch (`git checkout -b feature-branch`)
    3. Commit your changes (`git commit -m "Add new feature"`)
    4. Push to the branch (`git push origin feature-branch`)
    5. Open a Pull Request
