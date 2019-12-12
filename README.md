## Sociozone (SCZ)
a mini instagram. This project is built using Go and PostgreSQL

## Creating Migrations
`migrate create -ext sql -dir src/database/postgres/migrations <migration_name>`

## Run migration UP
- `go run src/migration_runners/migrate_up.go`

## Run migration DOWN
- `go run src/migration_runners/migrate_down.go`