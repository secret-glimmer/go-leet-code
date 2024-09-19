package main

import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	var result []string
	var currentLine []string
	currentLineLen := 0

	for _, word := range words {
		if currentLineLen+len(word)+len(currentLine) > maxWidth {
			result = append(result, justifyLine(currentLine, currentLineLen, maxWidth))
			currentLine = []string{word}
			currentLineLen = len(word)
		} else {
			currentLine = append(currentLine, word)
			currentLineLen += len(word)
		}
	}

	// Handle the last line
	lastLine := ""
	for _, word := range currentLine {
		lastLine += word + " "
	}
	lastLine = lastLine[:len(lastLine)-1]        // Remove the trailing space
	lastLine += spaces(maxWidth - len(lastLine)) // Add spaces to the end
	result = append(result, lastLine)

	return result
}

func justifyLine(line []string, lineLen, maxWidth int) string {
	if len(line) == 1 {
		return line[0] + spaces(maxWidth-len(line[0]))
	}

	totalSpaces := maxWidth - lineLen
	spaceBetweenWords := totalSpaces / (len(line) - 1)
	extraSpaces := totalSpaces % (len(line) - 1)

	result := ""
	for i, word := range line {
		result += word
		if i < len(line)-1 {
			result += spaces(spaceBetweenWords)
			if i < extraSpaces {
				result += " "
			}
		}
	}
	return result
}

func spaces(n int) string {
	return strings.Repeat(" ", n)
}

func main() {
	words1 := []string{"This", "is", "an", "example", "of", "text", "justification."}
	maxWidth1 := 16
	fmt.Println(fullJustify(words1, maxWidth1))

	words2 := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	maxWidth2 := 16
	fmt.Println(fullJustify(words2, maxWidth2))

	words3 := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	maxWidth3 := 20
	fmt.Println(fullJustify(words3, maxWidth3))
}
