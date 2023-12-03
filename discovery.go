package main

import (
	"time"

	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/bluez/profile/adapter"
	"github.com/muka/go-bluetooth/bluez/profile/device"
	log "github.com/sirupsen/logrus"
)

type Device struct {
	/*
		Address The Bluetooth device address of the remote device.
	*/
	Address string

	/*
		Alias The name alias for the remote device. The alias can
				be used to have a different friendly name for the
				remote device.

				In case no alias is set, it will return the remote
				device name. Setting an empty string as alias will
				convert it back to the remote device name.

				When resetting the alias with an empty string, the
				property will default back to the remote name.
	*/
	Alias string

	/*
		Name The Bluetooth remote name. This value can not be
				changed. Use the Alias property instead.

				This value is only present for completeness. It is
				better to always use the Alias property when
				displaying the devices name.

				If the Alias property is unset, it will reflect
				this value which makes it more convenient.
	*/
	Name string
	/*
		TxPower Advertised transmitted power level (inquiry or
				advertising).
	*/
	TxPower int16

	/*
		RSSI Received Signal Strength Indicator of the remote
				device (inquiry or advertising).
	*/
	RSSI int16
}

type Scanner struct {
	Address string
	Name    string
	Alias   string
}

type Scan struct {
	Devices []Device
	Scanner Scanner
}

func ScannerProps(a adapter.Adapter1) Scanner {
	var s Scanner = Scanner{Address: a.Properties.Address, Name: a.Properties.Name, Alias: a.Properties.Alias}
	return s
}

func PowerCycle(a adapter.Adapter1) {
	log.Infof("Running BT power cycle")
	a.SetPowered(false)
	a.SetPowered(true)
}

/*
Add a device to the devices list only if its address is not in the list
*/
func addDevice(devices *[]Device, d Device) {
	if len(*devices) == 0 {
		*devices = append(*devices, d)
	}

	for j := range *devices {
		if (*devices)[j].Address == d.Address {
			break
		}
		if j == len(*devices)-1 {
			*devices = append(*devices, d)
		}
	}
}

func Run(adapterID string, timer int) (Scan, error) {
	var devices []Device

	//clean up connection on exit
	defer api.Exit()
	a, err := adapter.GetAdapter(adapterID)
	var s Scanner = ScannerProps(*a)

	PowerCycle(*a)

	log.Infof("Scanning started on: %s - %s", s.Address, s.Alias)

	if err != nil {
		return Scan{}, err
	}

	log.Infof("Flush cached devices")
	err = a.FlushDevices()
	if err != nil {
		return Scan{}, err
	}

	discovery, cancel, err := api.Discover(a, nil)
	if err != nil {
		return Scan{}, err
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

		device := Device{Address: dev.Properties.Address, Alias: dev.Properties.Alias, Name: dev.Properties.Name, RSSI: dev.Properties.RSSI, TxPower: dev.Properties.TxPower}
		log.Infof("New device discovered: addr=%s rssi=%d alias=%s name=%s ", device.Address, device.RSSI, device.Alias, device.Name)
		addDevice(&devices, device)
	}

	scan := Scan{Devices: devices, Scanner: s}

	return scan, nil
}
