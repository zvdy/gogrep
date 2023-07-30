# Go Grep

Go Grep is a simple version of the `grep` utility written in Go. It allows you to search for a specific pattern in a single file or recursively in a directory and its subdirectories. The program uses Goroutines to efficiently search for the pattern in multiple files concurrently.

## Table of Contents

- [Usage](#usage)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Contributing](#contributing)
- [License](#license)

## Usage

```bash
./grep [-R] <pattern> <file/dir> [<file/dir> ...]

    -R: Recursive search (optional). If provided, the program searches for the pattern in all files within the specified directory and its subdirectories.
    <pattern>: The pattern to search for.
    <file/dir>: The file or directory path(s) to search in.
```

 ## Getting Started
 
 ### Prerequisites
---
Before running the program, ensure you have Go installed on your system.

### Installation
---

    Clone the repository or download the grep.go file.

    Compile the Go code to create the executable:

```bash

go build grep.go
```
## Contributing

Contributions to this project are welcome! If you find any issues or have ideas for improvements, feel free to open an issue or submit a pull request.

    Fork the repository.
    Create your branch (git checkout -b feature/YourFeature).
    Commit your changes (git commit -m 'Add some feature').
    Push to the branch (git push origin feature/YourFeature).
    Open a pull request.

## License

This project is licensed under the [MIT License](https://mit-license.org/).