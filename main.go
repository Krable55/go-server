package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/krable55/go-server/pkg/db"
	"github.com/krable55/go-server/pkg/server"
)

// * Build environemnts
const (
	production  = "prod"
	development = "dev"
)

func init() {
	env := flag.String("env", development, "dev | prod")
	flag.Parse()
	os.Setenv("ENV", *env)
}

func main() {

	env := os.Getenv("ENV")

	fmt.Printf("Environment: %s\n", env)

	// Connect DB
	db.Connect()

	// Start server
	server.Start()
}
