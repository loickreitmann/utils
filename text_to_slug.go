package utils

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// The removeAccents function uses the Unicode Normalization Form D (NFD) to decompose
// characters into their base form and combining marks. It then removes diacritics by
// filtering out the characters categorized under `unicode.Mn` (mark, nonspacing).
func (u *Utils) removeAccents(s string) string {
	// Normalize the string to NFD form to separate base characters from accents
	t := norm.NFD.String(s)
	var bldr strings.Builder
	for _, r := range t {
		// Skip combining marks (accents), thus removing diacritics
		if unicode.Is(unicode.Mn, r) {
			continue
		}
		bldr.WriteRune(r)
	}
	return bldr.String()
}

// The TextToSlug function converts accented characters to their unaccented versions,
// replaces all non-alphanumeric characters with dashes, trims redundant dashes,
// and converts the string to lowercase.
// This approach makes the slug both URL-friendly and human-readable.
func (u *Utils) TextToSlug(input string) string {
	// Step 1: Remove accents
	normalized := u.removeAccents(input)

	// Step 2: Replace spaces and non-URL-friendly characters with dashes
	// Allow only letters, numbers, and replace other characters with dashes
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	slug := reg.ReplaceAllString(normalized, "-")

	// Step 3: Trim and lowercase the slug
	slug = strings.Trim(slug, "-")
	slug = strings.ToLower(slug)

	return slug
}
