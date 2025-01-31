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
    "github_url": "https://github.com/sejoroajose/HNG-INTERNSHIP/Task_0"
}
```

## Installation & Running

### Prerequisites
- Install [Go](https://go.dev/dl/)

### Steps
1. Clone the repository:
   ```sh
   git clone https://github.com/sejoroajose/HNG-INTERNSHIP.git
   ```
2. Navigate to the project directory:
   ```sh
   cd HNG-INTERNSHIP/Task_0
   ```
3. Install dependencies:
   ```sh
   go mod tidy
   ```
4. Run the server:
   ```sh
   go run main.go
   ```
5. The API will be available at:
   ```sh
   http://localhost:8080/api/user
   ```

## Want to hire a Backend Developer?
[https://hng.tech/hire/golang-developers](Hire backend Golang dev)

