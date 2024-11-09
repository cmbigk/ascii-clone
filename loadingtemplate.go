package main

import (
	"bufio"
	"os"
)

// LoadTemplate loads the ASCII template into a map
func LoadTemplate(filePath string) (map[rune][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	asciiMap := make(map[rune][]string)
	scanner := bufio.NewScanner(file)

	var character rune = ' ' // Start from space (' ') character
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// Save current character's ASCII representation and increment character
			if len(lines) > 0 {
				asciiMap[character] = lines
				character++
				lines = []string{}
			}
		} else {
			lines = append(lines, line)
		}
	}
	if len(lines) > 0 {
		asciiMap[character] = lines
	}
	return asciiMap, nil
}
