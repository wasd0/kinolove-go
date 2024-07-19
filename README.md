## Requirements

- [docker compose](https://docs.docker.com/compose/install/)
- [goose - migration tool](https://github.com/pressly/goose)
- [jet - sql generator](https://github.com/go-jet/jet)

## Getting started

- ### Install goose

```shell
go install github.com/pressly/goose/v3/cmd/goose@latest
```

This will install the `goose` binary to your `$GOPATH/bin` directory.

For a lite version of the binary without DB connection dependent commands, use the exclusive build tags:

```shell
go build -tags='no_postgres no_mysql no_sqlite3 no_ydb' -o goose ./cmd/goose
```

For macOS users `goose` is available as a [Homebrew Formulae](https://formulae.brew.sh/formula/goose#default):

```shell
brew install goose
```

- ### Install jet

Jet generator can be installed in one of the following ways:

- (Go1.16+) Install jet generator using go install:
```sh
go install github.com/go-jet/jet/v2/cmd/jet@latest
```
*Jet generator is installed to the directory named by the GOBIN environment variable,
which defaults to $GOPATH/bin or $HOME/go/bin if the GOPATH environment variable is not set.*

- Install jet generator to specific folder:
```sh
git clone https://github.com/go-jet/jet.git
cd jet && go build -o dir_path ./cmd/jet
```
*Make sure `dir_path` folder is added to the PATH environment variable.*

- ### Add environment `.env` file to project directory. See `.env.example`.

- ### Run docker compose. Linux example:

```bash
cd path/to/project && docker compose up -d
```

- ### Run makefile from command line. Linux example:

```bash
cd path/to/project && make run
```
