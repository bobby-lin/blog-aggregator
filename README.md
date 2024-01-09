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

## Observations
There are situation where we want to convert database objects to object with different types for usage in HTTP response.

Example: `sql.NullTime` returns a nested object.
```json
{
  "ID": "627d8a35-adbc-11ee-943d-e86a64565d8d",
  "LastFetchedAt": {
    "Time": "0001-01-01T00:00:00Z",
    "Valid": false
  }
}
```