package main

import (
	"fmt"
	"strings"
	"testing"
)

var fmtPrint = fmt.Print

// captureOutput captures the output of PrintASCII
func captureOutput(asciiMap map[rune][]string, text string) string {
	var builder strings.Builder
	// Override fmt.Print to write to builder
	old := fmtPrint
	defer func() { fmtPrint = old }()
	fmtPrint = func(args ...interface{}) (n int, err error) {
		return fmt.Fprintf(&builder, "%v", args...)
	}

	// Handle escape sequences in the input text
	text = strings.ReplaceAll(text, "\\n", "\n")

	// Call the PrintASCII function and capture output
	PrintASCII(asciiMap, text)

	// Normalize line endings to \n for consistency
	return strings.ReplaceAll(builder.String(), "\r\n", "\n")
}

// TestASCIIArt tests the PrintASCII function
func TestASCIIArt(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Empty Input",
			input: "",
			want:  "",
		},
		{
			name:  "Single Escape Sequence '\\n'",
			input: "\\n",
			want:  "\n",
		},
		{
			name:  "Simple Hello",
			input: "Hello",
			want: ` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                 `,
		},
		{
			name:  "Hello with newline",
			input: "Hello\\nThere",
			want: ` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                 
 _______   _                            
|__   __| | |                           
   | |    | |__     ___   _ __    ___  
   | |    |  _ \   / _ \ | '__|  / _ \ 
   | |    | | | | |  __/ | |    |  __/ 
   |_|    |_| |_|  \___| |_|     \___/  
                                        `,
		},
		{
			name:  "Quoted Input",
			input: "{Hello There}",
			want: `   __  _    _          _   _                 _______   _                           __    
  / / | |  | |        | | | |               |__   __| | |                          \ \   
 | |  | |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___   | |  
/ /   |  __  |  / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \   \ \ 
\ \   | |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/   / / 
 | |  |_|  |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___|  | |  
  \_\                                                                              /_/   
                                                                                          `,
		},
	}

	// Load the ASCII map from a test template (you'll need to implement LoadTemplate)
	asciiMap, err := LoadTemplate("standard.txt")
	if err != nil {
		t.Fatalf("Error loading template: %v", err)
	}

	// Iterate through each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Capture the output from PrintASCII
			output := captureOutput(asciiMap, tc.input)

			// Compare the result with the expected output
			if output != tc.want {
				t.Errorf("For input %s,\nExpected output:\n`%s`\nBut got:\n`%s`", tc.input, tc.want, output)
			}
		})
	}
}
