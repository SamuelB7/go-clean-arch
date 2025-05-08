package main

import (
	cfg "go-clean-arch/config"
	store "go-clean-arch/internal/store/postgres"
	"log"
)

func main() {

	cfg.LoadEnvs()

	store := store.NewPostgresStorage(nil)

	app := &application{
		config: config{
			addr: ":8080",
		},
		store: store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
