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