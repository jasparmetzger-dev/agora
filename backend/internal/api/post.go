package api

import (
	"log"

	"github.com/gin-gonic/gin"
	db "github.com/jasparmetzger-dev/agora/internal/database"
)

// PATCH, "posts/:id", auth required
func PatchPostHandler(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Fatal()
	}
}

// DELETE, "posts/:id", auth required
func DeletePostHandler(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Fatal()
	}
}
