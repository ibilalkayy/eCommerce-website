package main

// Importing the libraries
import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ibilalkayy/eCommerce/routes"
)

// Executing the routes and setting up the port at :4000
func Execute() error {
	routes.Routes()
	fmt.Println("Starting the server at :4000")
	// Port to make it work on Heroku
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "4000"
	}
	return http.ListenAndServe(":"+port, nil)
	// Port to make it work locally
	// return http.ListenAndServe(":4000", nil)
}

// Executing the Execute() function
func main() {
	if err := Execute(); err != nil {
		log.Fatal(err)
	}
}
