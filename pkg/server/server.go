package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const portNumber = ":8081"

var router *mux.Router

func Start() {

	router = mux.NewRouter()
	env := os.Getenv("ENV")

	fmt.Printf("Environment: %s\n", env)
	Create()
	// api.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
	// 	// * http://localhost:8081/api/example -> "From the API"
	// 	fmt.Fprintln(w, "From the API")
	// })

	if env == "prod" {
		// Serve static files if in production
		ServePlatform()
	} else {
		// Just serve the API if in development
		fmt.Printf("Serving api at: http://localhost%s/api\n", portNumber)
	}

	log.Fatal(http.ListenAndServe(portNumber, router))

	// * Alternative implementation with gin router:
	/*
		 router := gin.Default()

		// Serve frontend static files
		router.Use(static.Serve("/", static.LocalFile("./build", true)))

		// Setup route group for the API
		api := router.Group("/api")
		{
			api.GET("/", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "pong",
				})
			})
		}

		// Start and run the server
		fmt.Printf("Staring application on port %s", portNumber)
		router.Run(portNumber)
	*/

}
