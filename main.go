package main

import (
	"fmt"

	"github.com/muka/go-bluetooth/bluez/profile/adapter"
)

func main() {
	fmt.Println("Hello")
	aid := adapter.GetDefaultAdapterID()
	devices, err := Run(aid, 60)

	if err != nil {
		fmt.Println(err)
	}

	for i := range devices {
		d := devices[i]
		fmt.Println(d.Properties.Name, d.Properties.Address, d.Properties.RSSI)
	}
	return

}
