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
		for i := 0; i < maxLines; i++ {
			for _, char := range line {

				asciiLines, exists := asciiMap[char]
				if !exists {
					fmt.Printf("Error: Character '%c' not found in ASCII map\n", char)
					return
				}

				if i < len(asciiLines) {
					fmt.Print(asciiLines[i])
				} else {
					fmt.Print("        ")
				}
			}
			if !(lineIndex == len(lines)-1 && i == maxLines-1) {
				fmt.Println()
			}
		}
		if lineIndex < len(lines)-1 {
			fmt.Println()
		}
	}
}
