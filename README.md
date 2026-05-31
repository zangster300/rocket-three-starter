# Rocket Template

# Usage

## Setup

1. Go
```
go mod tidy
```

2.  Datastar Pro + Rocket

- Grab [`datastar-pro.js`](https://data-star.dev/pro/download) and drop it into the `/web/resources/static/datastar/` directory

3. Web Dependencies

```
go run cmd/web/build/main.go
```

- uses [`esbuild`](https://esbuild.github.io/api/#overview)
- can be used to bundle dependencies

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