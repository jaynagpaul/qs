package qs

import (
	"errors"
	"os"
	"path/filepath"
)

// Run will parse the directory at path in search of a quickstart.toml.
// After finding the file it will call ExecFile on it.
func Run(path string) error {
	p, err := findQS(path)

	if err != nil {
		return err
	}
}

// Search the directory for a quickstart.toml file
func findQS(path string) (string, error) {
	if p := filepath.Join(path, "quickstart.toml"); fileExists(p) {
		return p, nil
	}

	return "", errors.New("No quickstart.toml")
}

func fileExists(path string) bool {
	stat, err := os.Stat(path)

	return err == nil
}
