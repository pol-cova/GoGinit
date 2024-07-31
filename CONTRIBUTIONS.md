# Contributions Guide for GoGinit

Thank you for considering contributing to GoGinit! This guide will help you understand what the project is about, its goals, and how you can contribute effectively.

## What is GoGinit?

GoGinit is a CLI tool designed to initialize Go projects quickly and efficiently. It uses the `bubbletea` package for its terminal user interface (TUI) and supports multiple frameworks including Echo, Gin, Fiber, Martini, Chi, Mux, GoFr, and Native `net/http`. The tool helps in setting up a base folder structure, initializing a Go module, and downloading necessary dependencies, making the process of starting a new Go project seamless.

### Inspiration

GoGinit was inspired by tools like Django-admin and `npm init`, aiming to streamline the setup process for Go projects and make backend development faster and more efficient.

## Goals

1. **Ease of Use**: Simplify the process of initializing new Go projects with a user-friendly CLI.
2. **Flexibility**: Support a wide range of Go frameworks to cater to various project needs.
3. **Efficiency**: Reduce setup time and ensure best practices are followed in the initial project structure.
4. **Extensibility**: Enable the addition of new frameworks and features as the Go ecosystem evolves.
5. **Speed**: Accelerate backend development by automating repetitive tasks.

## Technical Aspects

### Project Structure

The base folder structure created by GoGinit includes the following directories and files:

```
GoGinit/
├── cmd/
│   └── root.go
├── internal/
│   └── tui/
│       └── tui.go
├── templates/
│   └── echo.go (or other framework templates)
├── go.mod
└── README.md
```

- **cmd**: Contains the main entry point of the application.
- **internal**: Contains application-specific code that should not be used outside the project.
- **templates**: Contains reusable templates for different frameworks.
- **go.mod**: The Go module file.
- **README.md**: Project documentation.

### Framework Support

GoGinit supports initializing projects with various frameworks. The tool prompts users to select a framework and sets up the project accordingly. Supported frameworks include:

- **Echo**
- **Gin**
- **Fiber**
- **Martini**
- **Chi**
- **Mux**
- **GoFr**
- **Native `net/http`** (default)

### Dependencies

GoGinit ensures that necessary dependencies are added to the `go.mod` file based on the chosen framework. For example, if Echo is selected, the `github.com/labstack/echo/v4` package is included in the `go.mod` file.

### Bubbletea TUI

The CLI interface is built using the `bubbletea` package, providing an interactive and user-friendly experience for initializing projects.

## How to Contribute

### Reporting Issues

If you encounter any issues or bugs, please report them using the GitHub Issues feature. Provide as much detail as possible, including steps to reproduce the issue, your environment setup, and any relevant logs.

### Suggesting Features

We welcome suggestions for new features or improvements. Please submit a new issue with the tag "feature request" and provide a detailed description of the feature and its potential benefits.

### Submitting Pull Requests

1. **Fork the Repository**: Start by forking the GoGinit repository on GitHub.
2. **Clone Your Fork**: Clone your forked repository to your local machine.
   ```bash
   git clone https://github.com/pol-cova/GoGinit.git
   cd GoGinit
   ```
3. **Create a New Branch**: Create a new branch for your feature or bug fix.
   ```bash
   git checkout -b feature/new-feature
   ```
4. **Make Changes**: Implement your feature or fix. Ensure your code follows the existing style and conventions.
5. **Commit Your Changes**: Write a clear and concise commit message.
   ```bash
   git add .
   git commit -m "Add new feature"
   ```
6. **Push to Your Fork**: Push your changes to your forked repository.
   ```bash
   git push origin feature/new-feature
   ```
7. **Create a Pull Request**: Go to the original GoGinit repository and create a pull request from your fork and branch. Provide a detailed description of your changes and any related issue numbers.

### Code Style and Guidelines

- **Consistency**: Follow the existing code style and conventions.
- **Documentation**: Update documentation as necessary for any new features or changes.
- **Testing**: Write tests for new features or bug fixes to ensure code quality.

### Getting Started with Development

1. **Install Go**: Ensure you have Go installed on your machine. You can download it from [here](https://golang.org/dl/).
2. **Clone the Repository**: Clone the repository to your local machine.
   ```bash
   git clone https://github.com/pol-cova/GoGinit.git
   cd GoGinit
   ```
3. **Install Dependencies**: Run the following command to install the necessary dependencies.
   ```bash
   go mod tidy
   ```
4. **Run the Project**: Use the following command to run the project.
   ```bash
   go run .
   ```

### Future Updates

- **Database Setup**: The next major update will focus on a full database initial setup configuration, enhancing the tool's capabilities for backend projects.
- **Improved Interface**: Continuous improvements to the user interface to enhance user experience.

## Conclusion

We appreciate your contributions to GoGinit! By working together, we can make this tool even more powerful and user-friendly. If you have any questions or need further assistance, feel free to open an issue or contact the maintainers.
