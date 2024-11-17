package server

import "net/http"

func setupRoutes() error {
	http.HandleFunc("/", greet)
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/users/", getUserById)
	return nil
}
