package etl

import "strings"

// Transform transforms the legacy data format to the shiny new format
func Transform(in map[int][]string) map[string]int {
	out := make(map[string]int)

    for points, letters := range in {
        for _, letter := range letters {
            transformedLetter := strings.ToLower(letter)
            out[transformedLetter] = points
        }
    }
	return out
}
