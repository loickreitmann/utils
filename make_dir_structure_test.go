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
func TestUtils_MakeDirStructure_Error(t *testing.T) {
	// ARRANGE
	var testUtils Utils
	notPermittedRoot := "./permission_denied_path"
	// creat a directory with permission to block the creation
	_ = os.Mkdir(notPermittedRoot, 0555)
	testPath := notPermittedRoot + "/sub_path"
	// ACT
	err := testUtils.MakeDirStructure([]string{testPath})
	// ASSERT
	if err == nil {
		t.Error("error was expected")
	}
	// CLEAN UP
	// get the root path
	err = os.RemoveAll(notPermittedRoot)
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
