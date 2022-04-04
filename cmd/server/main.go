package main

import (
	"context"
	"fmt"

	"github.com/Kylescottw/pulse-api/internal/comment"
	"github.com/Kylescottw/pulse-api/internal/db"
)

// Run - is going to be responsible for
// the instantiation and start up of the
// go application.
func Run() error {
	fmt.Println("starting up our application")

	
	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	} 
	fmt.Println("successfully connected and pinged the database")

	cmtService := comment.NewService(db)
	
	cmtService.PostComment(
		context.Background(),
		comment.Comment{
			ID: "49443e52-aebb-439e-adf9-0d7159872810",
			Slug: "manual-test",
			Author: "Elliot",
			Body: "Hello World",
		},
	)


	fmt.Println(cmtService.GetComment(
		context.Background(), 
		"49443e52-aebb-439e-adf9-0d7159872810",
	))

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		// TODO:: emit err to rollbar, or data dog... error monitoring.
		fmt.Println(err)
	}
}