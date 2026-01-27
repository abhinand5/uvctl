package config

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	EnvRoot   = "UVCTL_ROOT"
	EnvActive = "UVCTL_ACTIVE"
)

// GetRoot returns the root directory for uvctl environments.
// Resolution order:
// 1. $UVCTL_ROOT if set
// 2. ~/dev/envs (default)
func GetRoot() (string, error) {
	if root := os.Getenv(EnvRoot); root != "" {
		return filepath.Clean(root), nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("cannot determine home directory: %w", err)
	}

	return filepath.Join(home, "dev", "envs"), nil
}

// GetActive returns the currently active environment name, if any.
func GetActive() string {
	return os.Getenv(EnvActive)
}

// EnsureRoot ensures the root directory exists and is writable.
func EnsureRoot() (string, error) {
	root, err := GetRoot()
	if err != nil {
		return "", err
	}

	if err := os.MkdirAll(root, 0755); err != nil {
		return "", fmt.Errorf("cannot create root directory %s: %w", root, err)
	}

	// Check writability by attempting to create a temp file
	testFile := filepath.Join(root, ".uvctl_write_test")
	f, err := os.Create(testFile)
	if err != nil {
		return "", fmt.Errorf("root directory %s is not writable: %w", root, err)
	}
	f.Close()
	os.Remove(testFile)

	return root, nil
}

// IsRootWritable checks if the root directory exists and is writable.
func IsRootWritable() (string, bool) {
	root, err := GetRoot()
	if err != nil {
		return "", false
	}

	info, err := os.Stat(root)
	if err != nil || !info.IsDir() {
		return root, false
	}

	testFile := filepath.Join(root, ".uvctl_write_test")
	f, err := os.Create(testFile)
	if err != nil {
		return root, false
	}
	f.Close()
	os.Remove(testFile)

	return root, true
}
