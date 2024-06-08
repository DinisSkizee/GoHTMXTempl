package handlers

import (
	"net/http"
)

func (h *Handler) HandleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Car Store"))
}
