package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

func main() {
	user := User{
		ID:   1,
		Name: "Akbar Idria",
	}

	JSONData, err := json.MarshalIndent(user, "", " ")
	if err != nil {
		fmt.Println("failed to marshal struct")
	}

	fmt.Printf("Marshaled structed: %s \n\n", JSONData)

	JSONString := `[{"id": 1, "name": "Alice Smith"}]`

	var UserTwo []User
	err = json.Unmarshal([]byte(JSONString), &UserTwo)
	if err != nil {
		fmt.Println("Failed to unmarshall")
	}

	for _, item := range UserTwo {
		fmt.Printf("User item %#v \n", item)
	}
}
