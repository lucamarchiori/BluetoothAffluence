package main

import (
	"testing"
)

func TestPostScanResults(t *testing.T) {
	scan := Scanner{Address: "aaa", Name: "ScannerName", Alias: "ScannerAlias"}
	devs := []Device{{Address: "aa"}, {Address: "bb"}, {Address: "cc"}}

	s := Scan{Scanner: scan, Devices: devs}

	b := postScanResults(s)
	if b != nil {
		log.Info("Error")
	}
}
