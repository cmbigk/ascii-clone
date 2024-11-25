package main

import "fmt"

func PrintASCII(asciiMap map[rune][]string, text string) {
	var maxLines int
	for _, lines := range asciiMap {
		if len(lines) > maxLines {
			maxLines = len(lines)
		}
	}

	lines := SplitTextByLines(text)

	for lineIndex, line := range lines {
		if line == "" {
			// Only print a single blank line for empty input lines
			fmt.Println()
			continue
		}
		for i := 0; i < maxLines; i++ {
			for _, char := range line {
				asciiLines, exists := asciiMap[char]
				if !exists {
					fmt.Printf("Error: Character '%c' not found in ASCII map\n", char)
					return
				}

				// Print the ASCII line for the character if available, otherwise space
				if i < len(asciiLines) {
					fmt.Print(asciiLines[i])
				} else {
					fmt.Print("        ")
				}
			}
			fmt.Println() // Newline after each row of ASCII characters
		}
		// Only add a blank line if the next line is not empty
		if lineIndex < len(lines)-1 && lines[lineIndex+1] != "" {
			fmt.Println()
		}

	}
}
