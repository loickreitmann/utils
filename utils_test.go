package utils

import (
	"testing"
)

func TestUtils_RandomString(t *testing.T) {
	// arrange
	var testUtils Utils
	expectedLens := []int{1, 5, 7, 22, 36, 789, 10000}
	actualLens := []int{}

	// act
	for i := range expectedLens {
		str := testUtils.RandomString(expectedLens[i])
		actualLens = append(actualLens, len(str))
	}

	// assert
	for i := range actualLens {
		if actualLens[i] != expectedLens[i] {
			t.Errorf("wrong string length; expected %d, got %d", expectedLens[i], actualLens[i])
		}
	}
}
