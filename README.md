[![wakatime](https://wakatime.com/badge/github/dracuxan/GoGraphical.svg)](https://wakatime.com/badge/github/dracuxan/GoGraphical)

# GoGraphical - Using GraphQL to create a CRUD API using Go and MongoDB

- [x] Schema File
- [x] Resolvers
- [x] Controllers
- [ ] Test Files
- [x] Workflows

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
