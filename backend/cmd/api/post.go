package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/jasparmetzger-dev/agora/cmd/database"
)

// GET, "posts/:id", auth required
func GetPostMetaDataHandler(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := ValidatePostUUID(c)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		post, err := q.GetPostById(c, id)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"post": post})
	}
}

// GET, "posts/:id/video", auth required

func GetVideoHandler(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		path, err := getPostUrl(q, c)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.File(path)
	}
}

// PATCH, "posts/:id", auth required
func PatchPostHandler(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		//validate input
		var req struct {
			Title   string `json:"title"`
			Content string `json:"content"`
			// BE ABLE TO CHANGE VIDEO TOO??
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		//get post
		id, err := ValidatePostUUID(c)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}
		post, err := q.GetPostById(c, id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		//prepare post update
		if req.Content != "" {
			post.Content = req.Content
		}
		if req.Title != "" {
			post.Title = req.Title
		}
		params := db.UpdatePostByIdParams{}
		params.ID = id
		params.Title = post.Title
		params.Content = post.Content

		//update
		new_post, err := q.UpdatePostById(c, params)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "update successful", "post": new_post})

	}
}

// DELETE, "posts/:id", auth required
func DeletePostHandler(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := ValidatePostUUID(c)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}
		post, err := q.DeletePostById(c, id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "successful deletion", "post": post})
	}
}
