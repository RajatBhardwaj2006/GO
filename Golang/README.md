# Golang Examples

This directory contains independent Go programs covering Go fundamentals, functions and methods, HTTP servers, and a basic GoFr web application.

## Files

| File | Topics | Description |
|---|---|---|
| [`intro.go`](./intro.go) | Packages, imports, `main()` | Verifies that the Go environment is working and prints a message. |
| [`datatypes.go`](./datatypes.go) | Variables and data types | Practice code for strings, integers, floating-point values, booleans, and bytes. This file currently needs syntax corrections before it will compile. |
| [`function_method.go`](./function_method.go) | Functions, structs, methods | Defines a `rectangle` struct, calculates a sum, and calculates rectangle area through a method receiver. |
| [`https.go`](./https.go) | HTTP server, handlers, routing | Creates a standard-library HTTP server on port `8020` with a `GET /event` endpoint. |
| [`main.go`](./main.go) | GoFr, routes, web applications | Creates a GoFr application with a `GET /hello` endpoint. |

## Running the Examples

Run one file at a time because every example uses `package main` and several files define their own `main()` function.

From the repository root:

```bash
go run Golang/intro.go
go run Golang/function_method.go
```

You can also run the commands from this directory:

```bash
cd Golang
go run intro.go
go run function_method.go
```

Do not run all files together with `go run .` until the examples have been separated into packages or their duplicate `main()` functions have been removed.

## HTTP Server Example

Start the standard-library HTTP server:

```bash
go run Golang/https.go
```

The server listens on port `8020`. In another terminal, send a request:

```bash
curl http://localhost:8020/event
```

Expected response:

```text
Welcome to GoFr's event!
```

The `eventHandler` function accepts `GET` requests and returns HTTP status `405 Method Not Allowed` for other methods.

## GoFr Example

`main.go` uses the GoFr framework and requires a Go module with the GoFr dependency. From the repository root:

```bash
go mod init github.com/RajatBhardwaj2006/GO
go get gofr.dev/pkg/gofr
go run Golang/main.go
```

The application exposes:

```text
GET /hello
```

The route returns:

```text
Hello GoFr!
```

The exact listening port is determined by the GoFr runtime configuration.

## Concepts Practiced

- `package main` and program entry points
- Importing standard-library and third-party packages
- Console output with `fmt`
- Variables and explicit type declarations
- Short variable declarations using `:=`
- Functions with parameters and return values
- Struct definitions and composite literals
- Methods with value receivers
- HTTP handlers using `http.ResponseWriter` and `*http.Request`
- HTTP status codes and request-method validation
- Route registration with `http.HandleFunc`
- Basic GoFr routing and application startup

## Data-Type Example Status

`datatypes.go` is an unfinished learning exercise. Before running it, correct the formatting calls and variable names. For example:

```go
var i int = 20
var f float64 = 3.14

fmt.Printf("i: %v %T\n", i, i)
fmt.Printf("f: %v %T\n", f, f)
```

The current exercise also contains `format"..."` expressions and references `f` without declaring it correctly.

## Recommended Improvements

As this learning directory grows, consider organizing each topic into its own folder and adding:

- A root `go.mod` file
- Unit tests using the `testing` package
- Examples for slices, maps, pointers, interfaces, and errors
- Goroutine and channel exercises
- JSON and REST API examples
- Separate packages for reusable code
- `go fmt`, `go vet`, and automated tests

For the complete repository overview, see the [root README](../README.md).
