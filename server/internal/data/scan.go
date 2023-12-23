package data

import (
	"database/sql"
	"fmt"
)

type ScanModel struct {
	DB *sql.DB
}

type Scan struct {
	ScanTime  string `json:"scanTime"`
	ScannerId int    `json:"scannerId"`
}

func (s ScanModel) Insert(scan *Scan) (sql.Result, error) {

	query := "INSERT INTO scan (scanTime, scannerId) values (?, ?);"

	args := []interface{}{scan.ScanTime, scan.ScannerId}
	stm, err := s.DB.Prepare(query)

	if err != nil {
		fmt.Println("UNO", err)
		return nil, err
	}

	rst, err := stm.Exec(args...)

	if err != nil {
		fmt.Println("DUE", err)
		return nil, err
	}

	return rst, nil

}
