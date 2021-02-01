package server

import "os"

// PORT is the port the server is hosted from
var PORT = "3000"

func readENV() {
	envPort := os.Getenv("PORT")
	if envPort != "" {
		PORT = envPort
	}
}
