package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
    "github.com/gorilla/handlers"
)

type User struct {
    Email               string  `json:"email"`
    Current_datetime    string  `json:"current_datetime"`
    Github_url          string  `json:"github_url"`
}


func getUserInfo(w http.ResponseWriter, r *http.Request){
    user := User{
        Email:  "ajosesejoro@gmail.com",
        Current_datetime: time.Now().Format(time.RFC3339),
        Github_url: "https://github.com/sejoroajose/HNG-INTERNSHIP/", 
    }

    w.Header().Set("content-Type", "application/json")

    json.NewEncoder(w).Encode(user)
}

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/api/user", getUserInfo).Methods("GET")

    corsMiddleware := handlers.CORS(
        handlers.AllowedOrigins([]string{"*"}),                
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
    )

    handler := corsMiddleware(router)

    log.Println("Server starting on port 8080...")
    if err := http.ListenAndServe(":8080", handler); err !=nil{
        log.Fatal(err)
    }
}

