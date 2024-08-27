// Package of simple and well tested reusable Go utilities.
// Inspired by Trevor Sawler's "Building a Module in Go." Udemy course.
package utils

import "crypto/rand"

var sourceCharacters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_+!@#$%*")

// Utils is the type used to instantiate the module. Any variable of this type will have access to all
// the methods with the *Utils receiver.
type Utils struct{}

// RandomString generaters and returns a sstring of random characters of a specified length.,
// The characters are randomnly asembled from the source of the following posible characters:
// abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_+!@#$%*
func (o *Utils) RandomString(length int) string {
	randomString := make([]rune, length)
	for i := range randomString {
		primeNum, _ := rand.Prime(rand.Reader, len(sourceCharacters))
		x, y := primeNum.Uint64(), uint64(len(sourceCharacters))
		randomString[i] = sourceCharacters[x%y]
	}
	return string(randomString)
}
