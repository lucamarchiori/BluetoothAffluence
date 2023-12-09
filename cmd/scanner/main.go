package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/muka/go-bluetooth/bluez/profile/adapter"
)

var log = logrus.New()

func main() {
	log.Out = os.Stdout

	log.Infof("Bluetooth affluence script started")

	aid := adapter.GetDefaultAdapterID()
	scan, err := Run(aid, 20)

	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	res := postScanResults(scan)
	if !res {
		log.Error("Error sending data to server")
		os.Exit(1)
	}

	os.Exit(0)
}

func postScanResults(scan Scan) bool {
	var url string = "https://webhook.site/0435b241-6cd6-4707-a286-35da5feda75a"
	var d []byte
	var err error

	d, err = json.Marshal(scan)

	if err != nil {
		log.Error(err)
		return false
	}

	bd := []byte(d)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bd))

	if err != nil {
		log.Error(err)
		return false
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Error(err)
		return false
	}
	defer resp.Body.Close()

	return true
}
