package utils_test

import (
	"testing"

	"github.com/loickreitmann/utils"
)

var runeTestCases = []struct {
	name     string
	set      []rune
	subset   []rune
	contains bool
}{
	{name: "contained", set: []rune("ou812"), subset: []rune("o2"), contains: true},
	{name: "not contained", set: []rune("8675309"), subset: []rune("wer9"), contains: false},
}

func TestUtils_ContainsAllRunes(t *testing.T) {
	// ARRANGE
	var testUtils utils.Utils

	for _, testCase := range runeTestCases {
		// ACT & ASSERT
		if testUtils.ContainsAllRunes(testCase.set, testCase.subset) != testCase.contains {
			t.Errorf("[%s]: expected containment to be %v", testCase.name, testCase.contains)
		}
	}
}
