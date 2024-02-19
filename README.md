# Go Build Structure CLI

## Introduction
This is a command-line interface (CLI) tool written in Go. It helps you quickly generate the necessary files and folder structures for your Go projects.

## Installation
To install the Go Build Structure CLI, you can use `go get`:

```bash
go get -u github.com/drnugt86837/gbs-cli
```
Please note that this CLI is designed specifically for the `iii go sample` project and requires the installation of packages such as `gin` and `wire`.

## Usage
To use this tool, run the following command:

```bash
gbs-cli [command]
```

### Available Commands:
- `create`: Create a new module structure with predefined files and directories

### Creating a New Module Structure:
To create a new module structure, use the `create` command followed by the module name:

```bash
gbs-cli create <module-name>
```

Replace `<module-name>` with the name of your module. This will generate the necessary files and directories for your module.

#### Example:
To create a module structure for a user module, you can run:

```bash
gbs-cli create user
```

This will create the following structure:
- `user/`
    - `UserController.go`
    - `UserService.go`
    - `UserRoutes.go`
    - `dto/`
        - `CreateUserDto.go`

## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvement, please open an issue or submit a pull request on [GitHub](https://github.com/drnugt86837/gbs-cli).

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
