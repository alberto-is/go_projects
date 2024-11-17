package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hola Mundo")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")
	json.NewEncoder(w).Encode(GetAllUsers())

}

func getUserById(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")
	id := r.URL.Path[len("/users/"):]
	user, err := FindUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(*user)
}
