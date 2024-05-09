package main

import (
	"context"
	"fmt"

	"github.com/karokojnr/maoni/internal/comment"
	"github.com/karokojnr/maoni/internal/db"
	log "github.com/sirupsen/logrus"
)

// Run - responsible for the instantiation and startup of our go application
func Run() error {
	fmt.Println("Starting Maoni...")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to database")
	}
	if err := db.MigrateDB(); err != nil {
		log.Error("failed to migrate database")
		return err
	}

	commentService := comment.NewService(db)
	commentService.PostComment(context.Background(), comment.Comment{
		ID:   "5c137bb8-d468-45ad-8795-63d252be9d25",
		Slug: "test",
		Body: "test",
	})

	fmt.Println(commentService.GetComment(context.Background(), "5c137bb8-d468-45ad-8795-63d252be9d25"))
	return nil
}

func main() {
	fmt.Println("Maoni API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
