package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/jasparmetzger-dev/agora/cmd/database"
)

// GET, "/feed", auth required
func ShowFeedHandler(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(500, gin.H{"error": "not implemented."})
	}
}
