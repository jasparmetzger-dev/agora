package api

import (
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/jasparmetzger-dev/agora/cmd/database"
)

// POST, "/posts", requires auth
//requires data in form of an .mp4, title and description

func CreatePostHandler(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {

		//get ids
		userId, err := ValidateUUID(c)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		id := MakeId()

		//save video
		file, err := c.FormFile("video")
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		url, err := saveVideo(c, file, id.String())
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		//create post
		params := db.CreatePostParams{}
		params.ID = id
		params.Url = url
		params.Title = c.PostForm("title")
		params.Content = c.PostForm("description")
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
func saveVideo(c *gin.Context, file *multipart.FileHeader, id string) (pgtype.Text, error) { //returns url

	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		return pgtype.Text{}, nil
	}
	path := filepath.Join("uploads", id+filepath.Ext(file.Filename))
	if err := os.Rename(file.Filename, path); err != nil {
		return pgtype.Text{}, err
	}
	if err := c.SaveUploadedFile(file, path); err != nil {
		return pgtype.Text{}, err
	}

	return pgtype.Text{String: path, Valid: true}, nil
}
