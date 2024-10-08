package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (u *Utils) parseEnvVars(envFile string) error {
	scanner := bufio.NewScanner(strings.NewReader(envFile))
	for scanner.Scan() {
		line := scanner.Text()
		// ignore enpty lines of comments lines
		if line == "" || strings.HasPrefix(line, "#") {
			// move on to the next line
			continue
		}
		// split lines into key-value-pairs
		kvp := strings.Split(line, "=")
		if len(kvp) != 2 {
			// malformed line, ignore it
			continue
		}
		key := strings.TrimSpace(kvp[0])
		value := strings.TrimSpace(kvp[1])
		// store the key-value-pairs as OS Env Vars
		if err := os.Setenv(key, value); err != nil {
			return fmt.Errorf("error setting %s as en environment variable: %v", key, err)
		}
	}
	return nil
}

// LoadEnvVarsFromFile expects a string represnting the path to the environment variable file.
// This approach relies on the environment variables file existing in the file system and being
// readable.
func (u *Utils) LoadEnvVarsFromFile(filename string) error {
	// read in the file
	file, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading from %s file: %v", filename, err)
	}

	// convert it to a string
	envVars := string(file)

	// parse it's key-value-pairs into env vars
	if err := u.parseEnvVars(envVars); err != nil {
		return err
	}

	return nil
}

// LoadEnvVarsFromEmbed expects a string resulting from having uses Go's //go:embed directive to import an
// embedded environment variable file.
// This approach relies on the environment variables file being embedded directly in the binary, which might
// not be ideal for sensitive data in some cases.
func (u *Utils) LoadEnvVarsFromEmbed(goEmbedReadFile string) error {
	// parse it's key-value-pairs into env vars
	if err := u.parseEnvVars(goEmbedReadFile); err != nil {
		return err
	}

	return nil
}
