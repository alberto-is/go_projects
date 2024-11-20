package servers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Note: Change to c.JSON(http.StatusOK, users) for better performance
func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, getAllUsers()) // c.JSON(http.StatusOK, users) IS better for performance, but c.IndentedJSON is better for debugging
}

func GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	c.IndentedJSON(http.StatusOK, findUserByID(id))
}

func PostUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}
	addUser(user)
	c.IndentedJSON(http.StatusCreated, user)
}

func PutUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
	}

	if err := updateUser(user); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusNoContent)
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	if err := deleteUserByID(id); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
