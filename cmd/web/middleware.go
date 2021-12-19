package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

//  WriteToConsole complete

func WriteToConsole(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("listian")

		next.ServeHTTP(w, r)
	})
}

// NoSurf adds CSRF protections to all POST request
func NoSurf(next http.Handler) http.Handler {

	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		// Secure:   false,
		Secure: app.InProduction,

		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler

}

// SessionLoad loads and saves the sesion on very request

func SessionLoad(next http.Handler) http.Handler {

	return session.LoadAndSave(next)
}
