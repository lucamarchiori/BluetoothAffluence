package data

import (
	"database/sql"
	"math"
	"time"
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

type scanDeviceAvg struct {
	ScanTime string  `json:"scanTime"`
	Count    float64 `json:"count"`
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func daysBetween(startDateStr, endDateStr string) (float64, error) {
	startDate, err := time.Parse("2006-01-02 15:04:05", startDateStr)
	if err != nil {
		return 0, err
	}

	endDate, err := time.Parse("2006-01-02 15:04:05", endDateStr)
	if err != nil {
		return 0, err
	}

	duration := endDate.Sub(startDate)
	days := duration.Hours() / 24

	return days, nil
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
	AND devices.rssi > -90
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
			AND devices.rssi > -90
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

func (s ScanModel) CountScanDevicesTimeAvg(scannerId int, startDate string, endDate string) (*[]scanDeviceAvg, error) {
	query := `SELECT strftime('%H:%M', scan.scanTime) AS minuteOfDay, COUNT(devices.id) AS averageValue
	FROM scan
	LEFT JOIN devices ON scan.id = devices.scanID
	WHERE scan.scannerID = ? AND scan.scanTime BETWEEN ? AND ?
	AND devices.rssi > -90

	GROUP BY minuteOfDay;
				`

	rows, err := s.DB.Query(query, scannerId, startDate, endDate)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var results []scanDeviceAvg

	days, err := daysBetween(startDate, endDate)
	if err != nil {
		return nil, err
	}

	if days == 0 {
		days = 0.1
	}

	// Iterate through the result set
	for rows.Next() {
		var scanTime string
		var numDevices int

		// Scan values from the result set into variables
		err := rows.Scan(&scanTime, &numDevices)
		if err != nil {
			return nil, err
		}

		res := scanDeviceAvg{scanTime, toFixed(float64(numDevices)/days, 2)}
		results = append(results, res)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &results, nil
}
