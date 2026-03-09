package main

import (
	"context"
	"log"
	"os"

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
	//init database query
	var q *database.Queries = database.New(db)
	var ctx context.Context = context.Background()

}
