package utils_test

import (
	"fmt"

	"github.com/loickreitmann/utils"
)

func ExampleUtils_RandomString() {
	var u utils.Utils

	str := u.RandomString(22)

	fmt.Println("Random string of 22 characters:", str)

	// Output:
	// Random string of 22 characters: StkgGpQ+u0lQL+q!5@BfsA
}
