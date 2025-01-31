# HNG INTERNSHIP TASK 0

# User Info API

This is a simple REST API built with Go using the Gorilla Mux router. The API provides information about a user, including their email, current datetime, and GitHub URL.

## Features
- Returns user information in JSON format
- Supports CORS for cross-origin requests
- Lightweight and easy to deploy

## Technologies Used
- Go
- Gorilla Mux
- Gorilla Handlers (for CORS support)
- net/http package

## API Endpoint

### Get User Info
- **Endpoint:** `/api/user`
- **Method:** `GET`
- **Response:** JSON containing user details

#### Example Response:
```json
{
    "email": "ajosesejoro@gmail.com",
    "current_datetime": "2024-01-31T15:04:05Z",
    "github_url": "https://github.com/sejoroajose/HNG-INTERNSHIP"
}
```

## Want to know more about this?
https://hng.tech/hire/golang-developers