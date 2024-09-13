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
