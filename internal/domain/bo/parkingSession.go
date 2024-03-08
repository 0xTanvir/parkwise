package bo

import "time"

// ParkingSession represents a parking session for a vehicle.
type ParkingSession struct {
	SessionID     int64
	SlotID        int64
	VehicleNumber string
	StartTime     time.Time
	EndTime       time.Time
	Fee           int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Park struct {
	LotID         int64
	VehicleNumber string
}

type UnPark struct {
	SlotID int64
}

type Charge struct {
	VehicleNumber string
	Bill          int64
}
