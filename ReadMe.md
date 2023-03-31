# SpInvalidFileNameFinder

SpInvalidFileNameFinder is a command-line tool written in Go that helps you find and optionally rename files and folders with invalid names for SharePoint.

## Features

- Find files and folders with invalid names for SharePoint.
- Log invalid file and folder names to a log file.
- Optionally rename files and folders with invalid names.
- Specify the directory to process or use the current working directory by default.
- Skip files and folders that start with a tilde (~).
- Skip files and folders that start with a dot (.).

## Installation

To install SpInvalidFileNameFinder, make sure you have Go installed on your system, and then run the following command:

```sh
go get github.com/yourusername/SpInvalidFileNameFinder
```

## Usage

To use SpInvalidFileNameFinder, navigate to the directory containing the `main.go` file and run the following command:

```sh
go run main.go [flags]
```

### Flags

- `-rename`: Enable renaming of files and folders with invalid names (default: `false`).
- `-dir`: Specify the directory to process (defaults to the current working directory).

### Examples

- To find invalid file and folder names without renaming them:

  ```sh
  go run main.go
  ```

- To find invalid file and folder names and rename them:

  ```sh
  go run main.go -rename
  ```

- To find invalid file and folder names in a specific directory:

  ```sh
  go run main.go -dir "/path/to/your/directory"
  ```

- To find invalid file and folder names in a specific directory and rename them:

  ```sh
  go run main.go -dir "/path/to/your/directory" -rename
  ```

## Log File

SpInvalidFileNameFinder logs the invalid file and folder names to a file called `invalid_filenames.log` in the directory being processed.

## License

SpInvalidFileNameFinder is released under the [MIT License](LICENSE).
