# `dep` - Dependency Management CLI Tool

`dep` is a modular and flexible command-line tool for managing project dependencies across various package managers. This tool supports commands for updating, removing, adding dependencies, checking for outdated packages, and analyzing security vulnerabilities. It automatically detects the appropriate package manager from a configuration file and can execute commands for various package managers such as `npm`, `pip`, `go`, `yarn`, `pipenv`, `cargo`, and more.

## Features

- **Multi-package manager support**: Automatically detects and works with multiple package managers (e.g., `npm`, `yarn`, `pip`, `go`, `pipenv`, `cargo`, etc.).
- **Custom commands**: Easily extend the functionality by defining custom commands in the configuration file (`config.json`).
- **Execution Info**: Displays system and execution info, including time taken to execute the command and additional system data.
- **Live output**: Provides real-time output of the command execution with progress updates.
- **Cross-platform**: Works across various operating systems (Linux, macOS, Windows).
- **Modular structure**: Easily extendable to support additional package managers.

## Installation

### Prerequisites

- **Go** (version 1.18 or higher)
- The tool automatically detects package managers like `npm`, `yarn`, `go`, `pip`, etc. Ensure the relevant package managers are installed on your system.

### Building the Binary

To build the executable for your system, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/aeswibon/dep.git
   cd dep
   ```

2. Build the executable for your platform:

   ```bash
   go build -o dep
   ```

3. To cross-compile for other platforms (e.g., macOS, Windows):

   ```bash
   GOOS=linux GOARCH=amd64 go build -o dep-linux
   GOOS=darwin GOARCH=amd64 go build -o dep-macos
   GOOS=windows GOARCH=amd64 go build -o dep.exe
   ```

### Using Pre-built Binaries

Alternatively, you can download the pre-built binary for your system from the [releases page](https://github.com/aeswibon/dep#release). Extract the file and add it to your system's PATH for easier access.

## Configuration

The tool uses a `config.json` file to define the package manager and custom commands to be executed.

### Sample `config.json`

```json
{
  "package_manager": "npm",
  "commands": [
    {
      "name": "lint",
      "description": "Run linter",
      "args": ["run", "lint"]
    },
  ]
}
```

- **package_manager**: Defines the package manager (e.g., `npm`, `pip`, `go`, `yarn`, etc.).
- **commands**: List of custom commands that can be executed via the CLI. Each command includes:
  - `name`: The name of the command.
  - `description`: A short description of what the command does.
  - `args`: The arguments to pass to the command when executed.

Place the `config.json` file in the same directory where the tool is executed.

## Usage

Run the tool with one of the following commands:

1. **Basic Command**: Executes a pre-defined command based on the `config.json` file.

   ```bash
   ./dep <command_name>
   ```

   Example:

   ```bash
   ./dep install
   ```

2. **Verbose Mode**: Display detailed execution information, including system info and execution time.

   ```bash
   ./dep install --verbose
   ```

3. **Custom Commands**: If you have custom commands defined in `config.json`, you can execute them by name.

   ```bash
   ./dep <custom_command_name>
   ```

   Example:

   ```bash
   ./dep update
   ```

---

## Command-Line Flags

- **`--verbose`**: Enables detailed logging and system info during command execution.
- **`--config`**: (Optional) Specifies a custom `config.json` file if it’s not located in the current directory.

Example:

```bash
./dep install --verbose
```

## Extending the Tool

### Adding New Package Managers

To add support for a new package manager, follow these steps:

1. Add a new entry in the `commands` section of `config.json`.
2. Create a corresponding handler in the `commands` directory, following the structure of existing package manager handlers.

### Adding Custom Commands

You can define custom commands in the `config.json` file under the `commands` section. Each custom command includes a name, description, and arguments.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please fork the repository, create a new branch for your feature or bug fix, and submit a pull request.
