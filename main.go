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
	text := ParseEscapeSequences(os.Args[1])

	asciiMap, err := LoadTemplate("standard.txt")
	if err != nil {
		fmt.Println("Error loading template:", err)
		return
	}

	PrintASCII(asciiMap, text)
}
