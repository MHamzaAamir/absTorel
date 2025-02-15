# Relative Path to Absolute Path Converter

This Go program converts a **relative path** into an **absolute path** based on a given **current directory**. It resolves paths accurately, handling `./` (current directory), `../` (parent directory), and nested paths correctly.

## Features

- Converts relative paths to absolute paths efficiently.
- Handles various edge cases, including multiple `..`, `.` and redundant slashes.
- **Pre-built executable (**``**) is included**, so you don’t need to install Go to run it.

## Usage

You can run this program using one of the following methods:

### 1. Running the Executable (`relative.exe`)

#### If You Don't Have Go Installed

You can directly run the pre-built `relative.exe` file:

- **For Command Prompt (CMD):**
  ```sh
  relative.exe "<current_directory>" "<relative_path>"
  ```
- **For PowerShell:**
  ```powershell
  .\relative.exe "<current_directory>" "<relative_path>"
  ```


### 2. Running the Executable (`relative.exe`) Using Go

If you have Go installed, you can still run the `.exe` file using the `go run` command:

```sh
go run main.go "<current_directory>" "<relative_path>"
```

## Author

**Muhammad Hamza Aamir** This project is open-source. Feel free to modify and improve it! 🚀

