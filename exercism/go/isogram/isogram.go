package isogram

import "strings"

// IsIsogram checks if the input word or phrase is without a repeating letter.
// Spaces and hyphens are allowed to appear multiple times.
func IsIsogram(word string) bool {
    caps := strings.ToUpper(word)
	letters := make(map[rune]int, len(caps)) // preallocate the map size

    // put each letter in the map, filter out hyphens and spaces
    for _, char := range caps {
        if char == '-' || char == ' ' {
            char = -1
        }
        letters[char]++
	}

    // check for duplicates
    for _, char := range caps {
        if letters[char] > 1 {
            return false
        }
    }
	return true
}