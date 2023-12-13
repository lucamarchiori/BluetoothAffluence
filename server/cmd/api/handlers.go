package main

import (
	"net/http"
)

func (app *application) store(w http.ResponseWriter, r *http.Request) {
	type Device struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Alias   string `json:"alias"`
		TxPower int16  `json:"txPower"`
		RSSI    int16  `json:"rssi"`
	}

	type Scanner struct {
		Address string `json:"address"`
		Name    string `json:"name"`
		Alias   string `json:"alias"`
	}

	type Scan struct {
		Devices []Device `json:"devices"`
		Scanner Scanner  `json:"scanner"`
	}

	s := Scan{}

	err := app.readJSON(w, r, &s)
	if err != nil {
		app.logger.Error(err)
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	// Dump the contents of the input struct in a HTTP response.
	app.logger.Info(s)

	// Store the request content into the database
}
