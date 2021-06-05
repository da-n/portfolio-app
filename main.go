package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if err := envCheck(); err != nil {
		log.Fatalf("One or more environment variables are not defined, terminating application: %v", err)
	}
}

// envCheck verify all necessary environment variables are set to run application
func envCheck() error {
	envVars := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
	}

	for _, k := range envVars {
		if os.Getenv(k) == "" {
			return fmt.Errorf("Environment variable %s not defined. Terminating application...", k)
		}
	}

	return nil
}

