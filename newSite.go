//go:generate go generate
package main

import (
	_ "embed"
	"log"
	"os"
	"path/filepath"
)

//go:embed embed/base.html
var baseHTML string

//go:embed embed/index.html
var indexHTML string

//go:embed embed/about.html
var aboutHTML string

//go:embed embed/style.css
var styleCSS string

//go:embed embed/gitignore
var gitIgnore string

//go:embed embed/README.md
var readme string

func newSite(siteName string) error {
	// Define a map for directory paths and corresponding file contents
	directories := map[string]string{
		siteName + "/pages/index.html":    indexHTML,
		siteName + "/pages/about.html":    aboutHTML,
		siteName + "/templates/base.html": baseHTML,
		siteName + "/static/style.css":    styleCSS,
		siteName + "/.gitignore":          gitIgnore,
		siteName + "/README.md":           readme,
	}

	// Iterate over the map and create directories and write files
	for dir, content := range directories {
		if err := os.MkdirAll(filepath.Dir(dir), 0755); err != nil {
			return err
		}
		if err := os.WriteFile(dir, []byte(content), 0755); err != nil {
			log.Printf("Error writing to file %q: %v\n", dir, err)
		}
	}

	return nil
}
