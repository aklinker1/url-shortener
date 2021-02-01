package server

import (
	"fmt"
	"net/http"
)

// Start the application
func Start() {
	fmt.Println("Reading env...")
	readENV()

	fmt.Println("Creating router...")
	r := createRouter()

	fmt.Println("Connecting to postgres...")

	fmt.Printf("Starting server @ :%s\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), r)
}
