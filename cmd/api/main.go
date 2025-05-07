package main

import (
	cfg "go-clean-arch/config"
	"log"
)

func main() {

	cfg.LoadEnvs()

	app := &application{
		config: config{
			addr: ":8080",
		},
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
