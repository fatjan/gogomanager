# gogomanager
Manager manages users

### Run The App
1. Copy the .env.example file and rename it to .env.
2. Open the .env file and fill in the required credentials (e.g., database details).
3. Run the application using the following command:
```
go run cmd/gogomanager/main.go
```

### Run Migrations
1. Installing goose:
```
go install github.com/pressly/goose/v3/cmd/goose@latest
```
2. Going to the folder schema: `internal/database/schema`
3. Run the command
```
goose postgres "postgres://user:password@localhost:5432/dbname" up
```

### Linter
```
golangci-lint run
```
or
```
golangci-lint run path/to/your/file.go
```
