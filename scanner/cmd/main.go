package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/muka/go-bluetooth/bluez/profile/adapter"
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
	var aid string
	var scan Scan
	var err error
	sleepDuration := time.Duration(300) * time.Second
	log.Infof("Bluetooth affluence script started")
	for {
		aid = adapter.GetDefaultAdapterID()
		log.Info("Starting scanner phase")

		scan, err = Run(aid, 60)

		if err != nil {
			log.Error(err)
			os.Exit(1)
		}

		err = postScanResults(scan)
		if err != nil {
			log.Error("Error sending data to server")
			os.Exit(1)
		}
		log.Info("Starting sleeping phase")
		time.Sleep(sleepDuration)
	}
}

func postScanResults(scan Scan) error {
	var url string = "http://127.0.0.1:4000/store"
	var d []byte
	var err error

	d, err = json.Marshal(scan)

	if err != nil {
		return err
	}

	bd := []byte(d)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bd))

	if err != nil {
		return err
	}

	if req == nil {
		return errors.New("Invalid request")
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	resp.Body.Close()

	return nil
}
