package main

import (
	"log"
	"net/http"

	"github.com/mattdavis0351/day3api/views"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()

	// this adds /site/ on to /api/v1/, so strip the entire thing /api/v1/site
	api.PathPrefix("/site").Handler(http.StripPrefix("/api/v1/site", http.FileServer(http.Dir("assets/public/")))).Methods(http.MethodGet)
	api.HandleFunc("/byte-json", views.ByteRoute).Methods(http.MethodGet)
	api.HandleFunc("/struct-life", views.StructRoute).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}
