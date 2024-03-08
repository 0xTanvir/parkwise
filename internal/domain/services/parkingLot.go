package services

import (
	"context"
	"log/slog"
	"sync"

	"parkwise/internal/domain/bo"
	"parkwise/internal/domain/definition"
)

var onceInitParkingLotService sync.Once
var parkingLotServiceInstance *parkingLotService

type parkingLotService struct {
	repo definition.ParkingLotRepository
}

func ParkingLot(parkingLotRepo definition.ParkingLotRepository) *parkingLotService {
	onceInitParkingLotService.Do(func() {
		parkingLotServiceInstance = &parkingLotService{
			repo: parkingLotRepo,
		}
	})

	return parkingLotServiceInstance
}

func (s *parkingLotService) Status(ctx context.Context, parkingLotID int64) (bo.ParkingLotStatus, error) {
	return s.repo.ParkingLotStatus(ctx, parkingLotID)
}

func (s *parkingLotService) Maintenance(ctx context.Context, maintenance bo.Maintenance) error {
	return s.repo.Maintenance(ctx, maintenance)
}

func (s *parkingLotService) Create(ctx context.Context, parkingLot bo.ParkingLot) (int64, error) {
	if err := s.repo.CreateParkingLot(ctx, &parkingLot); err != nil {
		return -1, err
	}

	if parkingLot.LotID < 1 {
		slog.Warn("inserted parking lot has invalid id", slog.String("name", parkingLot.Name))
	}
	return parkingLot.LotID, nil
}

func (s *parkingLotService) DailyReport(ctx context.Context, dateStr string) (bo.DailyReport, error) {
	return s.repo.DailyReport(ctx, dateStr)
}
