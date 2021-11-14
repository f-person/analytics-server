package main

import "net/http"

// Creates an event
// @description Creates an event.
// @param event body domain.Event true "event"
// @success 200 {string} string "ok"
// @Router /events [post]
func CreateEvent(w http.ResponseWriter, r *http.Request) {

}
