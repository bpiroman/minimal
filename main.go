package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

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

func CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer sourceFile.Close() // Ensure the source file is closed

	destinationFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer destinationFile.Close() // Ensure the destination file is closed

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return fmt.Errorf("failed to copy file content: %w", err)
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
	err = CopyFile("templates/index.html", filepath.Join(cwd, "index.html"))
	if err != nil {
		log.Fatalf("Error copying file: %v", err)
	}
	// read and copy main.go - local server
	err = CopyFile("templates/main.go", filepath.Join(cwd, "main.go"))
	if err != nil {
		log.Fatalf("Error copying file: %v", err)
	}
	// read and copy styles.css
	err = CopyFile("templates/styles.css", filepath.Join(cwd, "static/css/styles.css"))
	if err != nil {
		log.Fatalf("Error copying file: %v", err)
	}
}
