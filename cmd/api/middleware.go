package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

func (app *application) AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				app.unauthorizedErrorResponse(w, r, fmt.Errorf("Auth header missing"))
				return
			}

			headerParts := strings.Split(authHeader, " ")

			if len(headerParts) != 2 || headerParts[0] != "Bearer" {
				app.unauthorizedErrorResponse(w, r, fmt.Errorf("Invalid auth header format"))
				return
			}

			token := headerParts[1]

			if token == "" {
				app.unauthorizedErrorResponse(w, r, fmt.Errorf("Token is empty"))
				return
			}

			ctx := context.WithValue(r.Context(), "token", token)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
