## CHIRPY PROJECT: A introduction to HTTP Server fundamentals
High performance REST API built in Go that handles user authentication and blog post management  using a PostgreSQL database.

## INSTALLATION:
1. clone the repo: `git clone <url>`
2. Install dependencies: `go mod download`
3. Run the Server: `go run main.go`

## CONFIG:
- Set up a mod file, `go mod init` command
Run the server: `go run .`
then visit `localhost:8080` in Browser

## USAGE / ENDPOINTS:
`GET /api/healthz` : Returns a 200 OK status to verify the server is healthy.
`POST /api/users`: Allows to create a user profile, and store the connection data in the database .
`POST /api/login`: Allows users to communicate with the server after authentication.
`POST /api/chirps`: Allows authenticated users to create a new chirp.

## Requirements:
- Go 1.22+
- PostgreSQL

## Environment Variables
Create a `.env` file with:
- `DB_URL`: PostgreSQL connection string 
- `JWT_SECRET`: Secret key for authentication.

