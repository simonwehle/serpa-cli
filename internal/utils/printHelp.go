package utils

import (
	"fmt"
)

func PrintHelp(toolName string) {
	fmt.Printf("Usage:\n")
	fmt.Printf("  %s -u <serpa-maps server url>\n\n", toolName)

	fmt.Println("Example:")
	fmt.Printf("  %s -u http://localhost:53164\n", toolName)
}

func PrintVersion(toolName, version string) {
	fmt.Printf("%s version %s\n", toolName, version)
}
