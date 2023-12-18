package data

import (
	"database/sql"
	"fmt"
)

type ScanModel struct {
	DB *sql.DB
}

type Scan struct {
	ScanTime string `json:"scanTime"`
}

func (s ScanModel) Insert(scan *Scan) (sql.Result, error) {

	query := "INSERT INTO scan (scanTime) values (?);"

	args := []interface{}{scan.ScanTime}
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
