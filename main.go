package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DinisSkizee/GoHTMXTempl/handlers"
	"github.com/DinisSkizee/GoHTMXTempl/store"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Initialize in-memory store
	store := store.NewInMemoryStore()
	handler := handlers.New(store)

	router.HandleFunc("/", handler.HandleHome).Methods("GET")
	router.HandleFunc("/cars", handler.HandleListCars).Methods("GET")
	router.HandleFunc("/cars", handler.HandleAddCar).Methods("POST")
	router.HandleFunc("/cars/{id}", handler.HandleDeleteCar).Methods("DELETE")
	router.HandleFunc("/cars/search", handler.HandleSearchCar).Methods("GET")

	// Serve files in public
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	fmt.Printf("Listening on %v\n", "localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
