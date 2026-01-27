package uv

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// Available checks if uv is available in PATH.
func Available() bool {
	_, err := exec.LookPath("uv")
	return err == nil
}

// Path returns the path to the uv binary.
func Path() (string, error) {
	return exec.LookPath("uv")
}

// CreateVenv creates a virtual environment at the specified path.
// Output is streamed to stdout/stderr.
func CreateVenv(dir string, pythonVersion string) error {
	args := []string{"venv", ".venv", "--python", pythonVersion}

	cmd := exec.Command("uv", args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("uv venv failed: %w", err)
	}

	return nil
}

// Version returns the uv version string.
func Version() (string, error) {
	cmd := exec.Command("uv", "--version")
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get uv version: %w", err)
	}
	return strings.TrimSpace(string(out)), nil
}

// ListPythons returns available Python versions that uv can use.
func ListPythons() ([]string, error) {
	cmd := exec.Command("uv", "python", "list")
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to list Python versions: %w", err)
	}

	var versions []string
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			// Extract just the version identifier (first column)
			parts := strings.Fields(line)
			if len(parts) > 0 {
				versions = append(versions, parts[0])
			}
		}
	}

	return versions, nil
}

// CheckPython verifies if a specific Python version is available via uv.
func CheckPython(version string) error {
	cmd := exec.Command("uv", "python", "find", version)
	var stderr bytes.Buffer
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderr)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Python %s not available", version)
	}

	return nil
}
