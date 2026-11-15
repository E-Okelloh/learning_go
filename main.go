package main

// Import the packages we need.
import (
	"fmt"
	"net/http"
)

// This function handles requests sent to our homepage.
func homeHandler(w http.ResponseWriter, r *http.Request) {

	// Write a response back to whoever made the request.
	fmt.Fprintln(w, "Welcome to my first Go API!")
}

// The main function is where our application starts.
func main() {

	// Register a handler for the root URL.
	http.HandleFunc("/", homeHandler)

	// Start the HTTP server on port 8080.
	fmt.Println("Server running at http://localhost:8080")

	// ListenAndServe keeps the server running.
	// It stops if the server encounters an error.
	err := http.ListenAndServe(":8080", nil)

	// Check whether the server encountered an error.
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
