package db

import (
	"fmt"
	"github.com/pol-cova/GoGinit/config"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

// Function to install the required Go package
func installGoPackage(pkg string, projectDir string) {
	fmt.Printf("Installing Go package: %-25s", pkg)

	// Ensure we are in the project directory
	cmd := exec.Command("go", "get", pkg)
	cmd.Dir = projectDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("\nFailed to install package %s: %v", pkg, err)
	}

	// Run `go mod tidy` to clean up go.mod and go.sum
	cmd = exec.Command("go", "mod", "tidy")
	cmd.Dir = projectDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("\nFailed to tidy go.mod: %v", err)
	}

	fmt.Println(" [OK]")
}

// SetupDatabase sets up the database structure
func SetupDatabase(projectName string, setupDB bool) {
	if !setupDB {
		fmt.Println("Skipping database setup...")
		return
	}

	fmt.Println("Setting up the database...")

	// Generate the go.mod file if it does not exist
	if err := config.GenerateGoMod(projectName); err != nil {
		log.Fatalf("Error generating go.mod file: %v", err)
	}

	// Install the SQLite package
	installGoPackage("github.com/mattn/go-sqlite3", projectName)

	// Define the directories
	dbDir := filepath.Join(projectName, "pkg", "db")
	dbFile := filepath.Join(dbDir, fmt.Sprintf("%s.db", projectName))

	// Check if the db directory exists, if not, create it
	if _, err := os.Stat(dbDir); os.IsNotExist(err) {
		fmt.Printf("Creating db directory: %-20s", dbDir)
		if err := os.MkdirAll(dbDir, 0755); err != nil {
			log.Fatalf("\nFailed to create db directory: %v", err)
		}
		fmt.Println(" [OK]")
	} else {
		fmt.Println("DB directory already exists [OK]")
	}

	// Create the database file
	fmt.Printf("Creating database file: %-20s", dbFile)
	file, err := os.Create(dbFile)
	if err != nil {
		log.Fatalf("\nFailed to create database file: %v", err)
	}
	defer file.Close()

	fmt.Println(" [OK]")

	// Optionally generate db.go with init code
	generateDBGoFile(projectName, dbDir)
}

// generateDBGoFile generates the db.go file with initialization code
func generateDBGoFile(projectName, dbDir string) {
	dbGoFile := filepath.Join(dbDir, "db.go")

	tmpl := `package db

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

func InitDB() *sql.DB {
    db, err := sql.Open("sqlite3", "./pkg/db/{{.ProjectName}}.db")
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    // Add any schema setup or other initialization here

    return db
}`

	// Create or overwrite db.go
	fmt.Printf("Generating db.go file: %-20s", dbGoFile)
	file, err := os.Create(dbGoFile)
	if err != nil {
		log.Fatalf("\nFailed to create db.go file: %v", err)
	}
	defer file.Close()

	// Parse and execute the template
	t := template.Must(template.New("dbGo").Parse(tmpl))
	if err := t.Execute(file, struct {
		ProjectName string
	}{
		ProjectName: projectName,
	}); err != nil {
		log.Fatalf("\nFailed to execute template: %v", err)
	}

	fmt.Println(" [OK]")
}
