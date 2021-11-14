package main

import (
	"flag"
	"net/http"
)

var app App

// @title Analytics API
// @version 1.0
// @BasePath /api/v1
// @accept json
// @produce json
// @schemes http https
// @x-extension-openapi {"example": "value on a json format"}
func main() {
	addr := flag.String(":addr", ":8000", "HTTP Network Address")
	flag.Parse()

	app = NewApp()

	server := &http.Server{
		Addr:    *addr,
		Handler: routes(),
	}

	app.InfoLog.Printf("Starting server on %s.", *addr)
	err := server.ListenAndServe()
	app.ErrorLog.Fatal(err)
}
