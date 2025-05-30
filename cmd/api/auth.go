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

	ctx := r.Context()

	response, err := app.userUseCase.SignIn(ctx, userSignIn)

	if err != nil {
		switch err.Error() {
		case "User already registered":
			writeJsonError(w, http.StatusConflict, "User already registered")
		case "Invalid credentials":
			writeJsonError(w, http.StatusBadRequest, "Password or email incorrect")
		case "Error processing password":
			writeJsonError(w, http.StatusInternalServerError, "Internal server error")
		default:
			writeJsonError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	if err := writeJson(w, http.StatusCreated, response); err != nil {
		writeJsonError(w, http.StatusInternalServerError, "Error writing response")
	}
}

func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {
	var userLogin user.UserLogInRequest

	if err := readJson(w, r, &userLogin); err != nil {
		writeJsonError(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()

	response, err := app.userUseCase.LogIn(ctx, userLogin)

	if err != nil {
		switch err.Error() {
		case "User not found":
			writeJsonError(w, http.StatusNotFound, "User not found")
		case "Invalid credentials":
			writeJsonError(w, http.StatusUnauthorized, "Invalid credentials")
		default:
			writeJsonError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	if err := writeJson(w, http.StatusOK, response); err != nil {
		writeJsonError(w, http.StatusInternalServerError, "Error writing response")
	}
}
