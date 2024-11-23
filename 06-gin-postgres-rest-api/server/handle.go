package servers

import (
	"06-gin-postgres-rest-api/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// NOTE: Test only
func InitDB() {
	db.InitTable()
}

// Note: Change to c.JSON(http.StatusOK, users) for better performance
func GetUsers(c *gin.Context) {
	users, err := db.GetUsers()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, users) // c.JSON(http.StatusOK, users) IS better for performance, but c.IndentedJSON is better for debugging
}

func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	user, errr := db.GetUserByID(id)
	if errr != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": errr.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func PostUser(c *gin.Context) {
	var user db.User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}
	userID, err := db.InsertUser(user.Name, user.Age)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.ID = userID
	c.IndentedJSON(http.StatusCreated, user)
}

// func PutUser(c *gin.Context) {
// 	var user User
// 	if err := c.BindJSON(&user); err != nil {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
// 	}

// 	if err := updateUser(user); err != nil {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 	}
// 	c.Status(http.StatusNoContent)
// }

// func DeleteUser(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
// 		return
// 	}
// 	if err := deleteUserByID(id); err != nil {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.Status(http.StatusNoContent)
// }
