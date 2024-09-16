package utils

// UniqueRunes takes a string and converts it to a slice of unique runes.
// This method preserves the order of the first appearance of runes in the input string.
// Because Go's rune type handles Unicode characters properly, this solution works
// correctly even with strings containing non-ASCII characters.
func (u *Utils) UniqueRunes(s string) []rune {
	// Create a map to track unique runes
	seen := make(map[rune]bool)

	// Create a slice to store the unique runes
	var unique []rune

	// Iterate over the string, character by character
	for _, r := range s {
		// Check if the rune has already been encountered
		if !seen[r] {
			// Mark the rune as seen
			seen[r] = true
			// Append the rune to the unique slice
			unique = append(unique, r)
		}
	}

	return unique
}
