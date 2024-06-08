package handlers

import (
	"github.com/sikozonpc/fullstackgo/store"
)

type Handler struct {
	store *store.InMemoryStore
}

func New(store *store.InMemoryStore) *Handler {
	return &Handler{store: store}
}
