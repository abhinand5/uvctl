package env

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/abhinand5/uvctl/internal/config"
	"github.com/abhinand5/uvctl/internal/uv"
)

// List returns all environment names under the root directory.
func List() ([]string, error) {
	root, err := config.GetRoot()
	if err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(root)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("cannot read root directory: %w", err)
	}

	var envs []string
	for _, entry := range entries {
		if entry.IsDir() {
			// Check if it has a .venv directory (valid env)
			venvPath := filepath.Join(root, entry.Name(), ".venv")
			if info, err := os.Stat(venvPath); err == nil && info.IsDir() {
				envs = append(envs, entry.Name())
			}
		}
	}

	return envs, nil
}

// Exists checks if an environment with the given name exists.
func Exists(name string) (bool, error) {
	root, err := config.GetRoot()
	if err != nil {
		return false, err
	}

	envPath := filepath.Join(root, name, ".venv")
	info, err := os.Stat(envPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return info.IsDir(), nil
}

// Path returns the full path to an environment directory.
func Path(name string) (string, error) {
	root, err := config.GetRoot()
	if err != nil {
		return "", err
	}
	return filepath.Join(root, name), nil
}

// VenvPath returns the path to the .venv directory within an environment.
func VenvPath(name string) (string, error) {
	root, err := config.GetRoot()
	if err != nil {
		return "", err
	}
	return filepath.Join(root, name, ".venv"), nil
}

// ActivatePath returns the path to the activate script.
func ActivatePath(name string) (string, error) {
	root, err := config.GetRoot()
	if err != nil {
		return "", err
	}
	return filepath.Join(root, name, ".venv", "bin", "activate"), nil
}

// FishActivatePath returns the path to the fish-specific activate script.
func FishActivatePath(name string) (string, error) {
	root, err := config.GetRoot()
	if err != nil {
		return "", err
	}
	return filepath.Join(root, name, ".venv", "bin", "activate.fish"), nil
}

// Create creates a new environment with the given name and Python version.
// This operation is atomic - on failure, no partial state is left.
func Create(name string, pythonVersion string) error {
	// Check uv is available
	if !uv.Available() {
		return fmt.Errorf("uv not found in PATH")
	}

	// Get root and ensure it exists
	root, err := config.EnsureRoot()
	if err != nil {
		return err
	}

	envPath := filepath.Join(root, name)

	// Check if environment already exists
	if _, err := os.Stat(envPath); err == nil {
		return fmt.Errorf("environment %q already exists", name)
	}

	// Create environment directory
	if err := os.MkdirAll(envPath, 0755); err != nil {
		return fmt.Errorf("cannot create environment directory: %w", err)
	}

	// Create venv using uv - on failure, clean up
	if err := uv.CreateVenv(envPath, pythonVersion, name); err != nil {
		os.RemoveAll(envPath)
		return err
	}

	return nil
}

// Delete removes an environment and all its contents.
func Delete(name string) error {
	exists, err := Exists(name)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("environment %q does not exist", name)
	}

	envPath, err := Path(name)
	if err != nil {
		return err
	}

	if err := os.RemoveAll(envPath); err != nil {
		return fmt.Errorf("cannot delete environment: %w", err)
	}

	return nil
}

// ValidateActivate checks that an environment is ready for activation.
func ValidateActivate(name string) error {
	exists, err := Exists(name)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("environment %q does not exist", name)
	}

	activatePath, err := ActivatePath(name)
	if err != nil {
		return err
	}

	if _, err := os.Stat(activatePath); err != nil {
		return fmt.Errorf("activate script not found at %s", activatePath)
	}

	return nil
}
