package web

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"parkwise/config"
	"parkwise/internal/domain/algo"
	"parkwise/internal/domain/bo"
	"parkwise/internal/infrastructure/datastores/mockdb"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetParkingLotAPI(t *testing.T) {
	parkingLot := randomParkingLot()

	// TODO: Add multiple test cases
	// for each error scenarios
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	appConfig, err := config.Parse()
	if err != nil {
		slog.Error("Error parsing config", "cause", err)
	}

	// Get a mockdb datastore instance
	ds := mockdb.GetInstance(ctrl)
	apiService := NewAPIService(*appConfig.Server, ds)

	parkingLotStore := ds.ParkingLot.(*mockdb.MockParkingLotRepository)
	// build stubs
	parkingLotStore.EXPECT().
		ParkingLotStatus(gomock.Any(), gomock.Eq(parkingLot.LotID)).
		Times(1).
		Return(parkingLot, nil)

	recorder := httptest.NewRecorder()
	url := fmt.Sprintf("/v1/parking-lots/%d", parkingLot.LotID)
	req, err := http.NewRequest("GET", url, nil)
	require.NoError(t, err)

	router := gin.Default()
	apiService.InstallRoutes(router)

	router.ServeHTTP(recorder, req)
	// Check response
	require.Equal(t, http.StatusOK, recorder.Code)
}

func randomParkingLot() bo.ParkingLot {
	return bo.ParkingLot{
		LotID:    int64(algo.GenerateRandomInteger(1, 1000)),
		Name:     algo.GenerateRandomString(10),
		Capacity: 20,
	}
}
