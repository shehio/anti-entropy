# Anti-Entropy Protocol Implementation

[![Build](https://github.com/shehio/anti-entropy/actions/workflows/build.yml/badge.svg)](https://github.com/shehio/anti-entropy/actions/workflows/build.yml)

A Go implementation of the anti-entropy protocol for distributed systems, featuring a Merkle tree for efficient state synchronization, primarily targeting data consistency.

### Features

- Distributed node communication
- Merkle tree-based state verification
- Gossip protocol for state propagation
- Version-based conflict resolution
- Configurable update intervals

## Project Structure

```
src/anti_entropy/
├── main.go           # Main program entry point
├── node/            # Node implementation
│   ├── node.go      # Node structure and methods
│   └── node_test.go # Node tests
└── merkle/          # Merkle tree implementation
    ├── merkle_node.go  # Merkle node structure
    ├── merkle_tree.go  # Merkle tree operations
    └── merkle_test.go  # Merkle tree tests
```

### Prerequisites

- Go 1.21 or later
- Bazel 8.1.1 or later

### Installation

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

## Testing

The project includes comprehensive tests for both the node and Merkle tree implementations:

```bash
# Run node tests
bazel test //src/anti_entropy/node:node_test

# Run Merkle tree tests
bazel test //src/anti_entropy/merkle:merkle_test
```

## Development Setup

## License

This project is licensed under the MIT License - see the LICENSE file for details.
