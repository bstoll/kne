# Agent Instructions

## Building and Testing

When working on this project, please use the provided `Makefile` targets to build and test the codebase:

- **Build**: Use `make build` to compile the `kne` CLI binary.
- **Test**: Use `make test` to run unit tests. Note that this target automatically ignores end-to-end tests located under the `cloudbuild/` directory.
- **Install**: Use `make install` to build and install the `kne` CLI binary to your local bin directory (`/usr/local/bin`).

## Code Style Guidelines

All Go code written for this project must strictly adhere to standard Go conventions and Google's specific Go style guidelines. When creating or modifying code, ensure compliance with the following resources:

1. [Effective Go](https://go.dev/doc/effective_go)
2. [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments)
3. [Google Style Guide (Go)](https://google.github.io/styleguide/go/guide)
4. [Go Style Decisions](https://google.github.io/styleguide/go/decisions)
5. [Go Best Practices](https://google.github.io/styleguide/go/best-practices)

Please refer to these documents for any stylistic or design questions when writing Go code in this repository.
