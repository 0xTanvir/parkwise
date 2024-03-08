package dto

import "time"

// MaintenanceLog represents a log entry for parking slot maintenance.
type MaintenanceLog struct {
	LogID     int64     `json:"logId"`
	SlotID    int64     `json:"slotId"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Maintenance struct {
	SlotID     int  `json:"slotId"`
    Maintenance bool `json:"maintenance"` // true to enable maintenance mode, false to disable
}