# GoGinit 🚀

**GoGinit** is a CLI tool for initializing Go backend projects with ease. Perfect for kickstarting your Go projects with structured templates and framework options!

## Features 🌟

- **Seamless Initialization**: Set up a new Go project in no time.
- **Framework Options**: Choose between Echo, Gin, Fiber, Fuego, or other supported frameworks.
- **Customizable**: Optionally include configuration files and setup scripts.
- **SQLite Database Setup**: Initialize a SQLite database with your project.
- **Start Command**: Easily run your Go project with a single command.

## Installation 🛠️

### Homebrew 🍺 (macOS/Linux)

1. **Tap the Homebrew repository:**

   ```sh
   brew tap pol-cova/goginit
   ```

2. **Install GoGinit:**

   ```sh
   brew install goginit
   ```

### Manual Installation

1. **Download the latest binary from the [releases page](https://github.com/pol-cova/GoGinit/releases).**
2. **Extract the binary and move it to a directory in your `PATH`:**

   ```sh
   tar -xzf goginit_<version>_<os>_<arch>.tar.gz
   sudo mv goginit /usr/local/bin/
   ```

## Usage 🎉

### Initialize a New Project

To initialize a new Go project, run:

```sh
goginit init 
```


### Start the Project

To run the main.go file located in `cmd/projectName/main.go`, use:

```sh
goginit start <projectName>
```

**Options:**

- **`<project-name>`**: Name of the project (required).

### Framework Options

- **`echo`**: For the Echo framework
- **`gin`**: For the Gin framework
- **`fiber`**: For the Fiber framework
- **`fuego`**: For the Fuego framework
- **`martini`**: For the Martini framework
- **`chi`**: For the Chi framework
- **`mux`**: For the Mux framework
- **`GoFr`**: For the GoFr framework
- **`Fuego`**: For the Fuego framework
- **`default`**: For the native `net/http`

## Future Updates 🔮

- **Optimization**: Upcoming updates will focus on optimizing performance and improving overall efficiency.

## Contribution Guidelines 🤝

We welcome contributions! To contribute:

1. **Fork the repository** and clone your fork.
2. **Create a new branch** for your changes.
3. **Make your changes** and test thoroughly.
4. **Submit a pull request** with a detailed description.

## License 📄

GoGinit is licensed under the [MIT License](LICENSE).

## Contact 📫

For support or questions, reach out to us at [polc394@gmail.com](mailto:polc394@gmail.com).

---

**Happy Coding!** 🎉

---
