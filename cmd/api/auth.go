package main

import (
	"go-clean-arch/internal/domain/usecase/user"
	"net/http"
)

func (app *application) signInHandler(w http.ResponseWriter, r *http.Request) {
	var userSignIn user.UserSignInRequest

	if err := readJson(w, r, &userSignIn); err != nil {
		writeJsonError(w, http.StatusBadRequest, err.Error())
		return
	}
}
