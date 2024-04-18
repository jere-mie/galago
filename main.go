package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		greet()
		build()
		return
	}

	executable := os.Args[0]
	command := os.Args[1]

	switch command {
	case "server", "serve":
		greet()
		if err := build(); err != nil {
			return
		}
		if len(os.Args) > 2 && os.Args[2] == "dev" {
			log.Println("Starting server in dev mode, project will rebuild every 2 seconds...")
			go serveAndAutoRebuild()
		} else {
			log.Println("Starting server...")
		}
		serve()
	case "build":
		greet()
		build()
	case "version":
		fmt.Println(getVersion())
	case "new":
		if len(os.Args) < 3 {
			log.Printf("Not enough arguments for `%s new`\n", executable)
			log.Printf("Usage: %s new [sitename]\n", executable)
			os.Exit(1)
		}
		greet()
		siteName := os.Args[2]
		log.Printf("Creating new site \"%s\"\n", siteName)
		if err := newSite(siteName); err != nil {
			log.Printf("Error creating new site `%s`, %s\n", siteName, err)
			return
		}
		log.Printf("Successfully created new site \"%s\"!\n\n", siteName)
	default:
		log.Println("Unknown command:", command)
		log.Printf("Usage: %s [serve|build|version|new <sitename>]\n", executable)
		os.Exit(1)
	}
}

func serve() {
	fileServer := http.FileServer(http.Dir("./public"))
	http.Handle("/", fileServer)
	log.Println("Running webserver on http://localhost:1112")
	err := http.ListenAndServe(":1112", nil)
	if err != nil {
		panic(err)
	}
}

func build() error {
	log.Println("Building project...")
	if err := os.RemoveAll("./public"); err != nil {
		log.Println("Error clearing ./public directory:", err)
		return err
	}
	log.Println("Successfully cleared ./public directory")

	if err := filepath.Walk("pages/", processTemplate); err != nil {
		log.Printf("Error walking the path './pages': %v\n", err)
		return err
	}

	if err := copyDir("./static", "./public/static"); err != nil {
		log.Println("Error copying static directory:", err)
	} else {
		log.Println("Successfully copied ./static directory")
	}
	return nil
}

func serveAndAutoRebuild() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("Rebuilding project...")
		if err := build(); err != nil {
			log.Println("Error rebuilding project:", err)
		} else {
			log.Println("Project rebuilt successfully.")
		}
	}
}
