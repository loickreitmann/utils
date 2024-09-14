package utils_test

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/loickreitmann/utils"
)

/*
The exmple .env.example file contains:
```
API_KEY=your_api_key
# this a comment ignored by ven var parsing
DATABASE_URL=your_database_url

PORT=8080

This=is a bad=line
```
*/

func ExampleUtils_LoadEnvVarsFromFile() {
	var u utils.Utils

	if err := u.LoadEnvVarsFromFile("testdata/.env.example"); err != nil {
		fmt.Printf("unexpected error loading env vars: %v\n", err)
	}

	fmt.Println("API_KEY", os.Getenv("API_KEY"))
	fmt.Println("DATABASE_URL", os.Getenv("DATABASE_URL"))
	fmt.Println("PORT", os.Getenv("PORT"))
	fmt.Println("This", os.Getenv("This")) // will be an empty string

	// Output:
	// API_KEY your_api_key
	// DATABASE_URL your_database_url
	// PORT 8080
	// This
}
