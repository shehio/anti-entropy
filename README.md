# Anti-Entropy

[![Build](https://github.com/shehio/anti-entropy/actions/workflows/build.yml/badge.svg)](https://github.com/shehio/anti-entropy/actions/workflows/build.yml)

A distributed system implementation focusing on anti-entropy protocols for data consistency.

## Prerequisites

- [Bazel](https://bazel.build/install) build system
- Go 1.21 or later
- [Homebrew](https://brew.sh/) (for macOS users)

## Setup

1. Install Go (if not already installed):
   ```bash
   # On macOS
   brew install go
   ```

2. Initialize the Go module (if not already done):
   ```bash
   cd src/anti_entropy
   go mod init github.com/shehio/anti-entropy
   go mod tidy
   ```

3. Install dependencies:
   ```bash
   go get github.com/gorilla/mux
   ```

## Building

To build the project:

```bash
# Build all targets
bazel build //...

# Build specific target
bazel build //src/anti_entropy:anti_entropy

# Run all tests
bazel test //...
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
