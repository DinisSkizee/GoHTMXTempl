package handlers

import (
	"net/http"

	"github.com/DinisSkizee/GoHTMXTempl/views"
)

func (h *Handler) HandleHome(w http.ResponseWriter, r *http.Request) {
	views.Home().Render(r.Context(), w)
}
