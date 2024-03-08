package dto

import "time"

// ParkingSlot represents a parking slot within a parking lot.
type ParkingSlot struct {
	SlotID     int64     `json:"slotId"`
	LotID      int64     `json:"lotId"`
	SlotNumber int64     `json:"slotNumber"`
	Status     string    `json:"status"` // e.g., "available", "occupied", "maintenance"
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type ParkingSlotStatus struct {
	SlotID        int64  `json:"slotId"`
	VehicleNumber string `json:"vehicleNumber"`
	Status        string `json:"status"`
}
