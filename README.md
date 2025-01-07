# gogomanager
Manager manages users

### Run Migrations
1. Installing goose:
```
go install github.com/pressly/goose/v3/cmd/goose@latest
``` 
2. Going to the folder schema.
3. Run the command 
```
goose postgres "postgres://user:password@tcp(localhost:5432)/dbname" up
```
