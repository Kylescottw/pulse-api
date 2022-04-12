package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Kylescottw/pulse-api/internal/comment"
	"github.com/Kylescottw/pulse-api/internal/db"
	transportHttp "github.com/Kylescottw/pulse-api/internal/transport/http"
)

func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var missingEnvKeys []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
		envKeyValue := scanner.Text() // envKey=envValue
		if strings.Contains(envKeyValue, "=") { // confirm current scanner.Text() is format in Env var format. I.E envKey=envValue
			envKey := strings.Split(envKeyValue, "=")[0] // grab envKey
			_, present := os.LookupEnv(envKey) // check env exists
			if !present {
				missingEnvKeys = append(missingEnvKeys, envKey) // append envKey to missingEnvKeys for error reporting
			}
		}
  }
  return missingEnvKeys, scanner.Err()
}

func PreFlight() (error) {
	missingEnvKeys, err := readLines("../../.env-example") // reference .env-example for required env variables
  if err != nil {
    log.Fatalf("readLines: %s", err)
		return err
  }
	
	if len(missingEnvKeys) > 0 {
		errStr := strings.Join(missingEnvKeys, ", ")
		return fmt.Errorf("ERROR: Env Variables required: %s", errStr)
	}

	return nil
}


func Run() error {

	if err := PreFlight(); err != nil {
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