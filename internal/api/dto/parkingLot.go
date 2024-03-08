package dto

import (
	"parkwise/internal/domain/bo"
	"time"
)

// ParkingLot represents a parking lot structure.
type ParkingLot struct {
	Name     string `json:"name"`
	Capacity int64  `json:"capacity"`
}

func (pl ParkingLot) Model() bo.ParkingLot {
	return bo.ParkingLot{
		Name:     pl.Name,
		Capacity: pl.Capacity,
	}
}

type ParkingLotStatus []ParkingSlotStatus

func ToParkingLotStatus(boPls bo.ParkingLotStatus) ParkingLotStatus {
	dtoPls := []ParkingSlotStatus{}

	for _, pls := range boPls.SlotsStatus {
		dtoPls = append(dtoPls, ParkingSlotStatus{
			SlotID:        pls.SlotID,
			VehicleNumber: pls.VehicleNumber,
			Status:        pls.Status,
		})
	}

	return dtoPls
}

type DailyReportDate struct {
	Date string `form:"date" json:"date,omitempty"`
}

type DailyReport struct {
	Date                   time.Time `json:"date"`
	TotalVehicles          int64     `json:"totalVehicles"`
	TotatlParkingTimeHours float64   `json:"totatlParkingTimeHours"`
	TotalFeesCollected     int64     `json:"totalFeesCollected"`
}

func ToDailyReport(dr bo.DailyReport) DailyReport {
	return DailyReport{
		Date:                   dr.Date,
		TotalVehicles:          dr.TotalVehicles,
		TotatlParkingTimeHours: dr.TotatlParkingTimeHours,
		TotalFeesCollected:     dr.TotalFeesCollected,
	}
}

type Charge struct {
	VehicleNumber string `json:"vehicleNumber"`
	Bill          int64  `json:"bill"`
}

func ToCharge(c bo.Charge) Charge {
	return Charge{
		VehicleNumber: c.VehicleNumber,
		Bill:          c.Bill,
	}
}
