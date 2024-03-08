package mockdb

import (
	"parkwise/internal/domain/definition"

	"go.uber.org/mock/gomock"
)

func GetInstance(ctrl *gomock.Controller) definition.DataStore {
	return definition.DataStore{
		ParkingLot: NewMockParkingLotRepository(ctrl),
		Vehicle:    NewMockVehicleRepository(ctrl),
	}
}
