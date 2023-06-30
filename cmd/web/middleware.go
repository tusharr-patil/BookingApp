package main

import (
	"net/http"

	noSurf "github.com/justinas/nosurf"
)

func CsrfHandler(next http.Handler) http.Handler {
	csrfHandler := noSurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
