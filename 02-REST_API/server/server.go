package server

import "net/http"

func New(addr string) *http.Server {
	setupRoutes()
	return &http.Server{
		Addr: addr,
	}
}
