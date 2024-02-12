package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read the story from a file
	story, err := readFile("story.txt")
	if err != nil {
		fmt.Println("Error reading story file:", err)
		return
	}

	// Define placeholder markers
	targetStart := "<"
	targetEnd := ">"

	// Find and collect unique placeholders
	words := collectPlaceholders(story, targetStart, targetEnd)

	// Create a map to store user-provided words
	answers := make(map[string]string)

	// Prompt for words and fill the map
	for word := range words {
		fmt.Printf("Enter a word for %q: ", word)
		answer, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		answers[word] = strings.TrimSpace(answer) // Remove trailing newline
	}

	// Replace placeholders with user-provided words
	updatedStory := replacePlaceholders(story, answers, targetStart, targetEnd)

	// Print the final story
	fmt.Println(updatedStory)
}

func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var content string
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}
	return content, scanner.Err()
}

func collectPlaceholders(story string, start, end string) map[string]bool {
	words := make(map[string]bool)
	startIndex := -1
	for i, char := range story {
		if char == start[0] {
			startIndex = i
		} else if char == end[0] && startIndex != -1 {
			word := story[startIndex+1 : i] // Extract word without markers
			words[word] = true
			startIndex = -1
		}
	}
	return words
}

func replacePlaceholders(story string, answers map[string]string, start, end string) string {
	for word, answer := range answers {
		story = strings.ReplaceAll(story, start+word+end, answer)
	}
	return story
}
