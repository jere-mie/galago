//go:generate go generate
package main

import (
	_ "embed"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/flosch/pongo2/v6"
)

//go:embed embed/greeting.txt
var greeting string

//go:embed version.txt
var version string

func copyDir(src, dst string) error {
	// Create destination directory if it does not exist
	if err := os.MkdirAll(dst, 0755); err != nil {
		return err
	}

	// Get list of files and directories in the source directory
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	// Loop through the entries
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		// If the entry is a directory, recursively call copyDir
		if entry.IsDir() {
			if err := copyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			// If the entry is a file, copy it
			if err := copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}
	return nil
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

func processTemplate(path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Printf("Error accessing path %q: %v\n", path, err)
		return nil
	}
	if !info.IsDir() {
		// Read file contents
		content, err := os.ReadFile(path)
		if err != nil {
			log.Printf("Error reading file %q: %v\n", path, err)
			return nil
		}

		// Parse the template
		template, err := pongo2.FromString(string(content))
		if err != nil {
			log.Println("Error parsing template:", err)
			return nil
		}

		// Render the template
		outputString, err := template.Execute(nil)
		if err != nil {
			log.Println("Error rendering template:", err)
			return nil
		}

		relativePath, err := filepath.Rel("pages", path)
		if err != nil {
			log.Printf("Error getting relative path for %q: %v\n", path, err)
			return nil
		}
		destPath := filepath.Join("public", relativePath)

		// Create the destination directory if it doesn't exist
		destDir := filepath.Dir(destPath)
		if err := os.MkdirAll(destDir, 0755); err != nil {
			log.Printf("Error creating directory %q: %v\n", destDir, err)
			return nil
		}

		if err := os.WriteFile(destPath, []byte(outputString), 0755); err != nil {
			log.Printf("Error writing to file %q: %v\n", destPath, err)
		}

		log.Printf("Page exported: %s\n", destPath)

	}
	return nil
}

func greet() {
	log.Printf("\n%s\n", greeting)
}

func getVersion() string {
	return version
}
