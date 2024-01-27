package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/store", app.store)
	router.HandleFunc("/scanner/count-scan-devices", app.countScanDevices)
	router.HandleFunc("/scanner/index", app.indexScanner)

	return router
}
