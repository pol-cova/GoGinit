package config

import (
	"fmt"
	"github.com/pol-cova/GoGinit/templates"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type FrameworkConfig struct {
	Name     string
	Template string
}

// GetFrameworkConfig returns the configuration for the specified framework
func GetFrameworkConfig(framework string) (FrameworkConfig, error) {
	configs := map[string]FrameworkConfig{
		"echo":    {"github.com/labstack/echo/v4", templates.EchoTemplate},
		"gin":     {"github.com/gin-gonic/gin", templates.GinTemplate},
		"fiber":   {"github.com/gofiber/fiber/v3", templates.FiberTemplate},
		"martini": {"github.com/go-martini/martini", templates.MartiniTemplate},
		"chi":     {"github.com/go-chi/chi/v5", templates.ChiTemplate},
		"mux":     {"github.com/gorilla/mux", templates.MuxTemplate},
		"gofr":    {"github.com/gofr-dev/gofr", templates.GoFrTemplate},
		"fuego":   {"github.com/go-fuego/fuego", templates.FuegoTemplate},
		"default": {"", templates.DefaultTemplate},
	}

	config, ok := configs[framework]
	if !ok {
		return FrameworkConfig{}, fmt.Errorf("unknown framework: %s", framework)
	}

	return config, nil
}

// GetGoVersion fetches the current Go version
func GetGoVersion() (string, error) {
	cmd := exec.Command("go", "version")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	parts := strings.Fields(string(output))
	if len(parts) < 3 {
		return "", fmt.Errorf("unexpected output from go version: %s", output)
	}

	version := strings.TrimPrefix(parts[2], "go")
	return version, nil
}

// GenerateGoMod initializes the Go module using `go mod init`
// GenerateGoMod initializes the Go module using `go mod init`
// if it hasn't already been initialized.
func GenerateGoMod(moduleName string) error {
	if moduleName == "" {
		return fmt.Errorf("moduleName cannot be empty")
	}

	// Check if go.mod already exists
	goModPath := filepath.Join(moduleName, "go.mod")
	if _, err := os.Stat(goModPath); !os.IsNotExist(err) {
		// go.mod already exists
		fmt.Printf("go.mod already exists in %s, skipping initialization\n", moduleName)
		return nil
	}

	// Initialize the Go module
	cmd := exec.Command("go", "mod", "init", moduleName)
	cmd.Dir = moduleName
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error initializing Go module: %v, Output: %s", err, output)
	}

	// Optionally, fetch the Go version (though `go mod init` handles it automatically)
	goVersion, err := GetGoVersion()
	if err != nil {
		return fmt.Errorf("error getting Go version: %v", err)
	}
	fmt.Printf("Initialized Go module %s with Go version %s\n", moduleName, goVersion)

	return nil
}

// FetchFrameworkDependencies fetches the latest version of the framework dependencies
func FetchFrameworkDependencies(moduleName, framework string) error {
	if moduleName == "" {
		return fmt.Errorf("moduleName cannot be empty")
	}

	// Change the directory to the project directory
	cmd := exec.Command("go", "get", framework)
	cmd.Dir = moduleName
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error fetching framework dependencies: %v, Output: %s", err, output)
	}
	return nil
}
