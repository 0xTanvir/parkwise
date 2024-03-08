package definition

import (
	"context"

	"parkwise/internal/domain/bo"
)

type DataStore struct {
	ParkingLot ParkingLotRepository
	Vehicle    VehicleRepository
}

// ParkingLotRepository is the interface that wraps the basic CRUD operations
// defines the rules around what a Parking Lot repository has to be able to perform
// For datastore implementations, see internal/infrastructure/datastores
type ParkingLotRepository interface {
	ParkingLotStatus(ctx context.Context, parkingLotID int64) (bo.ParkingLotStatus, error)
	CreateParkingLot(ctx context.Context, parkingLot *bo.ParkingLot) error
	Maintenance(ctx context.Context, maintenance bo.Maintenance) error
	DailyReport(ctx context.Context, date string) (bo.DailyReport, error)
}

// VehicleRepository is the interface that wraps the basic CRUD operations
// defines the rules around what a Vehicle repository has to be able to perform
// For datastore implementations, see internal/infrastructure/datastores
type VehicleRepository interface {
	Park(ctx context.Context, park bo.Park) error
	UnPark(ctx context.Context, unpark bo.UnPark) (bo.Charge, error)
}
