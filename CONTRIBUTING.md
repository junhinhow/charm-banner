# Contributing to charm-banner

Thanks for your interest in contributing!

## How to contribute

1. Fork the repository
2. Create a feature branch (`git checkout -b feat/my-feature`)
3. Write tests for your changes
4. Ensure all tests pass (`go test ./...`)
5. Commit your changes using [conventional commits](https://www.conventionalcommits.org/)
6. Push to your branch and open a Pull Request

## Development

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run linter
go vet ./...
```

## Code style

- Follow standard Go conventions (`gofmt`, `go vet`)
- Write descriptive variable and function names
- Add comments in Portuguese for internal documentation
- Keep functions small and focused

## Adding new characters

To add support for new characters in the font maps:

1. Add the character pattern to `smallFont`, `mediumFont`, and `largeFont` in `banner.go`
2. Each pattern must have the correct number of rows for the font size (3, 5, or 7)
3. Add a test case in `banner_test.go`

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
