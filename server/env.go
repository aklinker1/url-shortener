package server

import (
	"fmt"
	"os"
)

// PORT is the port the server is hosted from
var PORT = "3000"
var DATABASE_URL = ""
var IS_PROD = false

func readENV() {
	envPort := os.Getenv("PORT")
	if envPort != "" {
		PORT = envPort
	}

	envDatabaseURL := os.Getenv("DATABASE_URL")
	if envDatabaseURL == "" {
		fmt.Println("DATABASE_URL is a required env var")
		os.Exit(1)
	} else {
		DATABASE_URL = envDatabaseURL
	}

	envIsProd := os.Getenv("IS_PROD")
	if envIsProd != "" {
		IS_PROD = envIsProd == "true"
	}
}
