package utils

import "crypto/rand"

var sourceCharacters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_+!@#$%*")

// RandomString generates and returns a string of a specified length of random characters.
// The characters are randomnly asembled from the source of the following posible characters:
// abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_+!@#$%*
func (u *Utils) RandomString(length int) string {
	// Going for performance over readability. Using the below if statement directly operates on
	// integers without any additional function calls or type conversions.
	// Could have used:
	//  length = int(math.Abs(float64(length)))
	// but it would
	if length < 0 {
		length = -length
	}
	randomString := make([]rune, length)
	for i := range randomString {
		primeNum, _ := rand.Prime(rand.Reader, len(sourceCharacters))
		x, y := primeNum.Uint64(), uint64(len(sourceCharacters))
		randomString[i] = sourceCharacters[x%y]
	}
	return string(randomString)
}
