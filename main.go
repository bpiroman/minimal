package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

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

func writeFileIfNotExists(name string, content []byte) error {
	if _, err := os.Stat(name); err == nil {
		fmt.Printf("⚠️  %s already exists, skipping.\n", filepath.Base(name))
		return nil
	}
	if err := os.WriteFile(name, content, 0o644); err != nil {
		return fmt.Errorf("error writing %s: %w", name, err)
	}
	fmt.Printf("✅ Created %s\n", filepath.Base(name))
	return nil
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting current working director:", err)
		return
	}
	log.Println(cwd)

	fmt.Println("🚀 Creating minimal Go web template in:", cwd)

	// Create static directory folders
	staticDir := filepath.Join(cwd, "static")
	cssDir := filepath.Join(staticDir, "css")
	faviconDir := filepath.Join(staticDir, "favicon")

	createDir(staticDir)
	createDir(cssDir)
	createDir(faviconDir)

	mainGo, _ := templates.ReadFile("templates/main.go.txt")
	indexHTML, _ := templates.ReadFile("templates/index.html")

	fmt.Println("🚀 Creating minimal Go web template in:", cwd)
	if err := writeFileIfNotExists(filepath.Join(cwd, "main.go"), mainGo); err != nil {
		fmt.Println("Error:", err)
	}
	if err := writeFileIfNotExists(filepath.Join(cwd, "index.html"), indexHTML); err != nil {
		fmt.Println("Error:", err)
	}
}
