# Blog Aggregator

## Useful Commands

Create new migration with goose
```console
goose postgres CONN up --dir sql/schema
```

Update Schema
```console
sqlc generate
```

Run server
```console
go build -o server.exe && ./server
```