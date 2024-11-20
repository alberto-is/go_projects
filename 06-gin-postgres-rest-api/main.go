package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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

// Note: Change to c.JSON(http.StatusOK, users) for better performance
func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users) // c.JSON(http.StatusOK, users) IS better for performance, but c.IndentedJSON is better for debugging
}

func main() {

	//go.SetMode(gin.ReleaseMode) // Uncomment this line to set the mode to release
	router := gin.Default()

	router.GET("/users", getUsers)

	router.Run(":8080")

}
