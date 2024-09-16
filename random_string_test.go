package utils_test

import (
	"math"
	"testing"

	"github.com/loickreitmann/utils"
)

func TestUtils_RandomString(t *testing.T) {
	// ARRANGE
	var testUtils utils.Utils
	expectedLens := []int{-23, 0, 1, 5, 7, 22, 36, 789, 10000}
	actualLens := []int{}

	// ACT
	for i := range expectedLens {
		str := testUtils.RandomString(expectedLens[i])
		actualLens = append(actualLens, len(str))
	}

	// ASSERT
	for i := range actualLens {
		if actualLens[i] != int(math.Abs(float64(expectedLens[i]))) {
			t.Errorf("wrong string length; expected %d, got %d", expectedLens[i], actualLens[i])
		}
	}
}

func TestUtils_RandomString_CustomeSeed(t *testing.T) {
	// ARRANGE
	var testUtils utils.Utils
	expectedLens := []int{-23, 0, 1, 5, 7, 314}
	actualStrs := []string{}
	customSeed := []rune("ou812")

	// ACT
	for i := range expectedLens {
		str := testUtils.RandomString(expectedLens[i], customSeed)
		actualStrs = append(actualStrs, str)
	}

	// ASSERT
	for i := range actualStrs {
		if len(actualStrs[i]) != int(math.Abs(float64(expectedLens[i]))) {
			t.Errorf("wrong string length; expected %d, got %d", expectedLens[i], len(actualStrs[i]))
		}
		uniqs := testUtils.UniqueRunes(actualStrs[i])
		if !testUtils.ContainsAllRunes(customSeed, uniqs) {
			t.Error("characters found in string not contained i seed set")
		}
	}
}
