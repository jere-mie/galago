package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	build()
	serve()
}

func serve() {
	// Use http.FileServer to serve files from the "./public" directory
	fileServer := http.FileServer(http.Dir("./public"))

	// Handle all requests with the file server
	http.Handle("/", fileServer)

	// Start the server on port 1112
	fmt.Println("Running webserver on http://localhost:1112")
	err := http.ListenAndServe(":1112", nil)
	if err != nil {
		panic(err)
	}
}

func build() {
	// clear ./public directory before building site
	if err := os.RemoveAll("./public"); err != nil {
		fmt.Println("Error clearing ./public directory:", err)
		return
	} else {
		fmt.Println("Successfully cleared ./public directory")
	}

	// recursively render all templates in ./pages directory
	if err := filepath.Walk("pages/", process_template); err != nil {
		fmt.Printf("Error walking the path './pages': %v\n", err)
		return
	}

	// copying ./static directory into ./public/static/
	if err := copyDir("./static", "./public/static"); err != nil {
		fmt.Println("Error copying static directory:", err)
	} else {
		fmt.Println("Successfully copied ./static directory")
	}
}
