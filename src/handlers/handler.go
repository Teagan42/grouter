package handlers

import "net/http"

type handler struct {
	handle func(w http.ResponseWriter, r http.Request) bool
	next *handler
}