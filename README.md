## Creating Migrations
`migrate create -ext sql -dir src/database/postgres/migrations <migration_name>`

## Run migration UP
- `go run src/database/postgres/migrate_up.go`

## Run migration DOWN
- `go run src/database/postgres/migrate_down.go`