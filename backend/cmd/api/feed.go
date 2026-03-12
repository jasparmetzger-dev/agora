package api

import (
	"log"

	"github.com/gin-gonic/gin"
	db "github.com/jasparmetzger-dev/agora/cmd/database"
)

// GET, "/feed", auth required
func ShowFeedHandler(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Fatal()
	}
}
