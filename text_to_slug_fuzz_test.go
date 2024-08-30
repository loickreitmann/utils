package utils_test

import (
	"regexp"
	"strings"
	"testing"

	"github.com/loickreitmann/utils"
)

// FuzzTextToSlug is a fuzz test for the TextToSlug function
func FuzzUtils_TextToSlug(f *testing.F) {
	// Add seed corpus
	f.Add("Café con leche y croissants délicieux à l'étage")
	f.Add("Привет мир!")
	f.Add("こんにちは世界")
	f.Add("Hello, World! @ # $")
	f.Add("")

	// Define the fuzzing function
	f.Fuzz(func(t *testing.T, input string) {
		var u utils.Utils
		slug := u.TextToSlug(input)

		// Basic property checks:
		// Check if the output does not contain spaces
		if strings.Contains(slug, " ") {
			t.Errorf("Slug contains spaces: %q", slug)
		}

		// Check if the output is lowercase
		if slug != strings.ToLower(slug) {
			t.Errorf("Slug is not in lowercase: %q", slug)
		}

		// Check if the output does not contain consecutive dashes
		if strings.Contains(slug, "--") {
			t.Errorf("Slug contains consecutive dashes: %q", slug)
		}

		// Check if the output does not contain any non-permitted characters
		if matched, _ := regexp.MatchString(`[^a-z0-9-]`, slug); matched {
			t.Errorf("Slug contains non-permitted characters: %q", slug)
		}
	})
}
