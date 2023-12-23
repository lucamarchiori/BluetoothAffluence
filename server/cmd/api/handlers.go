package main

import (
	"net/http"

	"lucamarchiori.bluetoothAffluence/server/internal/data"
)

func (app *application) store(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Devices []struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Alias   string `json:"alias"`
			TxPower int16  `json:"txPower"`
			RSSI    int16  `json:"rssi"`
		}
		Scanner struct {
			Address string `json:"address"`
			Name    string `json:"name"`
			Alias   string `json:"alias"`
		}
		ScanTime string `json:"scanTime"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.logger.Error(err)
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
	}

	// Dump the contents of the input struct
	app.logger.Info(input)

	// Store SCANNER info
	sc := &data.Scanner{
		Address: input.Scanner.Address,
		Name:    input.Scanner.Name,
		Alias:   input.Scanner.Alias,
	}

	res, err := app.models.Scanner.Insert(sc)
	if err != nil {
		app.logger.Error(err)
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
	}

	// Get the scanner DB Obj
	sc, err = app.models.Scanner.GetByAddress(sc.Address)

	if err != nil {
		app.logger.Error(err)
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
	}

	if sc == nil {
		app.logger.Error("Scanner not found")
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
	}

	// Handle SCAN info
	s := &data.Scan{
		ScanTime:  input.ScanTime,
		ScannerId: (*sc).Id,
	}

	res, err = app.models.Scan.Insert(s)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	scan_id, _ := res.LastInsertId()

	// Store the request content into the database
	for _, device := range input.Devices {
		d := &data.Device{
			Name:    device.Name,
			Alias:   device.Alias,
			RSSI:    device.RSSI,
			TxPower: device.TxPower,
			Address: device.Address,
			ScanId:  int16(scan_id),
		}

		_, err := app.models.Device.Insert(d)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}
	}

}
