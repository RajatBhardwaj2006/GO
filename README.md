# Go Learning and Practice Repository

A hands-on Go learning repository containing beginner language exercises, functions and methods, data-type experiments, and HTTP server examples. It also includes an introductory GoFr web application that exposes a `/hello` endpoint.

The repository is intended for learning Go syntax and programming fundamentals while gradually moving toward backend and web development with Go.

## Repository Contents

```text
GO/
├── Golang/
│   ├── intro.go              # Basic Go setup and console output
│   ├── datatypes.go          # Variables and data-type experiments
│   ├── function_method.go    # Functions, structs, and methods
│   ├── https.go              # HTTP server with a GET /event endpoint
│   └── main.go               # GoFr application with a GET /hello endpoint
└── README.md
```

## Go Concepts Covered

### Package and `main` Function

Each source file uses:

```go
package main
```

A Go executable begins execution from the `main()` function. The practice files are separate examples, and each one has its own `main()` function.

### Imports

The examples use Go's standard library packages:

- `fmt` for formatted console input and output
- `net/http` for creating an HTTP server and handling requests

The GoFr example imports:

- `gofr.dev/pkg/gofr` for a lightweight web application and routing API

### Variables and Data Types

The data-type exercise introduces common Go types and declaration styles:

- `string`
- `int`
- `float64`
- `bool`
- `byte`
- Explicit variable declarations with `var`
- Grouped declarations
- Short variable declarations with `:=`

Example:

```go
var name string = "GOFr Event"
flag := true
```

### Functions

`function_method.go` demonstrates named functions with typed parameters and return values:

```go
func calculateSum(num1, num2 int) int {
    return num1 + num2
}
```

This example introduces:

- Function declarations
- Multiple parameters of the same type
- Return types
- Calling functions
- Storing returned values

### Structs

The repository defines a `rectangle` struct with two fields:

```go
type rectangle struct {
    length, breadth int
}
```

Structs allow related values to be grouped into a custom type.

### Methods and Receivers

The rectangle example defines an `area()` method with a value receiver:

```go
func (r rectangle) area() int {
    return r.length * r.breadth
}
```

This demonstrates the difference between a normal function and a method associated with a type.

### HTTP Routing

`https.go` uses the standard `net/http` package to:

- Register a handler for `/event`
- Accept only `GET` requests
- Return HTTP 200 for valid requests
- Return HTTP 405 for unsupported methods
- Start a server on port `8020`

The handler is implemented by `eventHandler`.

### GoFr Web Application

`main.go` demonstrates a GoFr application:

- Creates an application with `gofr.New()`
- Registers a `GET /hello` route
- Returns `Hello GoFr!`
- Starts the application with `app.Run()`

## Technologies Used

- Go
- Go standard library
- `fmt`
- `net/http`
- GoFr (`gofr.dev/pkg/gofr`)

## Requirements

Install Go from the official website:

<https://go.dev/dl/>

Verify the installation:

```bash
go version
```

A recent Go installation is recommended. The repository does not currently include a `go.mod` file, so dependency management and module initialization are required before running the GoFr example.

## Clone the Repository

```bash
git clone https://github.com/RajatBhardwaj2006/GO.git
cd GO
```

## Running the Examples

Because the repository contains multiple independent files with separate `main()` functions, the entire `Golang/` directory should not be run as one package. Run individual examples instead.

### Run the introductory example

```bash
go run Golang/intro.go
```

Expected output:

```text
Go setup working ✔
```

### Run the functions and methods example

```bash
go run Golang/function_method.go
```

This calculates the sum of two integers and the area of a rectangle.

### Run the HTTP server example

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

The endpoint accepts `GET` requests. Other HTTP methods return status `405 Method Not Allowed`.

### Run the GoFr example

First initialize a Go module from the repository root:

```bash
go mod init github.com/RajatBhardwaj2006/GO
go get gofr.dev/pkg/gofr
go run Golang/main.go
```

Then request the endpoint exposed by the application:

```bash
curl http://localhost:8000/hello
```

The exact port is controlled by GoFr's runtime configuration. If the framework uses a different configured port, use that port when sending the request.

## Building an Example

Individual examples can be compiled into executables:

```bash
go build -o bin/intro Golang/intro.go
./bin/intro
```

Build the HTTP example with:

```bash
go build -o bin/http-server Golang/https.go
./bin/http-server
```

## Important Repository Notes

### Files Are Independent Examples

Several files use `package main` and define their own `main()` function. This is useful for learning, but it means that commands such as the following will cause conflicts:

```bash
go run Golang/*.go
go build ./Golang
```

Run or build one example at a time, as shown above.

### Data-Type Example Needs Cleanup

`Golang/datatypes.go` is currently an unfinished practice file. It contains formatting expressions that need correction before it can compile, such as `format"..."` instead of `fmt.Printf("...", ...)`, and it references `f` even though the declared variable is named `i`.

A corrected version of the central example would look like this:

```go
package main

import "fmt"

func main() {
    var name string = "GOFr Event"
    fmt.Printf("Welcome to %v\n", name)

    var i int = 20
    var f float64 = 3.14

    fmt.Printf("i: %v %T\n", i, i)
    fmt.Printf("f: %v %T\n", f, f)

    flag := true
    fmt.Printf("b: %v %T\n", flag, flag)

    var x byte = 'A'
    fmt.Printf("x: %v %T\n", x, x)
}
```

## Suggested Next Steps

To expand this repository into a complete Go learning path, consider adding:

- A `go.mod` file
- Separate folders for each learning topic
- Unit tests using the `testing` package
- Error-handling examples
- Arrays, slices, maps, and loops
- Pointers and interfaces
- Goroutines and channels
- File handling
- JSON encoding and decoding
- REST API projects
- Database connectivity
- Middleware and request validation
- A `Makefile` or task runner
- GitHub Actions for `go fmt`, `go vet`, and tests

## Learning Roadmap

```text
Go Installation and Syntax
        ↓
Variables and Data Types
        ↓
Functions
        ↓
Structs and Methods
        ↓
Slices, Maps, and Pointers
        ↓
Interfaces and Error Handling
        ↓
Goroutines and Channels
        ↓
HTTP Servers and APIs
        ↓
GoFr and Backend Applications
        ↓
Testing, Modules, and Production Practices
```

## Learning Philosophy

> Learn the syntax, write small programs, understand the runtime behavior, and build useful services.

This repository is a work in progress and will grow as more Go concepts, backend examples, and projects are added.

## License

This repository is maintained for educational and practice purposes.
