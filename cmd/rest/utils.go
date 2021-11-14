package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *App) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Output(1, trace)

	http.Error(
		w,
		http.StatusText(http.StatusInternalServerError),
		http.StatusInsufficientStorage,
	)
}
