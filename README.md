# GO Clean Arch API
Simple golang api being made with clean architecture principles

## Requirements
- Docker compose

## To run the project
 - Create environment file:

 ```sh
 cp .env.example .env
 ```

 - Start the container:

 ```sh
 docker compose up --build
 ```

 The api will start at `localhost:8080/v1`

### Run the migrations
- install the migrate cli
```sh
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

- Execute the migrations to populate local database
```sh
make migrate-up
```

### Swagger Documentation

To generate and update de docs run the command:

```sh
 make gen-docs
```

To access the swagger docs, access: `http://localhost:8080/v1/swagger/index.html`

### Tests

#### Execute all tests
```sh
go test ./...
```

#### Execute specific test
```sh
go test ./cmd/api/
go test ./internal/adapter/repository/mock/
```

#### Execute tests with verbose
```sh
go test -v ./cmd/api/
```