package main

import (
	"02-REST_API/server"
)

func main() {
	var err error

	srv := server.New(":8080")

	err = srv.ListenAndServe()

	if err != nil {
		panic(err)
	}

}
