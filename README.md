## ASCII-Art

### Overview

The ASCII-Art project is a Go program that takes a string input and generates a graphical representation of the text using ASCII characters. The program reads pre-designed "banner files" containing ASCII representations of characters and outputs the string in the specified format. It supports various inputs, including letters, numbers, special characters, spaces, and newlines.


### Objectives
- Learn and implement the Go file system (fs) for file handling.
- Manipulate data to generate ASCII-based graphical text.
- Apply good coding practices


### Features
- Supports input strings with:
    * Letters (uppercase and lowercase)
    * Numbers
    * Special characters (e.g., !, {, })
    * Spaces and newlines (\n)
- Utilizes ASCII templates from preformatted "banner files".
- Generates text art with consistent height (8 lines per character).
- Handles edge cases like empty strings and unsupported characters.
- Outputs the result in a format that can be processed by tools like cat.

### Input Example

**Input** : "Hello"
```
 _    _          _   _
| |  | |        | | | | 
| |__| |   ___  | | | |   ___
|  __  |  / _ \ | | | |  / _ \
| |  | | |  __/ | | | | | (_) |
|_|  |_|  \___| |_| |_|  \___/

````

### How It Works
1. **Read ASCII Banners:**
    - The program loads "banner files" (e.g., shadow, standard, thinkertoy) that define ASCII art for each printable character.
    - Each character representation is 8 lines tall and separated by a newline.
2. **Parse Input:**
    - The input string is split into lines.
    - Special characters, spaces, and newlines are handled appropriately.
3. **Generate ASCII Art:**
    - For each character in the input, the corresponding ASCII representation is fetched from the banner.
    - Lines are aligned and printed together for proper formatting.
4. **Output:**
    - The generated ASCII art is printed line by line to the command line in the specified format.


### Usage

Run the Program:
````
go run . "Input" "banner"
````
Pipe the output to a file or another program:
````
go run . "Input" "banner" | cat -e
````
