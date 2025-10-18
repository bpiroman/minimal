package main

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
)

//go:embed templates
var templates embed.FS

// createDir creates a directory if it doesn't exist
func createDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0o755); err != nil {
			return fmt.Errorf("failed to create %s: %w", path, err)
		}
		fmt.Printf("📁 Created directory: %s\n", path)
	} else {
		fmt.Printf("⚠️  Directory %s already exists, skipping.\n", path)
	}
	return nil
}

func main() {
	// get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting current working director:", err)
		return
	}

	fmt.Println("🚀 Creating minimal Go web template in:", cwd)

	// Create static directory folders
	staticDir := filepath.Join(cwd, "static")
	cssDir := filepath.Join(staticDir, "css")
	faviconDir := filepath.Join(staticDir, "favicon")

	createDir(staticDir)
	createDir(cssDir)
	createDir(faviconDir)

	// read and copy index.html
	indexHTML, _ := templates.ReadFile("templates/index.html")
	err = writeFile(filepath.Join(cwd, "index.html"), indexHTML)
	if err != nil {
		fmt.Println("error writing file:", err)
		return
	}
	// read and copy main.go
	mainGO, _ := templates.ReadFile("templates/main.go")
	err = writeFile(filepath.Join(cwd, "main.go"), mainGO)
	if err != nil {
		fmt.Println("error writing file:", err)
		return
	}
	// read and copy styles.css
	stylesCSS, _ := templates.ReadFile("templates/styles.css")
	err = writeFile(filepath.Join(cwd, "static/css/styles.css"), stylesCSS)
	if err != nil {
		fmt.Println("error writing file:", err)
		return
	}
}

func writeFile(filePath string, data []byte) error {
	// Check if the file exists.
	if _, err := os.Stat(filePath); err == nil {
		// File exists, so we return nil to indicate success (a successful skip).
		return nil
	} else if !os.IsNotExist(err) {
		// os.Stat failed for a real reason (e.g., permission error), so return that error.
		return err
	}
	err := os.WriteFile(filePath, data, 0o644)
	return err
}
