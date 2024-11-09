package main

import "strings"

// ParseEscapeSequences replaces the escape sequence "\n" in the input string
// with actual newline characters ("\n"). This allows us to handle text that
// uses the string representation of newline characters (e.g., "\n") and convert
// them into real newlines that can be processed correctly by the program.
func ParseEscapeSequences(input string) string {
	return strings.ReplaceAll(input, `\n`, "\n")
}

// SplitTextByLines splits the input text into a slice of strings based on the
// newline character ("\n"). Each element in the resulting slice represents
// one line of text, allowing us to process each line individually.
func SplitTextByLines(text string) []string {
	return strings.Split(text, "\n")
}
