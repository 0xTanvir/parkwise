package bo

import (
	"errors"
	"time"
)

var (
	// ErrLotNotFound is returned when a lot is not found.
	ErrLotNotFound = errors.New("the parking lot was not found")
	ErrNoDataFound = errors.New("no data found")
)

// ParkingLot represents a parking lot structure.
type ParkingLot struct {
	LotID     int64
	Name      string
	Capacity  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ParkingLotStatus struct {
	SlotsStatus []ParkingSlotStatus
}

type DailyReport struct {
	Date                   time.Time
	TotalVehicles          int64
	TotatlParkingTimeHours float64
	TotalFeesCollected     int64
}
