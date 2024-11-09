package main

import "fmt"

func PrintASCII(asciiMap map[rune][]string, text string) {
	// Determine the maximum number of lines for any character
	var maxLines int
	for _, lines := range asciiMap {
		if len(lines) > maxLines {
			maxLines = len(lines)
		}
	}
	lines := SplitTextByLines(text)

	for _, line := range lines {
		for i := 0; i < maxLines; i++ {
			for _, char := range line {
				asciiLines, exists := asciiMap[char]
				if !exists {
					fmt.Printf("Error: Character '%c' not found in ASCII map\n", char)
					return
				}
				// Print the line if it exists, otherwise print spaces for padding
				if i < len(asciiLines) {
					fmt.Print(asciiLines[i])
				} else {
					fmt.Print("        ") // Adjust padding as needed
				}
			}
			fmt.Println()
		}
		fmt.Println() // Add a line break after each ASCII line block
	}
}
