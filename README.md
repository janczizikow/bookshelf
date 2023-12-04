# bookshelf 📚

## Installation

```sh
go mod download
```

## Usage

### Start the HTTP server

```sh
make run
```

### Adding new migration

install [golang migrate](https://github.com/golang-migrate/migrate)

```sh
migrate create -ext sql -dir ./internal/database/migrations {name}
```
