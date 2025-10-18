package main

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
)

// create var with embed.FS type
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

	// read templates
	// mainGo, _ := templates.ReadFile("templates/main.go.txt")
	indexHTMLContent, err := templates.ReadFile("templates/index.html")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(string(indexHTMLContent))

	// if err := writeFileIfNotExists(filepath.Join(cwd, "main.go"), mainGo); err != nil {
	// 	fmt.Println("Error:", err)
	// }
	if err := writeFileIfNotExists(filepath.Join(cwd, "index.html"), indexHTMLContent); err != nil {
		fmt.Println("Error:", err)
	}
}
