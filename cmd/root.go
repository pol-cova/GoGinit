package cmd

import (
	"fmt"
	"github.com/pol-cova/GoGinit/internal/tui"
	"github.com/pol-cova/GoGinit/templates"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goginit",
	Short: "GoGinit is a CLI tool for initializing Go backend projects",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new Go project",
	Long:  `This command will initialize a new Go backend project with a base template and allow you to choose a framework.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Banner
		fmt.Println(`
		
  _________  ______      _ __ 
 / ___/ __ \/ ___(_)__  (_) /_
/ (_ / /_/ / (_ / / _ \/ / __/
\___/\____/\___/_/_//_/_/\__/ 
                              

			`)
		fmt.Println("Welcome to GoGinit! Let's initialize a new Go project.")
		// Version
		fmt.Println("version: 0.1.0")
		projectName, framework := tui.GetUserInput()
		if projectName == "" || framework == "" {
			fmt.Println("Project name or framework not selected.")
			return
		}
		createProjectSkeleton(projectName, framework)
	},
}

func createProjectSkeleton(projectName, framework string) {
	// Create the necessary directories
	dirs := []string{
		filepath.Join(projectName, "cmd", projectName),
		filepath.Join(projectName, "internal", "handlers"),
		filepath.Join(projectName, "internal", "middleware"),
		filepath.Join(projectName, "internal", "routes"),
		filepath.Join(projectName, "pkg", "models"),
		filepath.Join(projectName, "pkg", "db"),
		filepath.Join(projectName, "pkg", "utils"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	}

	// Create empty files with comments
	files := map[string]string{
		filepath.Join(projectName, "internal", "handlers", "handlers.go"):     "// handlers package",
		filepath.Join(projectName, "internal", "middleware", "middleware.go"): "// middleware package",
		filepath.Join(projectName, "internal", "routes", "routes.go"):         "// routes package",
		filepath.Join(projectName, "pkg", "models", "models.go"):              "// models package",
		filepath.Join(projectName, "pkg", "db", "db.go"):                      "// db package",
		filepath.Join(projectName, "pkg", "utils", "utils.go"):                "// utils package",
	}

	for path, content := range files {
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
	}

	// Initialize Go module
	cmd := exec.Command("go", "mod", "init", projectName)
	cmd.Dir = projectName
	if err := cmd.Run(); err != nil {
		fmt.Println("Error initializing Go module:", err)
		return
	}

	var mainFileContent string
	var getDependenciesCmd *exec.Cmd

	// Select the appropriate template and dependency command
	switch framework {
	case "echo":
		mainFileContent = templates.EchoTemplate
		getDependenciesCmd = exec.Command("go", "get", "github.com/labstack/echo/v4")
	case "gin":
		mainFileContent = templates.GinTemplate
		getDependenciesCmd = exec.Command("go", "get", "github.com/gin-gonic/gin")
	case "fiber":
		mainFileContent = templates.FiberTemplate
		getDependenciesCmd = exec.Command("go", "get", "github.com/gofiber/fiber/v3")
	case "martini":
		mainFileContent = templates.MartiniTemplate
		getDependenciesCmd = exec.Command("go", "get", "github.com/go-martini/martini")
	case "chi":
		mainFileContent = templates.ChiTemplate
		getDependenciesCmd = exec.Command("go", "get", "github.com/go-chi/chi/v5")
	case "mux":
		mainFileContent = templates.MuxTemplate
		getDependenciesCmd = exec.Command("go", "get", "github.com/gorilla/mux")
	case "gofr":
		mainFileContent = templates.GoFrTemplate
		getDependenciesCmd = exec.Command("go", "get", "github.com/gofr-dev/gofr")

	default:
		mainFileContent = templates.DefaultTemplate
	}

	// Create the main.go file in the appropriate directory
	mainFilePath := filepath.Join(projectName, "cmd", projectName, "main.go")
	if err := os.WriteFile(mainFilePath, []byte(mainFileContent), 0644); err != nil {
		fmt.Println("Error creating main.go file:", err)
		return
	}

	// Download the dependencies if required
	if getDependenciesCmd != nil {
		getDependenciesCmd.Dir = projectName
		if err := getDependenciesCmd.Run(); err != nil {
			fmt.Println("Error downloading dependencies:", err)
			return
		}
	}

	fmt.Println("Project initialized successfully!")
}
