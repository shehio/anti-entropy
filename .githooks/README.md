# Git Hooks

This directory contains git hooks used in the anti-entropy project.

## Available Hooks

### pre-commit
Runs `bazel build` and `bazel test` before each commit to ensure code quality.

## Installation

To install the hooks, run the following commands from the repository root:

```bash
# Make hooks executable
chmod +x .githooks/*

# Set git to use this directory for hooks
git config core.hooksPath .githooks
```

## Notes
- The pre-commit hook will prevent commits if the build or tests fail
- Tests are run with `--nocache_test_results` to ensure fresh test runs 