package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jasparmetzger-dev/agora/cmd/api"
	"github.com/jasparmetzger-dev/agora/cmd/auth"
	"github.com/jasparmetzger-dev/agora/cmd/database"
	"github.com/jasparmetzger-dev/agora/conf"
)

func main() {

	//init db
	db, err := database.NewPool(conf.DATABASE_URL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var q *database.Queries = database.New(db)

	//init gin routing
	var r *gin.Engine = gin.Default()
	r.SetTrustedProxies(nil)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	//auth
	r.POST("/register", auth.RegisterHandler(q))
	r.POST("/login", auth.LoginHandler(q))

	//user routing
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

	r.Run(conf.BACKEND_PORT)
}
