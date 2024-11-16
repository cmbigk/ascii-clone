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


	chars := 0
   for _, line := range lines {
       chars += len(line)
   }
   if chars == 0 {
       lines = lines[:len(lines)-1]
   }


   for _, line := range lines {


       if line == "" {
           fmt.Println()
       } else {


           for i := 0; i < 8; i++ {
               for _, char := range line {
                   asciiLines, exists := asciiMap[char]
                   if !exists {
                       fmt.Printf("Error: Character '%c' not found in ASCII map\n", char)
                       return
                   }
                   fmt.Print(asciiLines[i])
               }
               fmt.Println()
           }
       }
   }
}

