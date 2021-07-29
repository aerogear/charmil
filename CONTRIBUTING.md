# Contributing

## Prerequisites
The following should be to be installed on your device in order to contribute to this project.

- [Go >= v1.16](https://golang.org/dl)
- [golangci-lint](https://golangci-lint.run/)
- [Yarn](https://classic.yarnpkg.com/)

## Development

### Folder Structure
Charmil repository contains a number of components that can be consumed independently. In the root directory, the following folders can be found:

- **cli**  -command line tool for using charmil.
- **core** - charmil SDK 
- **starter** - charmil starter template project
- **validator** - validator and linter library to test cobra commands

### Running Charmil CLI Commands
You can run the CLI commands files directly with go run. All commands and subcommands are in the ./pkg/cmd folder.
```bash
go run ./cli/cmd/charmil
```

### Development Commands

#### `make setup/githooks`
Run the command to set up git hooks for the project. The following git hooks are currently available:
- **pre-commit** - This runs checks to ensure that the staged `.go` files passes formatting and standard checks using gofmt and go vet.

#### `make lint`
Runs a linter on the Go source code. Configuration can be found in [.golangci.yaml](./.golangci.yaml). There are a number of lint rules enabled. You can find a full list of rules here with usage and configuration guides.

#### `make test/unit`
Runs unit tests

## Community
- We are keeping all the communications open, so that everyone can sync and is free to contribute. So if you have any feature/bugs suggestions about anything please don't hesitate to open up an [issue](https://github.com/aerogear/charmil/issues/new/choose)
- You can join [Aerogear’s discord server](https://discord.gg/hsDJUPkAWH) to participate in the discussions happening

## License
By contributing, you agree that your contributions will be licensed under its Apache License 2.0
