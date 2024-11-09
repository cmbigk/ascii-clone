package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a valid text argument.")
		return
	}

	// Parse escape sequences (like "\n") in the provided text argument.
	// This function replaces any occurrences of `\n` with actual newline characters.
	text := ParseEscapeSequences(os.Args[1])

	// Load the ASCII art template from a file called "standard.txt".
	// This function reads the file, stores ASCII representations of characters in a map, and returns the map.
	asciiMap, err := LoadTemplate("standard.txt")
	if err != nil {
		fmt.Println("Error loading template:", err)
		return
	}
	PrintASCII(asciiMap, text)
}
