package server

import (
	"fmt"
	"net/http"
)

// * Serves the react App
func ServePlatform() {
	// Serve static files if in production
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./build/static/"))))

	// Serve index page on all unhandled routes
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./build/index.html")
	})
	fmt.Printf("Serving at: http://localhost%s", portNumber)
}
