package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/jasparmetzger-dev/agora/cmd/database"
)

// POST, "/posts", requires auth
func CreatePostHandler(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		//validate and get data
		var req struct {
			Video   []byte `json:"video"` //whatever the actual video is
			Title   string `json:"title"`
			Content string `json:"content"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		//create url and save the image
		url, err := saveVideo(req.Video)

		userId, err := ValidateUUID(c)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		//create post
		params := db.CreatePostParams{}
		params.Url = url
		params.Title = req.Title
		params.Content = req.Content
		params.UserID = userId

		created_post, err := q.CreatePost(c, params)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "post created successfully", "post": created_post})
	}
}

// GET, "/posts", requires auth
func GetAllPostsHandler(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := ValidateUUID(c)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		posts, err := q.GetPostsByUserId(c, userId)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "fetching posts successful", "posts": posts})
	}
}

// the saving logic
func saveVideo(video []byte) (pgtype.Text, error) { //returns url
	log.Fatal("not implemented")
	return pgtype.Text{}, nil
}
