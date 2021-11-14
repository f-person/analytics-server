package main

import (
	"fmt"
	"net/http"
)

func recoverPanic(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			// Builtin recover func checks whether there has been a panic or not
			if err := recover(); err != nil {
				w.Header().Set("Conntection", "close")
				app.ServerError(w, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	}
}

func setHeaders(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add(
			"Access-Control-Allow-Methods",
			"GET, OPTIONS, POST, PUT, DELETE",
		)
		w.Header().Add("Access-Control-Allow-Headers", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func logRequest(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.InfoLog.Printf(
			"%s - %s %s %s",
			r.RemoteAddr,
			r.Proto,
			r.Method,
			r.URL,
		)

		next.ServeHTTP(w, r)
	}
}
