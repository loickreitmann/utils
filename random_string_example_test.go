package utils_test

import (
	"fmt"

	"github.com/loickreitmann/utils"
)

func ExampleUtils_RandomString() {
	var u utils.Utils

	str := u.RandomString(22)

	fmt.Printf("Random string has %d characters.\n", len(str))

	// Output:
	// Random string has 22 characters.
}
