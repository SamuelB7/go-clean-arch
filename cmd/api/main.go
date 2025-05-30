package main

import (
	"fmt"
	cfg "go-clean-arch/config"
	repository "go-clean-arch/internal/adapter/repository/postgresql"
	"go-clean-arch/internal/domain/usecase/user"
	"log"
	"os"
)

func main() {

	cfg.LoadEnvs()

	db_host := os.Getenv("DATABASE_HOST")
	db_port := os.Getenv("DATABASE_PORT")
	db_user := os.Getenv("DATABASE_USER")
	db_password := os.Getenv("DATABASE_PASSWORD")
	db_name := os.Getenv("DATABASE_NAME")

	db, err := repository.NewPGSqlConnection(db_host, db_port, db_user, db_password, db_name, 30, 30)

	if err != nil {
		log.Fatal(fmt.Errorf("unable to open database connection: %w", err))
	}

	repo := repository.NewPostgresRepository(db)

	userUseCase := user.NewUserService(repo.Users())

	app := &application{
		config: config{
			addr: ":8080",
		},
		repository:  repo,
		userUseCase: userUseCase,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
