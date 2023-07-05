package main

import (
	"encoding/json"
	"log"
	"net/http"
)

//struct represents a sample resource
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users []User

//func handler for GET /users endpoint
func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

//func handler for POST /users endpoint
func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	//assigns new ID & adds user to list
	newUser.ID = len(users) + 1
	users = append(users, newUser)

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(newUser)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}


func main() {
	//initialize sample data
	users = []User{
		{ID: 1, Name: "John Doe"},
		{ID: 2, Name: "Jane Smith"},
	}

	//set up API endpoints
	http.HandleFunc("/users", getUsers).Methods("GET")
	http.HandleFunc("/users", createUser).Methods("POST")

	//server start
	log.Fatal(http.ListenAndServe(":8080", nil))
}
