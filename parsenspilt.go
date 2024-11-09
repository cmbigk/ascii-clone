package main

import "strings"

// ParseEscapeSequences replaces "\n" in the input with actual newlines
func ParseEscapeSequences(input string) string {
	return strings.ReplaceAll(input, `\n`, "\n")
}

// SplitTextByLines splits the text based on "\n" sequences
func SplitTextByLines(text string) []string {
	return strings.Split(text, "\n")
}
