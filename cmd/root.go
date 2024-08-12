package cmd

import (
	"fmt"
	"github.com/pol-cova/GoGinit/internal/db"
	"github.com/pol-cova/GoGinit/templates"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pol-cova/GoGinit/config"
	"github.com/pol-cova/GoGinit/internal/tui"
	"github.com/spf13/cobra"
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
	rootCmd.AddCommand(startCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new Go project",
	Long:  `This command will initialize a new Go backend project with a base template and allow you to choose a framework.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Banner
		fmt.Println(`
   ____              ____   _           _   _   
  / ___|   ___      / ___| (_)  _ __   (_) | |_ 
 | |  _   / _ \    | |  _  | | | '_ \  | | | __|
 | |_| | | (_) |   | |_| | | | | | | | | | | |_ 
  \____|  \___/     \____| |_| |_| |_| |_|  \__|
			`)
		fmt.Println("Welcome to GoGinit! Let's initialize a new Go project.")

		// Call the TUI to get user input
		projectName, framework, setupDB := tui.GetUserInput()
		if projectName == "" || framework == "" {
			fmt.Println("Project name or framework not selected.")
			return
		}
		// Create the project skeleton and handle any additional setup
		createProjectSkeleton(projectName, framework, setupDB)
	},
}

var startCmd = &cobra.Command{
	Use:   "start [projectName]",
	Short: "Start the backend server",
	Long:  `This command will run the main.go file located in cmd/projectName/main.go`,
	Args:  cobra.ExactArgs(1), // Ensure exactly one argument is provided
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0] // Get the project name from arguments
		runMain(projectName)
	},
}

func runMain(projectName string) {
	mainPath := filepath.Join("cmd", projectName, "main.go")
	cmd := exec.Command("go", "run", mainPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Running server: %s\n", mainPath)
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to run server: %v\n", err)
	}
}

func createProjectSkeleton(projectName, framework string, setupDB bool) {
	// Create the necessary directories
	dirs := []string{
		filepath.Join(projectName, "cmd", projectName),
		filepath.Join(projectName, "internal", "handlers"),
		filepath.Join(projectName, "internal", "middleware"),
		filepath.Join(projectName, "internal", "routes"),
		filepath.Join(projectName, "pkg", "models"),
		filepath.Join(projectName, "pkg", "db"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
	}

	// Create empty files with comments
	files := map[string]string{
		filepath.Join(projectName, "internal", "handlers", "handlers.go"):     "// handlers package\npackage handlers",
		filepath.Join(projectName, "internal", "middleware", "middleware.go"): "// middleware package\npackage middleware",
		filepath.Join(projectName, "internal", "routes", "routes.go"):         "// routes package\npackage routes",
		filepath.Join(projectName, "pkg", "models", "models.go"):              "// models package\npackage models",
	}

	// Conditionally add the db.go file
	if setupDB {
		files[filepath.Join(projectName, "pkg", "db", "db.go")] = "// db package\npackage db"
	}

	for path, content := range files {
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
	}

	// Initialize Go module
	if err := config.GenerateGoMod(projectName); err != nil {
		fmt.Println("Error generating go.mod file:", err)
		return
	}

	frameworkConfig, err := config.GetFrameworkConfig(framework)
	if err != nil {
		fmt.Println("Error getting framework configuration:", err)
		return
	}

	if err := config.FetchFrameworkDependencies(projectName, frameworkConfig.Name); err != nil {
		fmt.Println("Error getting framework dependencies:", err)
		return
	}

	var mainFileContent string

	// Select the appropriate template
	switch framework {
	case "echo":
		mainFileContent = templates.EchoTemplate
	case "gin":
		mainFileContent = templates.GinTemplate
	case "fiber":
		mainFileContent = templates.FiberTemplate
	case "martini":
		mainFileContent = templates.MartiniTemplate
	case "chi":
		mainFileContent = templates.ChiTemplate
	case "mux":
		mainFileContent = templates.MuxTemplate
	case "gofr":
		mainFileContent = templates.GoFrTemplate
	case "fuego":
		mainFileContent = templates.FuegoTemplate
	case "default":
		mainFileContent = templates.DefaultTemplate
	default:
		fmt.Println("Invalid framework selected.")
		return
	}

	// Create the main.go file in the appropriate directory
	mainFilePath := filepath.Join(projectName, "cmd", projectName, "main.go")
	if err := os.WriteFile(mainFilePath, []byte(mainFileContent), 0644); err != nil {
		fmt.Println("Error creating main.go file:", err)
		return
	}

	// Setup the database if required
	if setupDB {
		db.SetupDatabase(projectName, setupDB)
	}

	fmt.Println("Project created successfully, Happy Coding! 🎉")
}
