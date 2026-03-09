package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jasparmetzger-dev/agora/internal/auth"
	"github.com/jasparmetzger-dev/agora/internal/database"
)

func main() {

	//init db
	db, err := database.NewPool(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var q *database.Queries = database.New(db)

	//init gin routing
	var r *gin.Engine = gin.Default()
	r.SetTrustedProxies(nil)

	r.POST("/register", auth.RegisterHandler(q))
	r.POST("/login", auth.LoginHandler(q))

	var authorized *gin.RouterGroup = r.Group("/")
	authorized.Use(auth.AuthMiddleware())
	{
		authorized.PATCH("/profile", UpdateProfileHandler(q))
		authorized.GET("/profile", GetProfileHandler(q))

		authorized.POST("/posts", CreatePostHandler(q))
		authorized.GET("/posts", GetAllPostsHandler(q))
		authorized.PATCH("/posts/:id", PatchPostHandler(q))
		authorized.DELETE("/posts/:id", DeletePostHandler(q))

		authorized.GET("/feed", FeedHandler(q))

	}

	r.Run(os.Getenv("BACKEND_PORT"))
}
