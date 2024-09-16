package utils

import "crypto/rand"

var sourceCharacters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_+!@#$%*")

// RandomString generates and returns a string of a specified length of random characters.
// The characters are randomnly asembled from othe ptional characterSeed parameter, or they
// are sourced from the default set made up of the following characters:
// abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_+!@#$%*
func (u *Utils) RandomString(length int, characterSeed ...[]rune) string {
	// Going for performance over readability. Using the below if statement directly operates on
	// integers without any additional function calls or type conversions.
	// Could have used:
	//  length = int(math.Abs(float64(length)))
	// but it would
	if length < 0 {
		length = -length
	}
	seedSource := sourceCharacters
	if len(characterSeed) > 0 {
		seedSource = characterSeed[0]
	}
	randomString := make([]rune, length)
	for i := range randomString {
		primeNum, _ := rand.Prime(rand.Reader, len(seedSource))
		x, y := primeNum.Uint64(), uint64(len(seedSource))
		randomString[i] = seedSource[x%y]
	}
	return string(randomString)
}
