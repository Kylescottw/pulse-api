package main

import (
	"context"
	"fmt"

	"github.com/Kylescottw/pulse-api/internal/db"
)

// Run - is going to be responsible for
// the instantiation and start up of the
// go application.
func Run() error {
	fmt.Println("starting up our application")


	fmt.Println("hello world")
	
	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		return err
	}

	fmt.Println("successfully connected and pinged the database")

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		// TODO:: emit err to rollbar, or data dog... error monitoring.
		fmt.Println(err)
	}
}