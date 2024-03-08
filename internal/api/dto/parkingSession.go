package dto

import (
	"parkwise/internal/domain/bo"
	"time"
)

// ParkingSession represents a parking session for a vehicle.
type ParkingSession struct {
	SessionID     int64     `json:"sessionId"`
	SlotID        int64     `json:"slotId"`
	VehicleNumber string    `json:"vehicleNumber"`
	StartTime     time.Time `json:"startTime"`
	EndTime       time.Time `json:"endTime"`
	Fee           int64     `json:"fee"` // This could be calculated at the end based on the parking duration
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type Park struct {
	LotID         int64  `json:"lotId"`
	VehicleNumber string `json:"vehicleNumber"`
}

func (p Park) Model() bo.Park {
	return bo.Park{
		LotID:         p.LotID,
		VehicleNumber: p.VehicleNumber,
	}
}

type UnPark struct {
	SlotID int64 `json:"slotId"`
}

func (up UnPark) Model() bo.UnPark {
	return bo.UnPark{
		SlotID: up.SlotID,
	}
}
