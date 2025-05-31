# Contributing to Memory Allocation Tests

Thank you for considering contributing to the Memory Allocation Tests project! This document outlines the process for contributing and standards for code contributions.

## Getting Started

1. **Fork the repository** on GitHub
2. **Clone your fork** locally
3. **Set up the project** by following the instructions in README.md

## Adding a New Test

To add a new memory allocation test:

1. **Create a new package** in the `tests/` directory appropriate for your test category
2. **Create model structs** in `model/` directory if needed
3. **Implement your test logic** in a new file, following the existing test structure
4. **Register your test** in `main.go` by adding it to the tests map:

```go
tests := map[string]memory.MemoryTest{
    // Existing tests
    "your-test-name": &yourpackage.YourTest{},
}
```

### Test Structure Guidelines

For consistency, follow these guidelines when implementing a new test:

1. **Separate concerns** - Keep model definitions separate from test logic
2. **Use shared utilities** - Leverage existing utilities in the `pkg/` directory
3. **Follow naming conventions** - Use descriptive names for tests and functions
4. **Document your test** - Include comments explaining what your test measures and why
5. **Report meaningful results** - Add relevant metrics to the `OtherStats` map in the result

## Code Style

Follow standard Go style guidelines:

- Use `gofmt` to format your code
- Follow [Effective Go](https://golang.org/doc/effective_go) style recommendations
- Write meaningful comments and documentation
- Include appropriate error handling

## Adding a New Visualizer

To add a new visualization format:

1. Implement the `Visualizer` interface in the `pkg/visualizer` package
2. Add your visualizer to the factory function in `visualizer.go`
3. Update the documentation to include your new format

## Pull Request Process

1. **Create a branch** with a descriptive name
2. **Make your changes**, following the guidelines above
3. **Add tests** for any new functionality
4. **Update documentation** as necessary
5. **Submit a pull request** with a clear description of the changes

## Code Review

All submissions require review. We use GitHub pull requests for this process.

## Adding Documentation

If you're adding new features, please update:

1. Code comments
2. README.md if appropriate
3. Examples if needed

## Reporting Issues

When reporting issues, please include:

1. A clear description of the problem
2. Steps to reproduce
3. Expected behavior
4. Actual behavior
5. Go version and environment details

## License

By contributing to this project, you agree that your contributions will be licensed under the project's MIT license.
