package datatypes

import (
	"strings"
)

func concatenateExample(s string) string {
	// Use concatenation to format the result to be "Classic <string>"
	return "Classic " + s
}

func concatenate(s string) string {
	// Use concatenation to format the result to be "Hello <string>!"
	return "Hello " + s + "!"
}

func substrings(word string) string {
	// Return the first 4 letters from the word using substring
	if len(word) >= 4 {
		return word[:4]
	}
	return word
}

func capitalize(word string) string {
	// Capitalize the first letter of the word
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(word[:1]) + word[1:]
}

func uppercase(s string) string {
	// Uppercase all letters in the string
	return strings.ToUpper(s)
}

func downcase(s string) string {
	// Downcase all letters in the string
	return strings.ToLower(s)
}

func emptyString(s string) bool {
	// Return true if the string is empty
	return len(s) == 0
}

func stringLength(s string) int {
	// Return the length of the string
	return len(s)
}

func reverse(s string) string {
	// Return the same string with all of its characters reversed
	var x string
	for _, v := range s {
		x = string(v) + x
	}
	return x
}

func spaceRemover(s string) string {
	// Remove all the spaces in the string using strings.ReplaceAll
	return strings.ReplaceAll(s, " ", "")
}
