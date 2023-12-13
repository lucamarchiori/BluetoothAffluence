package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/store", app.store)

	return router
}
