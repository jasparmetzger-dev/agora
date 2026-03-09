package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jasparmetzger-dev/agora/internal/auth"
	"github.com/jasparmetzger-dev/agora/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	//init db
	db, err := database.NewPool(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var q *database.Queries = database.New(db)

	//init gin routing
	var r *gin.Engine = gin.Default()
	r.POST("/register", auth.RegisterHandler(q))
	r.POST("/login", auth.LoginHandler(q))

	var authorized *gin.RouterGroup = r.Group("/")
	authorized.Use()
	{

	}

	r.Run(os.Getenv("BACKEND_PORT"))
}
