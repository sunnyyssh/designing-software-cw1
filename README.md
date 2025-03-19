# Desiging software. Control work 1

# Usage
## Prepare
Build CLI
```shell
go build -o bankcli ./cmd/main.go
```
Then run PostgreSQL
```shell
docker run -d \
  --name postgres-bank \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=secret \
  -e POSTGRES_DB=bank \
  -v ./migrations:/migrations:ro \
  -p 5432:5432 \
  postgres:latest
```
I don't wanna spam services and tools, so apply migration youself please:
```shell
export PGPASSWORD=secret
docker exec -it postgres-bank psql -d bank -U postgres -f /migrations/V001__init.sql
```
And export connection string, so app can connect DB
```shell
export PG_CONN_STRING="postgresql://postgres:secret@localhost:5432/bank?sslmode=disable"
```
## Use, enjoy
Watch help
```shell
./bankcli help`
```