package main

import (
	"net/http"

	"lucamarchiori.bluetoothAffluence/server/internal/data"
)

/*
Handles the HTTP request for storing Bluetooth scan information.
*/
func (app *application) store(w http.ResponseWriter, r *http.Request) {

	// Define a struct for parsing JSON input
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

	// Parse JSON input from the request body
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

	// Insert scanner information into the database
	res, err := app.models.Scanner.Insert(sc)
	if err != nil {
		app.logger.Error(err)
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
	}

	// Retrieve the scanner object from the database
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

	// Insert scan information into the database
	res, err = app.models.Scan.Insert(s)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// Retrieve the scan ID from the database
	scan_id, _ := res.LastInsertId()

	// Store information about each discovered device into the database
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

	// Respond with success
	//app.logger.Info("Bluetooth scan data stored successfully.")
	//w.WriteHeader(http.StatusOK)

}

/*
Handle requests to retrieve the number of scanned devices
*/
func (app *application) countScanDevices(w http.ResponseWriter, r *http.Request) {
	// Set CORS header
	w.Header().Set("Access-Control-Allow-Origin", "*")

	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	startDate = "2023-12-12 00:00:00" //TODO: HARDCODED DATES FOR DEMO PURPOSES
	endDate = "2023-12-12 23:59:59"

	if startDate == "" || endDate == "" {
		app.errorResponse(w, r, http.StatusBadRequest, "Invalid date range")
		return
	}

	// Get the scanner id from the request
	//scannerId, err := strconv.Atoi(r.URL.Query().Get("scanner_id")) //TODO: HARDCODED SCANNER ID FOR DEMO PURPOSES
	// if err != nil {
	// 	app.logger.Error(err)
	// 	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
	// }

	scannerId := 1

	if scannerId == 0 {
		app.errorResponse(w, r, http.StatusBadRequest, "Invalid scanner id")
		return
	}

	// --- Moving AVG ---
	res, err := app.models.Scan.CountScanDevicesMovingAVG(scannerId, startDate, endDate)

	if err != nil {
		app.logger.Error(err)
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
	}

	movAvgData := map[string]interface{}{"count": res, "scanner_id": scannerId, "start_date": startDate, "end_date": endDate}

	// --- Count ---
	res, err = app.models.Scan.CountScanDevices(scannerId, startDate, endDate)

	if err != nil {
		app.logger.Error(err)
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
	}

	countData := map[string]interface{}{"count": res, "scanner_id": scannerId, "start_date": startDate, "end_date": endDate}

	// --- Count Time AVG ---

	startDate = "2023-12-01 00:00:00" //TODO: HARDCODED DATES FOR DEMO PURPOSES
	endDate = "2023-12-31 23:59:59"
	resAvg, err := app.models.Scan.CountScanDevicesTimeAvg(scannerId, startDate, endDate)

	if err != nil {
		app.logger.Error(err)
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
	}

	timeAvg := map[string]interface{}{"count": resAvg, "scanner_id": scannerId, "start_date": startDate, "end_date": endDate}

	rp := customResp{Message: "Scanners retrived", Data: map[string]interface{}{"count": countData, "countMovingAvg": movAvgData, "timeAvg": timeAvg}, Status: 200}

	err = app.writeJSON(w, http.StatusOK, rp, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

/*
IndexScanner handles the HTTP request for retrieving a list of scanners.
*/
func (app *application) indexScanner(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Call the Index method on the Scanner model to retrieve the list of scanners
	res, err := app.models.Scanner.Index()

	// Handle any errors that occurred during the Index method call
	if err != nil {
		app.logger.Error(err)
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
	}

	// Prepare a custom response with the retrieved scanners
	rp := customResp{Message: "Scanners retrived", Data: map[string]interface{}{"scanners": res}, Status: 200}

	// Write the JSON response to the client
	err = app.writeJSON(w, http.StatusOK, rp, nil)
	if err != nil {
		// Handle any errors that occurred during writing the JSON response
		app.serverErrorResponse(w, r, err)
	}
}
