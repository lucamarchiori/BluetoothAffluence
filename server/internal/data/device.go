package data

import (
	"database/sql"
	"fmt"
)

var err error

type DeviceModel struct {
	DB *sql.DB
}

type Device struct {
	Address string `json:"address"`
	Alias   string `json:"alias"`
	Name    string `json:"name"`
	TxPower int16  `json:"txPower"`
	RSSI    int16  `json:"rssi"`
}

func (d DeviceModel) Insert(device *Device) (sql.Result, error) {

	query := "INSERT INTO devices (address, alias, name, txPower, rssi) values (?, ?, ?, ?, ?);"

	args := []interface{}{device.Address, device.Alias, device.Name, device.TxPower, device.RSSI}
	stm, err := d.DB.Prepare(query)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	rst, err := stm.Exec(args...)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return rst, nil

}
