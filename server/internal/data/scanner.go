package data

import (
	"database/sql"
	"fmt"
)

type ScannerModel struct {
	DB *sql.DB
}

type Scanner struct {
	Id      int    `json:"id"`
	Address string `json:"address"`
	Alias   string `json:"alias"`
	Name    string `json:"name"`
}

// Insert a scanner in the database
func (s ScannerModel) Insert(scanner *Scanner) (sql.Result, error) {

	query := "INSERT OR IGNORE INTO scanners (address, alias, name) values (?, ?, ?);"

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

// Get all scanners
func (s ScannerModel) Index() (*[]Scanner, error) {
	query := "SELECT * FROM SCANNERS"
	rsp, err := s.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rsp.Close()
	var scanners []Scanner

	for rsp.Next() {
		var scanner Scanner
		err := rsp.Scan(&scanner.Id, &scanner.Address, &scanner.Alias, &scanner.Name)
		if err != nil {
			return nil, err
		}
		scanners = append(scanners, scanner)
	}

	if err != nil {
		return nil, err
	}

	return &scanners, nil
}

// Get a scanner by its address
func (s ScannerModel) GetByAddress(address string) (*Scanner, error) {

	var rows []Scanner
	query := "SELECT * FROM scanners where address = ?;"
	rsp, err := s.DB.Query(query, address)

	if err != nil {
		return nil, err
	}

	defer rsp.Close()

	for rsp.Next() {
		var scanner Scanner
		err := rsp.Scan(&scanner.Id, &scanner.Address, &scanner.Alias, &scanner.Name)
		if err != nil {
			return nil, err
		}
		rows = append(rows, scanner)
	}

	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, nil
	}

	return &rows[0], nil
}
