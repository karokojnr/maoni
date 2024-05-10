package main

import (
	"fmt"

	"github.com/karokojnr/maoni/internal/comment"
	"github.com/karokojnr/maoni/internal/db"
	transportHTTP "github.com/karokojnr/maoni/internal/transport/http"
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
	httpHandler := transportHTTP.NewHandler(commentService)

	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Maoni API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
