package main

import "fmt"

// PrintASCII takes an ASCII map and a text input, and prints the corresponding
// ASCII art representation of each character in the text based on the template.
func PrintASCII(asciiMap map[rune][]string, text string) {

	var maxLines int //use maxLines to find the maximum number of lines any character's ASCII representation occupies
	// to ensure that we print the ASCII art correctly, with all lines aligned.
	for _, lines := range asciiMap {
		if len(lines) > maxLines {
			maxLines = len(lines) // Track the character with the most lines in its ASCII art
		}
	}

	// Split the input text into separate lines based on the newline character
	lines := SplitTextByLines(text)

	for _, line := range lines {
		for i := 0; i < maxLines; i++ {
			for _, char := range line {
				// Check if the character exists in the ASCII map
				asciiLines, exists := asciiMap[char]
				if !exists {
					fmt.Printf("Error: Character '%c' not found in ASCII map\n", char)
					return
				}

				// If the current line index is within the character's ASCII lines, print the corresponding line
				// Otherwise, print spaces for padding to keep the formatting aligned
				if i < len(asciiLines) {
					fmt.Print(asciiLines[i])
				} else {
					// Print blank spaces if the character's ASCII representation is shorter than the longest one
					fmt.Print("        ")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
