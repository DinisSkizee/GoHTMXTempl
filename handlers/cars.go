package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DinisSkizee/GoHTMXTempl/store"
	"github.com/gorilla/mux"
)

func (h *Handler) HandleListCars(w http.ResponseWriter, r *http.Request) {
	cars := h.store.ListCars()
	json.NewEncoder(w).Encode(cars)
}

func (h *Handler) HandleAddCar(w http.ResponseWriter, r *http.Request) {
	var car store.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.store.AddCar(car)
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) HandleDeleteCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	h.store.DeleteCar(id)
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) HandleSearchCar(w http.ResponseWriter, r *http.Request) {
	make := r.URL.Query().Get("make")
	cars := h.store.SearchCar(make)
	json.NewEncoder(w).Encode(cars)
}
