package cmd

import (
	"flag"
	"fmt"
	"os"

	"serpa-cli/internal/utils"
)

const toolName = "serpa-cli"
const version = "0.0.1"

func Execute() {
	baseUrl := flag.String("u", "", "Serpa Maps base url")
	showHelp := flag.Bool("h", false, "Show this help message")
	showVersion := flag.Bool("v", false, "Show version")

	if len(os.Args) == 1 {
		utils.PrintHelp(toolName)
		os.Exit(0)
	}

	flag.Parse()

	if *showHelp {
		utils.PrintHelp(toolName)
		return
	}

	if *showVersion {
		fmt.Printf("%s version %s\n", toolName, version)
		return
	}

	// Serpa Maps server base url must be set. Set it using -u (base-url)

	if *baseUrl == "" {
		if _, err := os.Stat("categories.csv"); err != nil {
			fmt.Fprintln(os.Stderr, "Error: No categories.csv found")
			os.Exit(1)
		}
		if _, err := os.Stat("places.csv"); err != nil {
			fmt.Fprintln(os.Stderr, "Error: No places.csv found")
			os.Exit(1)
		}
	}
}
