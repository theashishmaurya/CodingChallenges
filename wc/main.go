package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"unicode"
)

// Define command-line flags for counting options
var (
	c bool // Option for counting bytes
	l bool // Option for counting lines
	w bool // Option for counting words
	m bool // Option for counting characters
)

// Initialize command-line flags
func init() {
	flag.BoolVar(&c, "c", false, "Outputs the number of bytes in a file")
	flag.BoolVar(&l, "l", false, "Outputs the number of lines in a file")
	flag.BoolVar(&w, "w", false, "Outputs the number of words in a file")
	flag.BoolVar(&m, "m", false, "Outputs the number of character in a file")
}

// Read the content of a file and return it as a string
func readFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// Count the number of bytes, lines, words, and characters in the provided string
func getCount(fileStr string) (int, int, int, int) {
	numberOfByte := len(fileStr)
	var numberOfLine int
	var numberOfWords int
	var numberOfChar int

	inWord := false

	for i := 0; i < len(fileStr); i++ {
		// Loop over fileStr

		// Count the number of lines
		if fileStr[i] == '\n' {
			numberOfLine++
		}

		// Check for word boundaries
		if isWordBoundary(rune(fileStr[i])) {
			// If the character is a space, punctuation, or newline, we are not in a word
			inWord = false
		} else {
			// If the character is not a space, punctuation, or newline and we were not in a word,
			// increment the word count
			if !inWord {
				numberOfWords++
			}
			inWord = true
		}

	}

	// Count the number of characters excluding newline characters
	runes := []rune(fileStr)
	numberOfChar = len(runes)

	return int(numberOfByte), numberOfLine, numberOfWords, numberOfChar
}

// Check if a character is a word boundary (space)
func isWordBoundary(char rune) bool {
	return unicode.IsSpace(char)
}

func main() {
	// Parse command-line flags
	flag.Parse()

	var input string
	args := flag.Args()
	var fileStr string
	var filePath string
	// Check if a file path is provided as a command-line argument
	if len(args) != 0 {

		filePath = args[0]
		var err error
		fileStr, err = readFile(filePath)
		if err != nil {
			fmt.Println("Error: Something went wrong", err)
			return
		}
	} else {
		// If no file path is provided, read from stdin
		stat, _ := os.Stdin.Stat()

		if (stat.Mode() & os.ModeCharDevice) == 0 {
			// Data is being piped in through stdin
			reader := bufio.NewReader(os.Stdin)
			for {
				line, err := reader.ReadString('\n')
				input += line
				// Check for the end of input
				if err != nil {
					break
				}
			}
			fileStr = input
		} else {
			// No file path and no stdin input provided
			fmt.Println("Error: No input provided. Please provide either a file path or input through stdin.")
			return
		}
	}

	// Get counts based on the specified options
	numberOfByte, numberOfLine, numberOfWords, numberOfChar := getCount(fileStr)

	// If no counting options are specified, print all counts and the file path
	if c == false && l == false && w == false && m == false {
		fmt.Println(numberOfByte, numberOfLine, numberOfWords, filePath)
		return
	}

	// Print counts based on the specified options
	if c {
		fmt.Println(numberOfByte, "", filePath)
	}

	if l {
		fmt.Println(numberOfLine, "", filePath)
	}

	if w {
		fmt.Println(numberOfWords, "", filePath)
	}

	if m {
		fmt.Println(numberOfChar, "", filePath)
	}

	return
}
