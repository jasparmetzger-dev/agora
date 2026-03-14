package auth

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	db "github.com/jasparmetzger-dev/agora/cmd/database"
	"github.com/jasparmetzger-dev/agora/conf"
)

func RegisterHandler(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Username string `json:"username" binding:"required"`
			Email    string `json:"email" binding:"required,email"`
			Password string `json:"password" binding:"required,min=8"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		token, err := register(req.Username, req.Email, req.Password, c.Request.Context(), q)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error(), "message": "registering didnt work"})
			return
		}
		c.JSON(200, gin.H{"token": token})
	}
}
func register(username, email, password string, ctx context.Context, q *db.Queries) (string, error) {
	// Check if the username already exists
	_, err := q.GetUserByUsername(ctx, username)
	if err == nil {
		return "", errors.New("username already exists")
	}

	//make user
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return "", err
	}
	params := db.CreateUserParams{
		Username:     username,
		Email:        email,
		PasswordHash: hashedPassword,
	}
	_, err = q.CreateUser(ctx, params)
	if err != nil {
		return "", err
	}

	//get user id
	user, err := q.GetUserByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	//generate token
	token, err := GenerateJWT(user.ID.String(), conf.SECRET_KEY)
	if err != nil {
		return "", err
	}
	return token, nil
}
