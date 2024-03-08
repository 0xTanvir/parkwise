package bo

import "time"

// ParkingSlot represents a parking slot within a parking lot.
type ParkingSlot struct {
	SlotID     int64
	LotID      int64
	SlotNumber int64
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ParkingSlotStatus struct {
	SlotID        int64
	VehicleNumber string
	Status        string
}