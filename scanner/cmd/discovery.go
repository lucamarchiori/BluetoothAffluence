package main

import (
	"time"

	"github.com/muka/go-bluetooth/api"
	"github.com/muka/go-bluetooth/bluez/profile/adapter"
	"github.com/muka/go-bluetooth/bluez/profile/device"
)

type Device struct {
	/*
		Address The Bluetooth device address of the remote device.
	*/
	Address string `json:"address"`

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
	Alias string `json:"alias"`

	/*
		Name The Bluetooth remote name. This value can not be
				changed. Use the Alias property instead.

				This value is only present for completeness. It is
				better to always use the Alias property when
				displaying the devices name.

				If the Alias property is unset, it will reflect
				this value which makes it more convenient.
	*/
	Name string `json:"name"`
	/*
		TxPower Advertised transmitted power level (inquiry or
				advertising).
	*/
	TxPower int16 `json:"txPower"`

	/*
		RSSI Received Signal Strength Indicator of the remote
				device (inquiry or advertising).
	*/
	RSSI int16 `json:"rssi"`
}

type Scanner struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	Alias   string `json:"alias"`
}

type Scan struct {
	Devices  []Device `json:"devices"`
	Scanner  Scanner  `json:"scanner"`
	ScanTime string   `json:"scanTime"`
}

func ScannerProps(a *adapter.Adapter1) Scanner {
	var s Scanner
	if a != nil && a.Properties != nil && a.Properties.Address != "" {
		s.Address = a.Properties.Address
	}
	if a != nil && a.Properties != nil && a.Properties.Name != "" {
		s.Name = a.Properties.Name
	}
	if a != nil && a.Properties != nil && a.Properties.Alias != "" {
		s.Alias = a.Properties.Alias
	}
	return s
}

func PowerCycle(a *adapter.Adapter1) {
	log.Infof("Running BT power cycle")
	if a != nil {
		a.SetPowered(false)
		a.SetPowered(true)
	}
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

/*
Run performs Bluetooth device scanning using the specified adapterID and timer duration.
*/
func Run(adapterID string, timer int) (Scan, error) {
	// Initialize devices slice
	var devices []Device

	// Clean up connection on exit
	//defer api.Exit()

	// Get the Bluetooth adapter by adapterID
	a, err := adapter.GetAdapter(adapterID)

	// Create a Scanner object from the adapter properties
	var s Scanner = ScannerProps(a)

	// Power cycle the adapter
	PowerCycle(a)

	log.Infof("Scanning started on: %s - %s", s.Address, s.Alias)

	if err != nil {
		return Scan{}, err
	}

	// Flush cached devices
	log.Infof("Flush cached devices")
	err = a.FlushDevices()
	if err != nil {
		return Scan{}, err
	}

	// Start device discovery
	discovery, cancel, err := api.Discover(a, nil)
	if err != nil {
		return Scan{}, err
	}

	// Set up a timer to cancel the discovery after a specified duration
	go func() {
		log.Infof("Starting timer %d seconds", timer)
		time.Sleep(time.Duration(timer) * time.Second)
		log.Infof("Timer expired")
		cancel()
	}()

	// Process events from the discovery
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

		// Add the discovered device to the devices slice
		addDevice(&devices, device)
	}

	// Get the current time and format it
	dt := time.Now().Format(time.RFC3339)

	// Create a Scan object with the collected information
	scan := Scan{Devices: devices, Scanner: s, ScanTime: dt}

	return scan, nil
}

/*
RunMock simulates a Bluetooth scan and returns a Scan object along with any potential error.
*/
func RunMock() (Scan, error) {
	// Create an empty slice to store devices
	var devices []Device

	// Create a Scanner object with mock information
	scanner := Scanner{Address: "AA:AA:AA:AA:AA:AA", Name: "Mock scanner", Alias: "Mock scanner"}

	// Create and add two mock devices to the devices slice
	device := Device{Address: "99:99:99:99:99:99", Name: "Mock device", TxPower: 99, RSSI: 99}
	addDevice(&devices, device)
	device = Device{Address: "88:99:99:99:99:99", Name: "Mock device 2", TxPower: 99, RSSI: 99}
	addDevice(&devices, device)

	// Get the current time and format it
	dt := time.Now().Format(time.RFC3339)

	// Create a Scan object with the collected information
	scan := Scan{Devices: devices, Scanner: scanner, ScanTime: dt}

	// Return the Scan object and a nil error
	return scan, nil
}
