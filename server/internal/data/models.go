package data

import (
	"database/sql"
	"errors"
)

var err error

// Define a custom ErrRecordNotFound error. We'll return this from our Get() method when
// looking up a movie that doesn't exist in our database.
var (
	ErrRecordNotFound = errors.New("record not found")
)

// Create a Models struct which wraps the MovieModel. We'll add other models to this,
// like a UserModel and PermissionModel, as our build progresses.
type Models struct {
	Device  DeviceModel
	Scan    ScanModel
	Scanner ScannerModel
}

// For ease of use, we also add a New() method which returns a Models struct containing
// the initialized MovieModel.
func NewModels(db *sql.DB) Models {
	return Models{
		Device:  DeviceModel{DB: db},
		Scan:    ScanModel{DB: db},
		Scanner: ScannerModel{DB: db},
	}
}
