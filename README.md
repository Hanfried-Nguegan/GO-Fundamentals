# Go Fundamentals

A practical, progressive course for learning the foundations of Go from the ground up. Each lesson is kept small and focused so that the language's syntax, conventions, and standard library become familiar through repetition and practice.

## Learning Goals

By the end of this course, you should be able to:

- Read and write idiomatic Go programs.
- Declare variables, constants, and custom types.
- Control program flow with conditions, loops, and `switch` statements.
- Work confidently with arrays, slices, maps, and structs.
- Write reusable functions and methods.
- Handle errors explicitly and design useful error messages.
- Use pointers, interfaces, and Go's composition-based design style.
- Organize code into packages and modules.
- Write tests, benchmarks, and examples with Go's standard tooling.
- Build concurrent programs with goroutines and channels.
- Use the standard library to work with files, JSON, HTTP, and command-line input.

## Prerequisites

- A basic understanding of programming concepts is helpful, but not required.
- Go installed from [go.dev](https://go.dev/dl/).
- A terminal and a code editor such as VS Code.

Check your installation with:

```bash
go version
```

## Course Roadmap

### 1. Getting Started

- [Running Go programs](Learn%20to%20run%20Go/textio.go)
- The `package main` declaration
- The `main` function
- Imports and formatted output
- `go run`, `go build`, and `go fmt`
- Go source file and package conventions

### 2. Variables and Constants

- [Basic variables](Basic%20Variables/basicVariables.go)
- [Declaring a variable](Declaring%20a%20Variable/code.go)
- Explicit variable declarations with `var`
- Short declarations with `:=`
- Constants with `const`
- Zero values
- Type inference
- Scope and naming conventions

### 3. Basic Types and Operators

- Booleans and boolean expressions
- Integers and floating-point numbers
- Strings and runes
- Type conversion
- Arithmetic, comparison, and logical operators
- Bitwise operators and hexadecimal values
- `iota` and enumerated constants

### 4. Control Flow

- `if`, `else if`, and `else`
- Initialization statements in `if`
- `for` loops
- `range`
- `break` and `continue`
- `switch` statements
- `defer`

### 5. Functions

- Parameters and return values
- Multiple return values
- Named return values
- Variadic functions
- Anonymous functions and closures
- Recursion
- Function values

### 6. Collections

- Arrays
- Slices and their length and capacity
- Creating slices with `make`
- Appending and copying
- Two-dimensional slices
- Maps
- Checking whether a map key exists
- Deleting map entries

### 7. Structs and Custom Types

- Defining structs
- Struct literals
- Accessing and updating fields
- Struct embedding
- Defining named types
- Methods and receivers
- Pointer receivers
- JSON tags

### 8. Pointers and Interfaces

- Addresses and dereferencing
- When to use pointers
- `nil`
- Interface values
- Implicit interface implementation
- The empty interface and `any`
- Type assertions
- Type switches

### 9. Errors and Reliability

- The `error` interface
- Returning and checking errors
- Creating errors with `errors.New`
- Wrapping errors with `fmt.Errorf`
- `errors.Is` and `errors.As`
- Custom error types
- `panic` and `recover`
- Input validation

### 10. Packages and Modules

- Package boundaries
- Exported and unexported identifiers
- Import paths
- Initializing a module with `go mod init`
- Managing dependencies with `go mod tidy`
- Package documentation
- Internal packages

### 11. Files and the Standard Library

- Reading and writing files
- Buffered I/O
- Working with directories and paths
- Command-line arguments and flags
- Environment variables
- Dates and times
- JSON encoding and decoding
- Regular expressions
- Useful packages including `fmt`, `strings`, `strconv`, `os`, `io`, and `path/filepath`

### 12. Testing and Tooling

- Table-driven tests
- Test files and `go test`
- Subtests
- Test helpers
- Benchmarks
- Examples
- Code coverage
- `go vet`
- Formatting with `gofmt`
- Static analysis and documentation tools

### 13. Concurrency

- Goroutines
- Channels
- Buffered and unbuffered channels
- Sending and receiving values
- Closing channels
- `select`
- `sync.WaitGroup`
- Mutexes and shared state
- Race detection with `go test -race`
- Context cancellation

### 14. Networking and HTTP

- HTTP clients and requests
- HTTP servers
- Handlers and routing
- Request methods, headers, and status codes
- Query parameters and request bodies
- JSON APIs
- Timeouts and cancellation
- Testing HTTP handlers

### 15. Practical Go Projects

The concepts in this course should eventually be combined into small projects such as:

- A command-line task manager
- A file or directory organizer
- A JSON-backed notes API
- A URL shortener
- A concurrent web scraper
- A small HTTP service with tests

## Running a Lesson

Most early lessons are standalone Go programs. Change into the lesson's directory and provide the file to `go run`:

```bash
cd "Basic Variables"
go run basicVariables.go
```

You can also run every Go file in the current package with:

```bash
go run .
```

Format a lesson before committing it:

```bash
gofmt -w basicVariables.go
```

## Recommended Study Routine

1. Read the lesson and predict what the program will print.
2. Run the program and compare the result with your prediction.
3. Change one value or statement and run it again.
4. Write a small variation without copying the original solution.
5. Record questions and review the related roadmap section.

## Progress Checklist

- [x] Run a basic Go program
- [x] Print text with `fmt.Println`
- [x] Declare and initialize variables
- [ ] Use constants and short declarations
- [ ] Control program flow
- [ ] Write functions
- [ ] Work with slices and maps
- [ ] Define structs and methods
- [ ] Handle errors
- [ ] Create packages and modules
- [ ] Read and write files
- [ ] Write tests
- [ ] Use goroutines and channels
- [ ] Build a complete Go project

The checklist reflects the lessons currently present in the repository and can be updated as new exercises are added.

## Repository Structure

```text
GO-Fundamentals/
├── Basic Variables/
│   └── basicVariables.go
├── Declaring a Variable/
│   └── code.go
├── Learn to run Go/
│   └── textio.go
└── README.md
```

## Go Commands Cheat Sheet

| Command | Purpose |
| --- | --- |
| `go run file.go` | Compile and run one file |
| `go run .` | Compile and run the current package |
| `go build` | Compile the current package |
| `go test ./...` | Run all tests in the module |
| `go fmt ./...` | Format Go source files |
| `go vet ./...` | Report suspicious code patterns |
| `go mod init module-name` | Create a Go module |
| `go mod tidy` | Add missing and remove unused dependencies |
| `go doc package` | Display package documentation |

## Helpful Resources

- [A Tour of Go](https://go.dev/tour/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go standard library documentation](https://pkg.go.dev/std)
- [Go specification](https://go.dev/ref/spec)

## License

This repository is a personal learning project. Add a license here if you decide to share or distribute the code under specific terms.
