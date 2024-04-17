package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// show startup text
	greet()

	if len(os.Args) < 2 {
		build()
		return
	}

	// Extract the executable
	executable := os.Args[0]

	// Extract the command
	command := os.Args[1]

	// Handle different commands
	switch command {
	case "server", "serve":
		log.Println("Building project...")
		err := build()
		if err != nil {
			return
		}
		log.Println("Starting server...")
		go serve()
	case "build":
		log.Println("Building project...")
		build()
	case "new":
		if len(os.Args) < 3 {
			log.Printf("Not enough arguments for `%s new`\n", executable)
			log.Printf("Usage: %s new [sitename]\n", executable)
			os.Exit(1)
		}
		siteName := os.Args[2]
		log.Printf("Creating new site \"%s\"\n", siteName)
		if err := newSite(siteName); err != nil {
			log.Printf("Error creating new site `%s`, %s\n", siteName, err)
			return
		}
		log.Printf("Successfully created new site \"%s\"!\n\n", siteName)
	default:
		log.Println("Unknown command:", command)
		log.Printf("Usage: %s [serve|build]\n", executable)
		os.Exit(1)
	}
}

func serve() {
	// Use http.FileServer to serve files from the "./public" directory
	fileServer := http.FileServer(http.Dir("./public"))

	// Handle all requests with the file server
	http.Handle("/", fileServer)

	// Start the server on port 1112
	log.Println("Running webserver on http://localhost:1112")
	err := http.ListenAndServe(":1112", nil)
	if err != nil {
		panic(err)
	}
}

func build() error {
	log.Println("Building project...")

	// clear ./public directory before building site
	if err := os.RemoveAll("./public"); err != nil {
		log.Println("Error clearing ./public directory:", err)
		return err
	} else {
		log.Println("Successfully cleared ./public directory")
	}

	// recursively render all templates in ./pages directory
	if err := filepath.Walk("pages/", process_template); err != nil {
		log.Printf("Error walking the path './pages': %v\n", err)
		return err
	}

	// copying ./static directory into ./public/static/
	if err := copyDir("./static", "./public/static"); err != nil {
		log.Println("Error copying static directory:", err)
	} else {
		log.Println("Successfully copied ./static directory")
	}
	return nil
}
