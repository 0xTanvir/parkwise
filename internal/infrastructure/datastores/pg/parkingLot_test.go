package pg

import (
	"context"
	"testing"

	"parkwise/internal/domain/algo"
	"parkwise/internal/domain/bo"

	"github.com/stretchr/testify/require"
)

// Return parkingLot with happy case
func createGoodRandomParkingLot(t *testing.T) *bo.ParkingLot {
	parkingLotCreate := &bo.ParkingLot{
		Name:     algo.GenerateRandomString(10),
		Capacity: 50,
	}

	err := testStore.ParkingLot.CreateParkingLot(context.Background(), parkingLotCreate)
	require.NoError(t, err)
	require.NotEmpty(t, parkingLotCreate.LotID)

	return parkingLotCreate
}

func TestAddParkingLot(t *testing.T) {
	createGoodRandomParkingLot(t)
	parkingLotWithoutName := &bo.ParkingLot{
		Capacity: 50,
	}

	err := testStore.ParkingLot.CreateParkingLot(context.Background(), parkingLotWithoutName)
	require.Error(t, err)
	require.Empty(t, parkingLotWithoutName.LotID)

	parkingLotEmpty := &bo.ParkingLot{}
	err = testStore.ParkingLot.CreateParkingLot(context.Background(), parkingLotEmpty)
	require.Error(t, err)
	require.Empty(t, parkingLotEmpty.LotID)
}

func TestParkingLotStatus(t *testing.T) {
	parkingLotCreate := createGoodRandomParkingLot(t)

	parkingLotGet, err := testStore.ParkingLot.ParkingLotStatus(context.Background(), parkingLotCreate.LotID)
	require.NoError(t, err)
	require.NotEmpty(t, parkingLotGet)
}
