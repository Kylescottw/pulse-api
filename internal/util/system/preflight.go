package system

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func parseEnvExample(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var missingEnvKeys []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
		envKeyValue := scanner.Text() // envKey=envValue
		if strings.Contains(envKeyValue, "=") { // confirm current scanner.Text() is format in format. I.E envKey=envValue
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
	missingEnvKeys, err := parseEnvExample("../../.env-example") // reference .env-example for required env variables
  if err != nil {
    log.Fatalf("ERROR: Parsing Env Example: %s", err)
		return err
  }
	
	if len(missingEnvKeys) > 0 {
		errStr := strings.Join(missingEnvKeys, ", ")
		return fmt.Errorf("ERROR: Env Variables required: %s", errStr)
	}

	return nil
}