package main

import "net/http"

func (app *application) unauthorizedErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	writeJsonError(w, http.StatusUnauthorized, "Unauthorized")
}
