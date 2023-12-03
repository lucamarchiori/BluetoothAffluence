package main

import (
	"fmt"

	"github.com/muka/go-bluetooth/bluez/profile/adapter"
)

func main() {
	aid := adapter.GetDefaultAdapterID()
	scan, err := Run(aid, 60)

	if err != nil {
		fmt.Println(err)
	}

	for i := range scan.Devices {
		d := scan.Devices[i]
		fmt.Println(d.Address, d.RSSI, d.Name)
	}

}
