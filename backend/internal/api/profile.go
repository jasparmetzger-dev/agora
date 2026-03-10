package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/jasparmetzger-dev/agora/internal/database"
)

// GET, "/profile"
func GetProfileHandler(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := ValidateUUID(c)
		if err != nil {
			c.JSON(401, gin.H{"error": err})
		}
		user, err := q.GetUserById(c, id)
		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}

// PATCH, "/profile"
func UpdateProfileHandler(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		//validate req
		var req struct {
			Username string `json:"username"`
			Email    string `json:"email"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		//get user
		id, err := ValidateUUID(c)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
		}
		user, err := q.GetUserById(c, id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}

		//update user
		if req.Email != "" {
			user.Email = req.Email
		}
		if req.Username != "" {
			user.Username = req.Username
		}
		q.UpdateUserById(c, user)
	}
}

// PATCH, "/profile/changepassword"
func ChangePasswordHandler(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
