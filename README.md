## Creating Migrations
`migrate create -ext sql -dir ./migrations -seq <migration_name>`

## Running Migrations
- example
`migrate -source file://path/to/migrations -database postgres://localhost:5432/database up 2`
