package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	// Setting the logger to a file
	log.Out = os.Stdout
	// file, err := os.OpenFile("./btlog.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err == nil {
	// 	log.Out = file
	// } else {
	// 	log.Info("Failed to log to file, using default stderr")
	// }
}

func main() {
	log.Infof("Bluetooth affluence script started")

	//aid := adapter.GetDefaultAdapterID()
	//scan, err := Run(aid, 20)
	scan, err := RunMock()

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
	var url string = "http://127.0.0.1:4000/store"
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

	log.Info(resp.Body)

	if err != nil {
		log.Error(err)
		return false
	}
	defer resp.Body.Close()

	return true
}