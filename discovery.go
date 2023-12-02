package main

import (
	"fmt"
	"time"

	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/bluez/profile/adapter"
	"github.com/muka/go-bluetooth/bluez/profile/device"
	log "github.com/sirupsen/logrus"
)

var devices []device.Device1

func Run(adapterID string) ([]device.Device1, error) {
	fmt.Println("Scanning started")

	//clean up connection on exit
	defer api.Exit()

	a, err := adapter.GetAdapter(adapterID)

	// powercycle
	a.SetPowered(false)
	a.SetPowered(true)

	if err != nil {
		return nil, err
	}

	fmt.Println("Flush cached devices")
	err = a.FlushDevices()
	if err != nil {
		return nil, err
	}

	fmt.Println("Start discovery")
	discovery, cancel, err := api.Discover(a, nil)
	if err != nil {
		return nil, err
	}

	go func() {
		fmt.Println("Starting timer")
		time.Sleep(5 * time.Second)
		fmt.Println("Timer expired")
		cancel()
	}()

	for ev := range discovery {
		if ev.Type == adapter.DeviceRemoved {
			continue
		}

		dev, err := device.NewDevice1(ev.Path)
		if err != nil {
			log.Errorf("%s: %s", ev.Path, err)
			continue
		}

		if dev == nil {
			log.Errorf("%s: not found", ev.Path)
			continue
		}

		devices = append(devices, *dev)
		log.Infof("name=%s addr=%s rssi=%d", dev.Properties.Name, dev.Properties.Address, dev.Properties.RSSI)
		fmt.Println("New")
	}

	fmt.Println("FInish search loop")

	return devices, nil
}
