package main

import (
	"log"
	"net/http"
)

func (app *application) unauthorizedErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	writeJsonError(w, http.StatusUnauthorized, "Unauthorized")
}

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Internal server error: %s path: %s", r.Method, r.URL.Path, err)

	writeJsonError(w, http.StatusInternalServerError, "Internal server error")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Bad request exception: %s path: %s", r.Method, r.URL.Path, err)

	writeJsonError(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Not found exception: %s path: %s", r.Method, r.URL.Path, err)

	writeJsonError(w, http.StatusNotFound, "Not found")
}
