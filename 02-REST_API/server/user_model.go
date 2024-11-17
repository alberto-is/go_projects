package server

import "fmt"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User = []User{
	{ID: 1, Name: "Alberto", Age: 20},
	{ID: 2, Name: "Irene", Age: 19},
}

func GetAllUsers() []User {
	return users
}

func FindUserById(id string) (*User, error) {
	for _, user := range users {
		if fmt.Sprint(user.ID) == id {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("User %s Not found", id)
}
