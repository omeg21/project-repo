package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// Nosurf adds CSRF to all POST request
func Nosurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

//SessionLoad loads and save the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)

}

// func WriteToConsole(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("hit the button")
// 		next.ServeHTTP(w, r)
// 	})
// }
