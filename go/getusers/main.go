package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var Users = []User{
	{Name: "Akbar Idria", Email: "asdasdas"},
	{Name: "Akbar dua", Email: ""},
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(Users); err != nil {
		http.Error(w, "Failed to stream struct", http.StatusInternalServerError)
		log.Printf("failed encoding users: %v", err)
		return
	}
}

func main() {
	http.HandleFunc("/", HelloWorldHandler)
	http.HandleFunc("/users", HandleGetUsers)
	fmt.Println("Server starting on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server at 8080")
	}
}
