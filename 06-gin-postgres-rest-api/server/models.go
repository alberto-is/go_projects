package servers

import "errors"

var ErrorUserNotFound error = errors.New("user not found")

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User = []User{
	{ID: 1, Name: "Josefa", Age: 20},
	{ID: 2, Name: "Ramon", Age: 42},
	{ID: 3, Name: "Laura", Age: 25},
}

func getAllUsers() []User {
	return users
}

func findUserByID(id int) User {
	for _, user := range users {
		if user.ID == id {
			return user
		}
	}
	return User{}
}

func addUser(user User) {
	users = append(users, user)
}

func updateUser(user User) error {
	for i, u := range users {
		if u.ID == user.ID {
			users[i] = user
			return nil
		}
	}
	return ErrorUserNotFound

}

func deleteUserByID(id int) error {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return ErrorUserNotFound
}
