package main

import (
	"fmt"

	"github.com/Kylescottw/pulse-api/internal/db"
	"github.com/Kylescottw/pulse-api/internal/service/comment"
	transportHttp "github.com/Kylescottw/pulse-api/internal/transport/http"
	"github.com/Kylescottw/pulse-api/internal/util/system"
)




func Run() error {

	if err := system.PreFlight(); err != nil {
		return err
	}

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
	
	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err 
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		// TODO: emit err to rollbar, or data dog... error monitoring.
		fmt.Println(err)
	}
}