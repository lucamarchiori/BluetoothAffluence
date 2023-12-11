package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func routes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/store", store)
	err := http.ListenAndServe(":4000", r)
	log.Fatal(err)
	log.Info("Webserver started")
	return r
}
