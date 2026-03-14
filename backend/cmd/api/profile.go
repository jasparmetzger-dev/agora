package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jasparmetzger-dev/agora/cmd/auth"
	db "github.com/jasparmetzger-dev/agora/cmd/database"
)

// GET, "/profile"
func GetProfileHandler(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err, status := MakeUserFromHeader(q, c)
		if err != nil {
			c.JSON(status, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"user": user})
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
		user, err, status := MakeUserFromHeader(q, c)
		if err != nil {
			c.JSON(status, gin.H{"error": err.Error()})
			return
		}

		//update user
		if req.Email != "" {
			user.Email = req.Email
		}
		if req.Username != "" {
			user.Username = req.Username
		}

		new_user, err := UserUpdateHelper(q, c, user)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		//return
		c.JSON(200, gin.H{"message": "successfully updated user", "user": new_user})
	}
}

// PATCH, "/profile/changepassword"
// PASSWORD LENGTH CHECK??
func ChangePasswordHandler(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		//validate req
		var req struct {
			OldPassword string `json:"old_password"`
			NewPassword string `json:"new_password"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		//get user
		user, err, status := MakeUserFromHeader(q, c)
		if err != nil {
			c.JSON(status, gin.H{"error": err.Error()})
			return
		}
		//validate passwords
		if !auth.CheckPasswordHash(req.OldPassword, user.PasswordHash) {
			c.JSON(401, gin.H{"error": "invalid old_password, must match the set password"})
			return
		}
		//check pwd-requirements here
		//...
		//set new password
		user.PasswordHash, err = auth.HashPassword(req.NewPassword)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error(), "message": "new password could not be assigned"})
			return
		}
		new_user, err := UserUpdateHelper(q, c, user)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error(), "message": "user could not be updated"})
			return
		}
		c.JSON(200, gin.H{"user": new_user, "message": "updated password successfully"})
	}
}
