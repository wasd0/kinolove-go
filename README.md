## Requirements

- [docker compose](https://docs.docker.com/compose/install/)
- [goose - migration tool](https://github.com/pressly/goose)

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

- ### Add environment `.env` file to project directory. See `.env.example`.

- ### Run docker compose. Linux example:

```bash
cd path/to/project && docker compose up -d
```

- ### Run makefile from command line. Linux example:

```bash
cd path/to/project && make run
```
