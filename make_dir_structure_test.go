package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"testing"
)

func TestUtils_MakeDirStructure(t *testing.T) {
	// ARRANGE
	var testUtils Utils
	root, expectedPaths := generateRandomPathsSlice()
	// ACT
	err := testUtils.MakeDirStructure(expectedPaths)
	// ASSERT
	if err != nil {
		t.Errorf("error not expected: %s", err.Error())
	}
	// CLEAN UP
	// get the root path
	err = os.RemoveAll(root)
	if err != nil {
		t.Error(err.Error())
	}
}

func generateRandomPathsSlice() (string, []string) {
	var paths []string
	var u Utils
	max := big.NewInt(12)
	rnd, _ := rand.Int(rand.Reader, max)
	numb := int(rnd.Int64())
	root := "./testdirs/" + u.RandomString(numb)
	for i := 0; i < numb; i++ {
		paths = append(paths, fmt.Sprintf("%s/%s", root, u.RandomString(numb)))
	}
	return root, paths
}
