package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jasparmetzger-dev/agora/internal/api"
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
		authorized.GET("/profile", api.GetProfileHandler(q))
		authorized.PATCH("/profile", api.UpdateProfileHandler(q))
		authorized.PATCH("/profile/changepassword", api.ChangePasswordHandler(q))

		authorized.POST("/posts", api.CreatePostHandler(q))
		authorized.GET("/posts", api.GetAllPostsHandler(q))

		authorized.PATCH("/posts/:id", api.PatchPostHandler(q))
		authorized.DELETE("/posts/:id", api.DeletePostHandler(q))

		authorized.GET("/feed", api.ShowFeedHandler(q))

	}

	r.Run(os.Getenv("BACKEND_PORT"))
}
