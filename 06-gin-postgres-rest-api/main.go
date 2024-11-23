package main

import (
	servers "06-gin-postgres-rest-api/server"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/users", servers.GetUsers)
	router.GET("/users/:id", servers.GetUser)
	router.POST("/users", servers.PostUser)
	router.PUT("/users", servers.PutUser)
	router.DELETE("/users/:id", servers.DeleteUser)
	router.Run(":8080")

}
