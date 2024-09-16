package utils_test

import (
	_ "embed"
	"os"
	"testing"

	"github.com/loickreitmann/utils"
)

//go:embed testdata/.env.example
var envFile string

var expected = map[string]string{
	"API_KEY":      "your_api_key",
	"DATABASE_URL": "your_database_url",
	"PORT":         "8080",
}

func TestUtils_LoadEnvVarsFromEmbed(t *testing.T) {
	// ARRANGE
	var testUtils utils.Utils

	// ACT
	if err := testUtils.LoadEnvVarsFromEmbed(envFile); err != nil {
		t.Errorf("unexpected error loading env vars: %v", err)
	}

	// ASSERT
	for key, value := range expected {
		osKey := os.Getenv(key)
		if osKey != value {
			t.Errorf("epected %s to be %s; got %s", key, value, osKey)
		}
	}
}

func TestUtils_LoadEnvVarsFromFile(t *testing.T) {
	// ARRANGE
	var testUtils utils.Utils

	// ACT
	if err := testUtils.LoadEnvVarsFromFile("testdata/.env.example"); err != nil {
		t.Errorf("unexpected error loading env vars: %v", err)
	}

	// ASSERT
	for key, value := range expected {
		osKey := os.Getenv(key)
		if osKey != value {
			t.Errorf("epected %s to be %s; got %s", key, value, osKey)
		}
	}
}

func TestUtils_LoadEnvVarsFromFile_NoFile(t *testing.T) {
	// ARRANGE
	var testUtils utils.Utils

	// ACT
	err := testUtils.LoadEnvVarsFromFile("testdata/.env.bad.example")

	// ASSERT
	if err == nil {
		t.Error("Expected a \"no such file or directory\" error but got none")
	}
}
