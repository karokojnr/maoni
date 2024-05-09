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

	store, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to database")
	}
	err = store.MigrateDB()
	if err != nil {
		log.Error("failed to setup database")
		return err
	}

	commentService := comment.NewService(store)
	fmt.Println(commentService.GetComment(context.Background(), ""))

	return nil
}

func main() {
	fmt.Println("Maoni API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
