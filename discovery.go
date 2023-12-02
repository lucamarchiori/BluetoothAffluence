package main

import (
	"time"

	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/bluez/profile/adapter"
	"github.com/muka/go-bluetooth/bluez/profile/device"
	log "github.com/sirupsen/logrus"
)

var devices []device.Device1

func Run(adapterID string, timer int) ([]device.Device1, error) {
	log.Infof("Scanning started")

	//clean up connection on exit
	defer api.Exit()

	a, err := adapter.GetAdapter(adapterID)

	log.Infof("Running BT power cycle")
	a.SetPowered(false)
	a.SetPowered(true)

	if err != nil {
		return nil, err
	}

	log.Infof("Flush cached devices")
	err = a.FlushDevices()
	if err != nil {
		return nil, err
	}

	discovery, cancel, err := api.Discover(a, nil)
	if err != nil {
		return nil, err
	}

	go func() {
		log.Infof("Starting timer %d seconds", timer)
		time.Sleep(time.Duration(timer) * time.Second)
		log.Infof("Timer expired")
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
		log.Infof("New device discovered: name=%s addr=%s rssi=%d", dev.Properties.Name, dev.Properties.Address, dev.Properties.RSSI)
	}

	log.Infof("Scanning complete")

	return devices, nil
}
