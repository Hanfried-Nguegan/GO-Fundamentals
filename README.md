# Go Fundamentals

This repository is a hands-on Go learning workspace covering the foundations of the Go language, common patterns, and a small REST API project built with Gin and SQLite.

It contains many small standalone examples and mini-projects, each designed to reinforce a specific concept. The goal is to move from basic syntax and structures to practical application development in a structured, progressive way.

## Why this workspace exists

The project is organized as a collection of small exercises and sample apps rather than a single large application. That makes it easy to:

- learn one concept at a time,
- run examples in isolation,
- practice Go syntax and tooling,
- build toward a real backend project.

## Prerequisites

Before running the code in this workspace, make sure you have:

- Go installed on your machine
- A terminal and code editor, preferably VS Code
- Basic familiarity with programming concepts

Check your installation:

```bash
go version
```

## Workspace structure

At the root of the project you will find a mix of standalone beginner exercises and a few more advanced mini projects.

### Core learning folders

- [Arrays](Arrays)
- [Basic Variables](Basic%20Variables)
- [Control Structures](Control%20Structures)
- [Converting Between Types](Converting%20Between%20Types)
- [Declaring a Variable](Declaring%20a%20Variable)
- [Functions](Functions)
- [Go Essentials](Go%20Essentials)
- [Go Routines](Go%20Routines)
- [Interfaces](Interfaces)
- [Learn to run Go](Learn%20to%20run%20Go)
- [More Structs](More%20Structs)
- [Pointers](Pointers)
- [Practice](Practice)
- [Practice Slices](Practice%20Slices)
- [Price Calculator](Price%20Calculator)
- [Profit Calculator](Profit%20Calculator)
- [Same Line Declaration](Same%20Line%20Declaration)
- [Sandbox](Sandbox)
- [Short Variable Declaration](Short%20Variable%20Declaration)
- [Static Types](Static%20Types)
- [Structs](Structs)
- [Structure Practice](Structure%20Practice)

### Project-style examples

- [REST API PROJECT](REST%20API%20PROJECT) — a small CRUD API built with Gin and SQLite.

## Learning path

The workspace roughly follows a typical Go fundamentals progression:

1. Variables, constants, and basic types
2. Control flow and common operators
3. Functions and method design
4. Arrays, slices, maps, and structs
5. Pointers, interfaces, and standard library usage
6. Error handling and reusable patterns
7. File I/O and JSON handling
8. Concurrency with goroutines
9. HTTP APIs and database-backed services

Each folder is intentionally small and focused so you can experiment without getting lost in a large codebase.

## Running examples

Most folders are standalone Go modules and can be run from within their own directory.

### Example: run a simple program

```bash
cd "Basic Variables"
go run basicVariables.go
```

### Example: run a folder that contains a package

```bash
cd "Functions"
go run .
```

### Example: format Go files

```bash
gofmt -w basicVariables.go
```

### Example: tidy module dependencies

```bash
go mod tidy
```

## REST API project

The [REST API PROJECT](REST%20API%20PROJECT) is the most complete sample in this workspace. It demonstrates a simple event management API using:

- Go
- Gin web framework
- SQLite database
- JSON request/response handling
- CRUD routes for events

### Project layout

- [REST API PROJECT/main.go](REST%20API%20PROJECT/main.go) — app entry point
- [REST API PROJECT/db/db.go](REST%20API%20PROJECT/db/db.go) — database initialization and table creation
- [REST API PROJECT/models/event.go](REST%20API%20PROJECT/models/event.go) — event model and database operations
- [REST API PROJECT/routes/routes.go](REST%20API%20PROJECT/routes/routes.go) — route registration
- [REST API PROJECT/routes/events.go](REST%20API%20PROJECT/routes/events.go) — request handlers
- [REST API PROJECT/api-test](REST%20API%20PROJECT/api-test) — HTTP request examples for testing the API

### Run the API

From the project folder:

```bash
cd "REST API PROJECT"
go run .
```

The server runs on:

```text
http://localhost:8080
```

### Available routes

| Method | Route | Description |
| --- | --- | --- |
| GET | /events | Retrieve all events |
| GET | /events/:id | Retrieve one event |
| POST | /events | Create an event |
| PUT | /events/:id | Update an event |
| DELETE | /events/:id | Delete an event |

### Example payload

```json
{
  "name": "Test Event",
  "description": "A test event",
  "location": "A test location",
  "dateTime": "2025-01-01T15:30:00.000Z"
}
```

You can test the API using the included HTTP examples in [REST API PROJECT/api-test](REST%20API%20PROJECT/api-test) or with tools like Postman or VS Code REST Client.

## Popular commands

```bash
# run a single file
cd "Learn to run Go"
go run textio.go

# run a module
cd "REST API PROJECT"
go run .

# install dependencies
go mod tidy

# check for compile issues
go build ./...

# format code
gofmt -w .
```

## Tips for learning effectively

- Start with the simplest folders first.
- Run each example on its own before moving to the next concept.
- Read the code slowly and experiment by editing small values.
- Re-run examples after changing inputs to understand what the code is doing.
- Use the REST API project as a capstone exercise after the fundamentals.

## Recommended next steps

After you finish the beginner lessons, consider exploring:

- database integration with SQLite or PostgreSQL,
- authentication and authorization,
- testing with Go's built-in test framework,
- concurrency patterns and worker pools,
- deployment and containerization.

## Notes

This repository is a collection of educational examples and mini-projects. Some folders are intentionally simple and meant to demonstrate a single concept at a time, while the REST API project shows how those concepts combine in a real application.

1. Read the lesson and predict what the program will print.
2. Run the program and compare the result with your prediction.
3. Change one value or statement and run it again.
4. Write a small variation without copying the original solution.
5. Record questions and review the related roadmap section.

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

This project is licensed under the [MIT License](LICENSE).

Copyright (c) 2026 Hanfried Nguegan
