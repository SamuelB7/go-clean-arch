package main

import (
	"net/http"
	"os"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	environment := os.Getenv("ENVIRONMENT")
	version := os.Getenv("VERSION")

	data := map[string]string{
		"status":      "ok",
		"environment": environment,
		"version":     version,
	}

	if err := writeJson(w, http.StatusOK, data); err != nil {
		writeJsonError(w, http.StatusInternalServerError, err.Error())
	}
}
