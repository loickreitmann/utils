package utils_test

import (
	"math"
	"testing"

	"github.com/loickreitmann/utils"
)

func TestUtils_RandomString(t *testing.T) {
	// arrange
	var testUtils utils.Utils
	expectedLens := []int{-23, 0, 1, 5, 7, 22, 36, 789, 10000}
	actualLens := []int{}

	// act
	for i := range expectedLens {
		str := testUtils.RandomString(expectedLens[i])
		actualLens = append(actualLens, len(str))
	}

	// assert
	for i := range actualLens {
		if actualLens[i] != int(math.Abs(float64(expectedLens[i]))) {
			t.Errorf("wrong string length; expected %d, got %d", expectedLens[i], actualLens[i])
		}
	}
}

func TestUtils_RandomString_CustomeSeed(t *testing.T) {
	// arrange
	var testUtils utils.Utils
	expectedLens := []int{-23, 0, 1, 5, 7, 314}
	actualStrs := []string{}
	customSeed := []rune("ou812")

	// act
	for i := range expectedLens {
		str := testUtils.RandomString(expectedLens[i], customSeed)
		actualStrs = append(actualStrs, str)
	}

	// assert
	for i := range actualStrs {
		if len(actualStrs[i]) != int(math.Abs(float64(expectedLens[i]))) {
			t.Errorf("wrong string length; expected %d, got %d", expectedLens[i], len(actualStrs[i]))
		}
		if !containsAllRunes(customSeed, uniqs) {
		uniqs := testUtils.UniqueRunes(actualStrs[i])
			t.Error("characters found in string not contained i seed set")
		}
	}
}

// Function to check if all runes in `subset` are contained in `set`
func containsAllRunes(set, subset []rune) bool {
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
