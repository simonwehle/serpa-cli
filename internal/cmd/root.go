package cmd

import (
	"flag"
	"fmt"
	"os"

	"serpa-cli/internal/files"
	"serpa-cli/internal/utils"
)

const toolName = "serpa-cli"
const version = "0.0.1"

func Execute() {
	baseUrl := flag.String("u", "", "Serpa Maps base url")
	showHelp := flag.Bool("h", false, "Show this help message")
	showVersion := flag.Bool("v", false, "Show version")

	categoriesFile := "categories.csv"
	placesFile := "places.csv"

	if len(os.Args) == 1 {
		utils.PrintVersion(toolName, version)
		fmt.Fprintln(os.Stdout)
		utils.PrintHelp(toolName)
		os.Exit(0)
	}

	flag.Parse()

	if *showHelp {
		utils.PrintHelp(toolName)
		return
	}

	if *showVersion {
		utils.PrintVersion(toolName, version)
		return
	}

	if *baseUrl == "" {
		fmt.Fprintln(os.Stderr, "Error: Serpa Maps server base url must be set. Set it using -u (base-url)")
		os.Exit(1)
	}

	var (
		imagesExists = false
	)

	fileExistsOrExit(categoriesFile)
	fileExistsOrExit(placesFile)
	fmt.Printf("Files categories.csv and places.csv exist scanning for images ...")

	root := "."
	folders, images, err := files.CountFoldersAndImages(root)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if images != 0 {
		fmt.Printf("%d images in %d folders found\n", images, folders)
	} else {
		imagesExists = true
		fmt.Print("No images found; skipping upload assets step")
	}

	csvCategories, err := files.ReadCategoriesCSV(root, categoriesFile)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("CSV categories: %s", csvCategories)

	if imagesExists {

	}
}

func fileExistsOrExit(file string) {
	if _, err := os.Stat(file); err != nil {
		fmt.Fprintln(os.Stderr, "Error: No", file, "found")
		os.Exit(1)
	} 
}
