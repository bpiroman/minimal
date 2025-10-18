package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// version will be set at build time with -ldflags "-X main.version=...".
// Default is "dev" for local builds.
var version = "dev"

func usage() {
	prog := filepath.Base(os.Args[0])
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options] arg1 arg2 ...\n\n", prog)
	fmt.Fprintln(flag.CommandLine.Output(), "Options:")
	flag.PrintDefaults()
}

func main() {
	// Define flags
	showVersion := flag.Bool("version", false, "Print version and exit")
	help := flag.Bool("help", false, "Show help")

	// Example flag
	repeat := flag.Int("n", 1, "how many times to print the args")

	flag.Usage = usage
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	if *showVersion {
		fmt.Println(version)
		return
	}

	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "error: at least one positional argument is required")
		flag.Usage()
		os.Exit(2)
	}

	// Simple behavior: print args 'n' times
	for i := 0; i < *repeat; i++ {
		fmt.Printf("Invocation %d: %v\n", i+1, args)
	}
}
