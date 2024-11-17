package main

import (
	"fmt"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hola Mundo")
}

func main() {
	var err error

	http.HandleFunc("/", greet)
	server := http.Server{
		Addr: ":8080",
	}

	err = server.ListenAndServe()

	if err != nil {
		panic(err)
	}

}
