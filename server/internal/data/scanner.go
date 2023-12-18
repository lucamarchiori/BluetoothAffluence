package data

import (
	"database/sql"
	"fmt"
)

type ScannerModel struct {
	DB *sql.DB
}

type Scanner struct {
	Address string `json:"address"`
	Alias   string `json:"alias"`
	Name    string `json:"name"`
}

func (s ScannerModel) Insert(scanner *Scanner) (sql.Result, error) {

	query := "INSERT INTO scanners (address, alias, name) values (?, ?, ?);"

	args := []interface{}{scanner.Address, scanner.Alias, scanner.Name}
	stm, err := s.DB.Prepare(query)

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
