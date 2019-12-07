## Creating a Migration
- to create migration run `migrate create -ext sql -dir src/database/postgres/migrations/ migration_name`

## Run migration UP using golang-migrate CLI
- `migrate -path src/database/postgres/migrations/ -database 'psql://mac@tcp(localhost:5432)/scz_db?sslmode=disable' up`

## Run migration DOWN using golang-migrate CLI
- `migrate -path src/database/postgres/migrations/ -database 'psql://mac@tcp(localhost:5432)/scz_db?sslmode=disable' down`
