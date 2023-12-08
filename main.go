package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/muka/go-bluetooth/bluez/profile/adapter"
)

func main() {
	aid := adapter.GetDefaultAdapterID()
	scan, err := Run(aid, 20)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res := postScanResults(scan)
	if !res {
		log.Println("Error sending data to server")
	}

	os.Exit(0)
}

func postScanResults(scan Scan) bool {
	var url string = "https://webhook.site/0435b241-6cd6-4707-a286-35da5feda75a"
	var d []byte
	var err error

	d, err = json.Marshal(scan)

	if err != nil {
		log.Println(err)
		return false
	}

	bd := []byte(d)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bd))

	if err != nil {
		log.Println(err)
		return false
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return false
	}
	defer resp.Body.Close()

	return true
}
