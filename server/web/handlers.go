package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func store(w http.ResponseWriter, r *http.Request) {
	log.Info(r.Body)
	log.Info(r.Form)
	log.Info(r.Form.Get("scan"))

	var Scan struct {
		Devices []struct {
			Address string `json:"Address"`
			Name    string `json:"Name"`
			Alias   string `json:"Alias"`
		}
		Scanner struct {
			Address string
			Name    string
			Alias   string
		}
	}

	err := json.NewDecoder(r.Body).Decode(&Scan)
	if err != nil {
		log.Error(err)
		return
	}
	// Dump the contents of the input struct in a HTTP response.
	fmt.Fprintf(w, "%+v\n", Scan)

	// Store the request content into the database
}
