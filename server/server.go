package server

import (
	"fmt"
	"net/http"
	"io/fs"
)

// Start the application
func Start(ui *fs.FS, metaJSON string) {
	fmt.Println("Reading env...")
	readENV()

	fmt.Println("Creating router...")
	r := createRouter(ui, metaJSON)

	fmt.Println("Connecting to postgres...")
	connectDB()

	fmt.Printf("Starting server @ :%s\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), r)
}
