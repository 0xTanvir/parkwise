package bo

import "time"

// MaintenanceLog represents a log entry for parking slot maintenance.
type MaintenanceLog struct {
	LogID     int64
	SlotID    int64
	StartTime time.Time
	EndTime   time.Time
	Note      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Maintenance struct {
	SlotID     int  
    Maintenance bool 
}