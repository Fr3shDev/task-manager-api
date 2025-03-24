package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var users []User
var nextUserID = 1

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Only accepts post requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the incoming request to a user struct
	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Basic validation: username and password must be provided
	if newUser.Username == "" || newUser.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	// Simulate saving the user in memory
	newUser.ID = nextUserID
	nextUserID++
	users = append(users, newUser)

	log.Printf("User registered: %+v\n", newUser)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

// LoginHandler handles user login requests.
// It expects a POST request with a JSON body containing username and password.

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Only accepts post requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the incoming request to a user struct
	var loginData User
	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// For demo purposes, iterate over our in-memory users to find a match.
	for _, user := range users {
		if user.Username == loginData.Username && user.Password == loginData.Password {
			json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
			// json.NewEncoder(w).Encode(loginData)
			return
		}
	}

	http.Error(w, "Invalid username or password", http.StatusUnauthorized)
}
