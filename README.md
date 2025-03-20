# Desiging software. Control work 1

## Prepare usage
Build CLI
```shell
go build -o bankcli ./cmd/main.go
```
Then run PostgreSQL
```shell
docker run -d \
  --name postgres-bank \
  -e POSTGRES_USER=bank \
  -e POSTGRES_PASSWORD=secret \
  -e POSTGRES_DB=bank \
  -v ./migrations:/migrations:ro \
  -p 5433:5432 \
  postgres:latest
```
I don't wanna spam services and tools, so apply migration youself please:
```shell
export PGPASSWORD=secret
docker exec -it postgres-bank psql -d bank -U bank -f /migrations/V001__init.sql
```
And export connection string, so app can connect DB
```shell
export PG_CONN_STRING="postgresql://bank:secret@localhost:5433/bank?sslmode=disable"
```
## Use, enjoy
Watch help
```shell
./bankcli help
```

# Used Patterns
1. Repository pattern \
I implemented `BankAccountRepo`, `CategoryRepo`, and `OperationRepo` interfaces to abstract database operations.
2. Service Layer Pattern \
I created `BankAccountService`, `CategoryService`, and `OperationService` to handle business logic and coordinate interactions with the repositories.
3. Dependency Injection \
I passed dependencies (e.g., repositories) into services and CLI commands via constructors or function parameters.
4. Command Pattern \
I encapsulated a request as an object, allowing parameterization.
I used the cobra library to define CLI commands (e.g., `create`, `get`, `list`, `delete`).
Each command encapsulates a specific action and its parameters.
5. Data Transfer Object (DTO) Pattern \
I created DTOs like `BankAccountDTO`, `CategoryDTO`, and `OperationDTO` to transfer data between the service layer and the CLI.
6. Factory Method Pattern \
I created `NewCategory()`, `NewBankAccount` functions to implement creating domain objects with some validation inside.
7. Error as value Pattern \
As all Go code I defined errors as `error` values which allow me to handle them appropriately

# SOLID, GRASP, Clean Architecture
I hope, there is no principles that I've violated. Code is structured in Clean Architecture style, dependencies are center-forwarded as it should be. 