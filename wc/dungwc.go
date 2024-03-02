package main

import (
	"flag"
	"fmt"
	"os"
	"unicode"
)

func main() {
	wordCountFlag := flag.Bool("w", false, "Count words")
	lineCountFlag := flag.Bool("l", false, "Count lines")
	charCountFlag := flag.Bool("m", false, "Count characters")
	byteCountFlag := flag.Bool("c", false, "Count bytes")

	flag.Parse()

	if !*wordCountFlag && !*lineCountFlag && !*charCountFlag && !*byteCountFlag {
		fmt.Println("Usage: dungwc [-w] [-l] [-m] [-c] <filename>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: dungwc [-w] [-l] [-m] [-c] <filename>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	filename := flag.Args()[0]

	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(1)
	}

	lines := countLines(string(content))
	words := countWords(string(content))
	characters := countCharacters(string(content))
	bytes := countBytes(content)

	if *byteCountFlag {
		fmt.Printf("Bytes: %d\n", bytes)
	}

	if *lineCountFlag {
		fmt.Printf("Lines: %d\n", lines)
	}
	if *wordCountFlag {
		fmt.Printf("Words: %d\n", words)
	}
	if *charCountFlag {
		fmt.Printf("Characters: %d\n", characters)
	}
}

func countLines(text string) int {
	count := 0
	for _, char := range text {
		if char == '\n' {
			count++
		}
	}
	return count
}

func countWords(text string) int {
	words := 0
	inWord := false

	for _, char := range text {
		if unicode.IsSpace(char) || char == '\n' {
			inWord = false
		} else if !inWord {
			inWord = true
			words++
		}
	}

	return words
}

func countCharacters(text string) int {
	return len([]rune(text))
}

func countBytes(content []byte) int {
	return len(content)
}
