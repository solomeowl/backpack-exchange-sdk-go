# Contributing

Thank you for your interest in contributing to the Backpack Exchange Go SDK!

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/YOUR_USERNAME/backpack-exchange-sdk-go.git`
3. Create a branch: `git checkout -b feature/your-feature`
4. Make your changes
5. Run tests: `make test`
6. Run linter: `make lint`
7. Commit your changes
8. Push to your fork
9. Create a Pull Request

## Development Setup

```bash
# Install dependencies
go mod download

# Run tests
make test

# Run linter (requires golangci-lint)
make lint

# Format code
make fmt

# Build
make build
```

## Code Style

- Follow standard Go conventions
- Use `gofmt` for formatting
- Run `golangci-lint` before committing
- Write tests for new functionality
- Update documentation as needed

## Pull Request Guidelines

1. Keep PRs focused on a single change
2. Include tests for new features
3. Update documentation if needed
4. Follow the existing code style
5. Write clear commit messages

## Reporting Issues

When reporting issues, please include:
- Go version
- SDK version
- Steps to reproduce
- Expected vs actual behavior
- Error messages (if any)

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
