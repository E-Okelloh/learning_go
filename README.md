# Learn Go by Building a REST API

A step-by-step project for learning Go while building a small REST API. You will start with the language basics, then add HTTP routes, JSON, validation, tests, persistence, and production-minded improvements.

## Project: Task Tracker API

Build an API that lets clients create and manage tasks. Keep the first version intentionally small so you can focus on learning Go and the structure of a web service.

### Goals

- Learn Go syntax, packages, structs, methods, errors, and interfaces.
- Build HTTP handlers with Go's standard library.
- Work with JSON requests and responses.
- Implement CRUD endpoints and useful HTTP status codes.
- Write tests for the API.
- Add persistent storage after the in-memory version works.

## Prerequisites

- Go installed (use a current stable release): <https://go.dev/dl/>
- A terminal and a code editor
- Basic familiarity with HTTP and JSON is helpful, but not required

Check your Go installation:

```sh
go version
```

## Step-by-step learning path

Complete each step in order. Run the tests and commit your work before moving on.

### 1. Create the Go module

From the project directory, initialize the module:

```sh
go mod init learning-go-rest-api
go mod tidy
```

Create `main.go` and print a short startup message. Run it with:

```sh
go run .
```

**Learn:** Go modules, `package main`, imports, and the `main` function.

### 2. Define the task model

Create a `Task` struct with fields such as:

- `ID` — a unique integer
- `Title` — required text
- `Completed` — whether the task is done

Add JSON struct tags so the API uses lowercase JSON names (for example, `json:"title"`). Create a few task values and print them.

**Learn:** structs, field names, composite literals, and struct tags.

### 3. Start an HTTP server

Use the standard library's `net/http` package. Start a server on `localhost:8080` and add a simple `GET /health` handler that returns a success response.

Try it in another terminal:

```sh
curl -i http://localhost:8080/health
```

**Learn:** handlers, `http.ResponseWriter`, `*http.Request`, and server startup.

### 4. Return JSON

Set the response `Content-Type` to `application/json` and use `encoding/json` to encode a task or a small response object. Handle encoding errors rather than ignoring them.

**Learn:** JSON encoding, response headers, and Go error handling.

### 5. Add in-memory task storage

Store tasks in a map or slice and keep the next ID in a variable. Put the storage and its operations in a separate package or type instead of putting all application logic in `main`.

Add a `GET /tasks` route that returns all tasks. At this stage, data will reset whenever the server restarts; that is expected.

**Learn:** maps, slices, functions, methods, and separating responsibilities.

### 6. Implement CRUD endpoints

Use `http.ServeMux` and, if your Go version supports it, method-aware patterns. Otherwise, check `r.Method` in each handler and return `405 Method Not Allowed` for unsupported methods.

| Method | Path | Purpose | Typical success status |
| --- | --- | --- | --- |
| `GET` | `/health` | Check that the server is running | `200 OK` |
| `GET` | `/tasks` | List tasks | `200 OK` |
| `GET` | `/tasks/{id}` | Get one task | `200 OK` |
| `POST` | `/tasks` | Create a task | `201 Created` |
| `PUT` | `/tasks/{id}` | Replace/update a task | `200 OK` |
| `DELETE` | `/tasks/{id}` | Delete a task | `204 No Content` |

For `POST` and `PUT`, decode the request body with `json.Decoder`. Validate input, reject malformed JSON with `400 Bad Request`, and return `404 Not Found` when a task ID does not exist. Do not trust an ID supplied in the request body when creating a task; assign it on the server.

Example requests:

```sh
curl -i http://localhost:8080/tasks

curl -i -X POST http://localhost:8080/tasks \
  -H 'Content-Type: application/json' \
  -d '{"title":"Learn Go","completed":false}'

curl -i http://localhost:8080/tasks/1

curl -i -X PUT http://localhost:8080/tasks/1 \
  -H 'Content-Type: application/json' \
  -d '{"title":"Practice Go","completed":true}'

curl -i -X DELETE http://localhost:8080/tasks/1
```

**Learn:** routing, path values, request decoding, validation, and HTTP semantics.

### 7. Make errors consistent

Return JSON errors in one predictable format, such as `{"error":"task not found"}`. Set the status code before writing the response body. Avoid exposing internal error details to API clients.

**Learn:** reusable helpers, error values, and API design.

### 8. Write tests

Add tests for your handlers and storage. Use `net/http/httptest` to create requests and record responses without starting a real server.

Test at least:

- Health check returns success.
- Listing tasks returns valid JSON.
- Creating a task returns `201` and the created task.
- Invalid JSON returns `400`.
- Unknown task IDs return `404`.
- Unsupported methods return `405`.

Run the test suite:

```sh
go test ./...
go test -race ./...
go vet ./...
```

**Learn:** table-driven tests, test helpers, `httptest`, and basic race detection.

### 9. Improve the server

Use an `http.Server` with sensible read, write, and idle timeouts. Add graceful shutdown so the server can stop cleanly when it receives an interrupt signal. Keep configuration such as the listening address outside handler logic.

**Learn:** context cancellation, signals, configuration, and server lifecycle.

### 10. Add persistent storage

Once the in-memory API and tests are solid, replace the in-memory store with a database. SQLite is a manageable next step for a learning project; PostgreSQL is another option. Keep the HTTP handlers independent of the database by defining a small storage interface and implementing it for each store.

Add schema setup/migrations and tests for persistence. Keep database credentials and other secrets in environment variables; never commit secrets.

**Learn:** interfaces, dependency injection, SQL, migrations, and configuration.

### 11. Document and polish

- Document how to install, run, test, and configure the service.
- Add request/response examples and explain error responses.
- Format and check the code:

```sh
gofmt -w .
go test ./...
go vet ./...
```

- Consider adding pagination, filtering, structured logging, and API versioning as follow-up exercises.

## Suggested project structure

Start with only a few files and introduce packages when there is a clear reason to do so. One possible structure after the project grows:

```text
.
├── go.mod
├── main.go
├── internal/
│   ├── task/
│   │   ├── model.go
│   │   └── store.go
│   └── httpapi/
│       ├── handlers.go
│       └── handlers_test.go
└── README.md
```

## Completion checklist

- [ ] The service starts locally and `/health` responds.
- [ ] Task CRUD endpoints work with JSON.
- [ ] Invalid input and missing tasks return appropriate status codes.
- [ ] Handlers and storage have automated tests.
- [ ] `go test ./...` and `go vet ./...` pass.
- [ ] Tasks persist after restarting the service (after completing the database step).
- [ ] The README explains how to run and use the API.

## Further learning

- [A Tour of Go](https://go.dev/tour/)
- [How to Write Go Code](https://go.dev/doc/code)
- [Go standard library documentation](https://pkg.go.dev/std)
- [`net/http` documentation](https://pkg.go.dev/net/http)
