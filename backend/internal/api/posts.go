package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/jasparmetzger-dev/agora/internal/database"
)

// POST, "/posts", requires auth
func CreatePostHandler(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {

		//validate and get data
		var req struct {
			Url     pgtype.Text `json:"url"`
			Title   string      `json:"title"`
			Content string      `json:"content"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		userId, err := ValidateUUID(c)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}

		//create post
		params := db.CreatePostParams{}
		params.Url = req.Url
		params.Title = req.Title
		params.Content = req.Content
		params.UserID = userId

		created_post, err := q.CreatePost(c, params)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		c.JSON(200, gin.H{"message": "post created successfully", "post": created_post})

	}
}

// GET, "/posts", requires auth
func GetAllPostsHandler(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Fatal()
	}
}

// just logic
func postPost(post db.Post) {
	log.Fatal("not implemented")
}
