package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func routes() http.HandlerFunc {
	router := pat.New()

	return recoverPanic(setHeaders(logRequest(router)))
}
