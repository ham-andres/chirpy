# CHIRPY PROJECT: HTTP Server fundamentals

## requirements: 
  - install go
  

## config:
  - <go mod init> command to creat go.mod

## CH1 & CH2: Server & Routing:
This module covers the core components of building web servers using Go's `net/http` standard library.

## Key Concepts

- **ServeMux**: The standard library multiplexer used to match incoming request URLs to specific handlers.
- **Handlers & HandlerFuncs**: Implementing the `http.Handler` interface by defining `ServeHTTP`, or using `http.HandlerFunc` for functional logic.
- **Request vs. Response**:
    - `*http.Request`: Used to inspect incoming data (Method, Path, Headers).
    - `http.ResponseWriter`: An interface used to stream the response body and status codes back to the client.

## Core Pattern

1. Initialize a router (`NewServeMux`).
2. Register routes using handlers.
3. Start the server on a specific port using `ListenAndServe`.
 - Installation/Setup Logic:
  1. initialize a new http.ServeMux
  2. register handlers using mux.HandleFunc(pattern, handler)
  3. wrap mux with middleware functions if needed.
  4. pas the mux to http.ListenAndServe.

- In chapter 2, we transitioned from basic file serving to structed request handling.
We covered - middleware, (intercepting requests to add functionality like logging or metrics), Stateful handler( using structs to maintain application state like hit counters) across multiple http requests.
Advanced routing - implementing http.NewServeMux to dispatch requests based on specific paths.
Pattern matching - Understanding how Go selects the "longest match " and handles subtree paths versus fixed paths.

## CH3: Architecture: 
In this phase of the project, we focused on transitioning from a basic server to a structured Monolith Architecture, and also Decoupled deployment.

- Monolith Structure: Combined the file server (front-end assets) and API handlers into a single Go binary for simplicity.

- Namespacing & Routing: implemented an /admin namespace to separate internal metrics and manangement tools from public /api endpoints.

- stateful handlers: Leveraged custom middleware and shared structs to track server wide metrics, such as hits, without global variables.

## CH4: JSON Handling:
Implemented an HTTP server capable of parsing and serving JSON data. 
Used specialized HTTP clients like cURL and Postman to test POST requests and custom headers that standard browsers cannot easily handle.

Example test uses: 
````
curl -X POST http://localhost:8080/api/validate_chirp \
  -H "Content-Type: application/json" \
  -d '{"body": "This is a kerfuffle opinion"}'
````


## CH5: Storage: 
 - Set up SQLC to generate type-safe Go code from raw SQL queries.
 - Added sqlc.yaml config pointing to schema files, query files and generated output files.
 - Added reuired dependencies: PostgreSQL driver, UUID package and godotenv.
 - Stored the database connection string in a .env file and loaded it at app startup.
 - Opened a Postgres connection with database/sql.
 - Created a shared SQLC queries object and attached it to app config for handler access.



