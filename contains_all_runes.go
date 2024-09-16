package utils

// ContainsAllRunes checks if all runes in a `subset` slice are contained in `set` slice.
// This method correctly handles cases where runes are duplicated, both in the set and
// subset slices.
// This works with Unicode characters, as Go's rune type properly supports Unicode.
func (u *Utils) ContainsAllRunes(set, subset []rune) bool {
	// Create a map to count occurrences of each rune in `set`
	runeCount := make(map[rune]int)

	// Count occurrences of each rune in `set`
	for _, r := range set {
		runeCount[r]++
	}

	// Check if all runes in `subset` are present in `runeCount` with sufficient count
	for _, r := range subset {
		if runeCount[r] == 0 {
			// If a rune is not present or count is insufficient, return false
			return false
		}
		runeCount[r]--
	}

	// All runes in `subset` are contained in `set`
	return true
}
