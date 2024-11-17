package main

import (
	"02-REST_API/server"
)

func main() {
	var err error
	server := server.New(":8080")

	err = server.ListenAndServe()

	if err != nil {
		panic(err)
	}

}
