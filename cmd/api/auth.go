package main

import (
	"go-clean-arch/internal/domain/usecase/user"
	"net/http"
)

// signIn godoc
//
//	@Summary		User Sign In
//	@Description	Register a new user in the app
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		user.UserSignInRequest	true	"User sign in data"
//	@Success		201		{object}	user.SignInResponse		"User successfully created"
//	@Failure		400		{object}	map[string]string		"Bad request - invalid input"
//	@Failure		409		{object}	map[string]string		"Conflict - user already exists"
//	@Failure		500		{object}	map[string]string		"Internal server error"
//	@Router			/auth/sign-in [post]
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

// login godoc
//
//	@Summary		User Login
//	@Description	Login with existing user credentials
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		user.UserLogInRequest	true	"User login data"
//	@Success		200		{object}	user.UserLogInResponse	"User successfully logged in"
//	@Failure		400		{object}	map[string]string		"Bad request - invalid input"
//	@Failure		401		{object}	map[string]string		"Unauthorized - invalid credentials"
//	@Failure		404		{object}	map[string]string		"Not found - user does not exist"
//	@Failure		500		{object}	map[string]string		"Internal server error"
//	@Router			/auth/login [post]
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
