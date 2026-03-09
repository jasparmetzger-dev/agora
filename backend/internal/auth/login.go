package auth

import (
	"context"
	"errors"

	"os"

	"github.com/gin-gonic/gin"
	db "github.com/jasparmetzger-dev/agora/internal/database"
)

func LoginHandler(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		//handle input
		var req struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required,min=8"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		//log in
		token, err := login(req.Username, req.Password, c.Request.Context(), q)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error(), "message": "registering didnt work"})
			return
		}
		c.JSON(200, gin.H{"token": token})
	}
}

func login(username, password string, ctx context.Context, q *db.Queries) (string, error) {
	//verify user
	user, err := q.GetUserByUsername(ctx, username)
	if err != nil {
		return "", errors.New("Username not found.")
	}
	if !CheckPasswordHash(password, user.PasswordHash) {
		return "", errors.New("Password is incorrect.")
	}

	//generate token
	token, err := GenerateJWT(user.ID.String(), os.Getenv("SECRET_KEY"))
	if err != nil {
		return "", err
	}
	return token, nil
}
