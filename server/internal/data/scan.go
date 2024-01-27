package data

import (
	"database/sql"
)

type ScanModel struct {
	DB *sql.DB
}

type Scan struct {
	ScanTime  string `json:"scanTime"`
	ScannerId int    `json:"scannerId"`
}

type scanDeviceCount struct {
	ScanTime string `json:"scanTime"`
	Count    int    `json:"count"`
}

func (s ScanModel) Insert(scan *Scan) (sql.Result, error) {

	query := "INSERT INTO scan (scanTime, scannerId) values (?, ?);"

	args := []interface{}{scan.ScanTime, scan.ScannerId}
	stm, err := s.DB.Prepare(query)

	if err != nil {
		return nil, err
	}

	rst, err := stm.Exec(args...)

	if err != nil {
		return nil, err
	}

	return rst, nil
}

func (s ScanModel) CountScanDevices(scannerId int, startDate string, endDate string) (*[]scanDeviceCount, error) {

	query := `SELECT scan.scanTime, COUNT(devices.id) AS numDevices FROM scan
	LEFT JOIN devices ON scan.id = devices.scanID
	WHERE scan.scannerID = ?
	AND scan.scanTime BETWEEN ? AND ?
	GROUP BY scan.scanTime;`

	rows, err := s.DB.Query(query, scannerId, startDate, endDate)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var results []scanDeviceCount

	// Iterate through the result set
	for rows.Next() {
		var scanTime string
		var numDevices int

		// Scan values from the result set into variables
		err := rows.Scan(&scanTime, &numDevices)
		if err != nil {
			return nil, err
		}

		res := scanDeviceCount{scanTime, numDevices}
		results = append(results, res)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &results, nil
}


func (s ScanModel) CountScanDevicesMovingAVG(scannerId int, startDate string, endDate string) (*[]scanDeviceCount, error) {
    query := `SELECT scan.scanTime, COUNT(devices.id) AS numDevices FROM scan
              LEFT JOIN devices ON scan.id = devices.scanID
              WHERE scan.scannerID = ?
              AND scan.scanTime BETWEEN ? AND ?
              GROUP BY scan.scanTime;`

    rows, err := s.DB.Query(query, scannerId, startDate, endDate)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var results []scanDeviceCount

    // Iterate through the result set
    for rows.Next() {
        var scanTime string
        var numDevices int

        // Scan values from the result set into variables
        if err := rows.Scan(&scanTime, &numDevices); err != nil {
            return nil, err
        }

        res := scanDeviceCount{scanTime, numDevices}
        results = append(results, res)
    }

    // Check for errors from iterating over rows
    if err := rows.Err(); err != nil {
        return nil, err
    }

    // Calculate moving average
    windowSize := 5
    if len(results) >= windowSize {
        for i := windowSize - 1; i < len(results); i++ {
            sum := 0
            for j := i - (windowSize - 1); j <= i; j++ {
                sum += results[j].Count
            }
            results[i].Count = sum / windowSize
        }
    }

    return &results, nil
}
