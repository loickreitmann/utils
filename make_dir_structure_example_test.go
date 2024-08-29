package utils_test

import (
	"fmt"

	"github.com/loickreitmann/utils"
)

func ExampleUtils_MakeDirStructure() {
	var u utils.Utils

	pathsToMake := []string{
		"./05_test/a/1",
		"./05_test/a/2",
		"./05_test/a/3",
		"./05_test/b/1",
		"./05_test/b/2",
		"./05_test/b/3",
		"./05_test/c/1",
		"./05_test/c/2",
		"./05_test/c/3",
	}
	err := u.MakeDirStructure(pathsToMake)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Created the following directories:")
		u.CrawlLogPaths("05_test")
	}

	// Output:
	// Created the following directories:
	// 05_test
	// 05_test/a
	// 05_test/a/1
	// 05_test/a/2
	// 05_test/a/3
	// 05_test/b
	// 05_test/b/1
	// 05_test/b/2
	// 05_test/b/3
	// 05_test/c
	// 05_test/c/1
	// 05_test/c/2
	// 05_test/c/3
}
