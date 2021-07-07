# Contributing to CLI

Thank you for contributing to the CLI. See below for guides to help you contribute.

## Prerequisites

The following will need to be installed on your device in order to contribute to this project.

- [Go >= v1.16](https://golang.org/dl)
- [golangci-lint](https://golangci-lint.run)
- [GoReleaser](https://goreleaser.com/) (optional)

## Development

### Running CLI commands

You can run the CLI commands files directly with `go run`. All commands and subcommands are in the `./pkg/cmd` folder.

```shell
go run ./cmd/starter
```

### Development commands
 

#### `make lint`

Runs a linter on the Go source code. Configuration can be found in `.golangci.yaml`.
There are a number of lint rules enabled. You can find a full list of rules [here](https://golangci-lint.run/usage/linters/) with usage and configuration guides.

#### `make install`

Builds a binary in the `$GOPATH/bin` directory. Can be executed globally as it is in your `$PATH`.

#### `make binary`

Builds an executable binary of the CLI in the project root. Executable only inside the workspace.

#### `make format`

Formats source code.


### Testing

If you have the Go extension for VS Code, you can generate test stubs for a file, package or function. See [Go#Test](https://code.visualstudio.com/docs/languages/go#_test)

### `make test/unit`

Runs unit tests

## Internationalization

All text strings are placed in `./pkg/localize/locales` directory.

## Documentation

The main CLI documentation source files are stored in the `./pkg/localize/locales/en/cmd/` directory.

The CLI documentation output is generated in the `./docs` directory.

### Generating documentation

Documentation can be generated from the CLI commands.

```shell
make docs/generate
```

#### `make docs/generate`

After running the command, the documentation should be generated in AsciiDoc format.

## Best practices

- [Command Line Interface Guidelines](https://clig.dev/) is a great resource for writing command-line interfaces.
- Write clear and meaningful Git commit messages following the [Conventional Commits specification](https://www.conventionalcommits.org)
- Provide clear documentation comments.
- Make sure you include a clear and detailed PR description, linking to the related issue when it exists.
- Check out [CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments) when writing and reviewing Go code.

## Releases

This project follows [Semantic Versioning](https://semver.org/). Before creating a release, identify if it will be a major, minor, or patch release. In the following example, we will create a patch release `0.20.1`.

> NOTE: When creating a release it is good practice to create a pre-release first to monitor the integration tests to ensure everything works as it should.

### Create snapshot

For testing purposes we should always release a local snapshot version for testing (requires [GoReleaser](https://goreleaser.com/))

```shell
goreleaser --snapshot --rm-dist
```

### Creating the release

Execute `git tag v0.20.1` to create the release tag. Then execute `git push origin v0.20.1` to push to the tag to your remote (GitHub).
Once pushed, a [GitHub Action](https://github.com/aerogear/charmil/actions/workflows/release.yml) will create a release on GitHub and upload the binaries.

> NOTE: To create a pre-release, the tag should have appropriate suffix, e.g v0.20.1-alpha1

### Generate a changelog

> NOTE: This step is not required for pre-releases.
> NOTE: This step is automated using github actions.

[git-chglog](https://github.com/git-chglog/git-chglog) is used to generate a changelog for the current release.

Run `./scripts/generate-changelog.sh` to output the changes between the current and last stable releases. Paste the output into the description of the [release on GitHub](https://github.com/aerogear/charmil/releases/tag/latest).
