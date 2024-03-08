# Architecture

In this project structure, the cmd directory contains the entry point of the application, which is the main.go file. The internal directory contains the implementation details of the application, which are divided into three layers: 
- api: The api layer contains the input/output logic of the application
- domain: The domain layer contains the business logic and entities of the application
- infrastructure: The infrastructure layer contains the implementation details

# Tools
 * Make build tool.
 * [Gin](https://github.com/gin-gonic/gin) HTTP Web Framework.
 * [pgx](https://github.com/jackc/pgx) for PostgreSQL Driver.
 * [Migrate](https://github.com/golang-migrate/migrate) Database migrations CLI.
 * [Testify](https://github.com/stretchr/testify) Test framework.
 * [GoMock](https://github.com/uber-go/mock) Mocking framework.
 * [Swaggo](https://github.com/swaggo/swag) for generate RESTful API documentation with Swagger 2.0.

# Run application
### Run ParkWise Server
First configure environment variable on your local machine or update it on `Makefile` and `app.env`, then do db migration for initial db setup

```bash
make migrate-up
```
It will setup db. Then run server

```bash
make run
```
After running the application on localhost visit http://localhost:8080/swagger/index.html#/ for swagger doc of the project, or get the raw swagger json and yaml from `docs` directory

### Test
```bash
make test
```
Please check `internal/infrastructure/datastore/pg/parkingLot_test.go` for unit test, and `internal/api/web/parkingLot_test.go` for HTTP API test. Test will be performed automatically on Github Runner, where it will also setup postgresql db test.

Note: As it is a assessment, I am doing here only one test.

## Other Instructions:
- Dockerfile is available on `cmd/server/Dockerfile`
- All API Endpoints, Request Body, Query Params, and Response format are available on `docs` directory and also available on [swagger endpoint](http://localhost:8080/swagger/index.html#/)
- All the PostgreSQL schema is available on `db` directory

# SQL Schema
### Migrations
run (if not installed)
```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```
To create a new migration
```bash
migrate create -ext sql -dir database/migrations -seq name_for_migration
```

and then edit the new files in `/database/migrations`.

```bash
make migrate-up
```
for db tear-down
```bash
make migrate-down
```

### Schema Dump (if use docker)
```bash
docker exec pg-docker pg_dump -U postgres --schema-only --verbose -s parkwise > ts_ddl.sql
```

# GoMock: for mocking definitions
Install gomock
```bash
go install go.uber.org/mock/mockgen@latest
```
Generate mock datastore definition implementation
```bash
make mock-db
```
# API(Swagger doc generation)
```bash
make swag
```

# Deployment
### Docker build
```bash
docker build --progress=plain -f ./cmd/server/Dockerfile -t parkwise:latest -t parkwise:0.1 .
```

