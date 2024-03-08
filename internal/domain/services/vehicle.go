package services

import (
	"context"
	"sync"

	"parkwise/internal/domain/bo"
	"parkwise/internal/domain/definition"
)

var onceInitVehicleService sync.Once
var vehicleServiceInstance *vehicleService

type vehicleService struct {
	repo definition.VehicleRepository
}

func Vehicle(vehicleRepo definition.VehicleRepository) *vehicleService {
	onceInitVehicleService.Do(func() {
		vehicleServiceInstance = &vehicleService{
			repo: vehicleRepo,
		}
	})

	return vehicleServiceInstance
}

func (s *vehicleService) Park(ctx context.Context, park bo.Park) error {
	return s.repo.Park(ctx, park)
}

func (s *vehicleService) UnPark(ctx context.Context, unpark bo.UnPark) (bo.Charge, error) {
	return s.repo.UnPark(ctx, unpark)
}
