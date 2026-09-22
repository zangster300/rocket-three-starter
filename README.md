# Rocket Template

# Usage

## Setup

1. Go
```shell
go mod tidy
```

2.  Datastar + Rocket

```shell
task download # runs `go run cmd/web/downloader/main.go`
```

- Downloads the [`datastar-rocket.js`](https://data-star.dev/reference/rocket) bundle into the `/web/resources/static/datastar/` directory

3. Web Dependencies

```
go run cmd/web/build/main.go
```

- uses [`esbuild`](https://esbuild.github.io/api/#overview)

## Development Mode

```shell
task live
```

OR

```shell
go tool air -build.cmd "go build -tags=dev -o tmp/bin/main ./cmd/web" -build.entrypoint "tmp/bin/main" -misc.clean_on_exit true -build.include_ext "go,templ"

# watch and rebuild web assets + hotreload
go run cmd/web/build/main.go -watch

# watch and rebuild templ components
go tool templ generate -watch
````