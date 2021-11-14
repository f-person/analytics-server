package main

import "net/http"

type Event struct {
	Name string
	Data map[string]string `json:"data,omitempty"`
}

// Creates an event
// @description Creates an event.
// @param event body Event true "event"
// @success 200 {string} string "ok"
// @Router /events [post]
func CreateEvent(w http.ResponseWriter, r *http.Request) {

}
