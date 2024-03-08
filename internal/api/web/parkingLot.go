package web

import (
	"context"
	"log/slog"
	"net/http"

	"parkwise/internal/api/dto"
	"parkwise/internal/domain/bo"
	"parkwise/internal/domain/services"

	"github.com/gin-gonic/gin"
)

// Get ParkingLotStatus godoc
// @Summary      Get ParkingLotStatus
// @Description  Get ParkingLotStatus
// @Tags         ParkingLot
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Lot ID"
// @Success      200  {object}  dto.ParkingLotStatus
// @Failure      400  {string} string  "Invalid request body"
// @Failure      500  {string}  string  "Error"
// @Router       /v1/parking-lots/{id} [get]
func (r *repos) getStatus(ctx *gin.Context) {
	var wrappedID dto.IDWrapper
	if err := ctx.ShouldBindUri(&wrappedID); err != nil {
		slog.Error("unable to parse parking lot id", "cause", err)
		ctx.JSON(http.StatusBadRequest, dto.Builder().SetMessage("Invalid query value"))
		return
	}

	getStatusCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pls, err := services.ParkingLot(r.ds.ParkingLot).Status(getStatusCtx, wrappedID.ID)
	if err != nil {
		slog.Error("unable to get parking lot status", "cause", err)
		ctx.JSON(http.StatusInternalServerError, dto.Builder().SetMessage("Internal server error"))
		return
	}

	ctx.JSON(http.StatusOK, dto.ToParkingLotStatus(pls))
}

// Get Parking slot maintenance godoc
// @Summary      Parking slot maintenance
// @Description  Parking slot maintenance
// @Tags         ParkingLot
// @Accept       json
// @Produce      json
// @Param        request body dto.Maintenance  true  "Maintenance params"
// @Success      200  {string}  "Maintenance updated"
// @Failure      400  {string} string  "Invalid request body"
// @Failure      404  {object}  dto.Error
// @Failure      500  {string}  string  "Error"
// @Router       /v1/slots/maintenance [post]
func (r *repos) maintenance(ctx *gin.Context) {
	parkingLotDto := dto.Maintenance{}
	if err := ctx.ShouldBindJSON(&parkingLotDto); err != nil {
		slog.Error("unable to parse maintenanceCtx from request body", "cause", err)
		ctx.JSON(http.StatusBadRequest, dto.Builder().SetMessage("Invalid request body"))
		return
	}

	maintenanceCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := services.ParkingLot(r.ds.ParkingLot).Maintenance(maintenanceCtx, bo.Maintenance{SlotID: parkingLotDto.SlotID, Maintenance: parkingLotDto.Maintenance})
	if err != nil {
		slog.Error("unable to get parking lot from database: ", "cause", err)
		ctx.JSON(http.StatusInternalServerError, dto.Builder().SetMessage("Error"))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Maintenance updated"})
}

// Add Parking Lot godoc
// @Summary      Add a new Parking Lot
// @Description  Create a new Parking Lot in the system
// @Tags         ParkingLot
// @Accept       json
// @Produce      json
// @Param        request body dto.ParkingLot  true  "ParkingLot params"
// @Success      201  {object}  dto.IDWrapper
// @Failure      400  {string} string  "Invalid request body"
// @Failure      404  {object}  dto.Error
// @Failure      500  {string}  string  "Error"
// @Router       /v1/parking-lots [post]
func (r *repos) addParkingLot(ctx *gin.Context) {
	parkingLotDto := dto.ParkingLot{}
	if err := ctx.ShouldBindJSON(&parkingLotDto); err != nil {
		slog.Error("unable to parse parking lot from request body", "cause", err)
		ctx.JSON(http.StatusBadRequest, dto.Builder().SetMessage("Invalid request body"))
		return
	}

	model := parkingLotDto.Model()

	addParkingLotCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	id, err := services.ParkingLot(r.ds.ParkingLot).Create(addParkingLotCtx, model)
	if err != nil {
		slog.Error("unable to create parking lot", "cause", err)
		ctx.JSON(http.StatusInternalServerError, dto.Builder().SetMessage("Internal server error"))
		return
	}

	ctx.JSON(http.StatusOK, dto.IDWrapper{ID: id})
}

// DailyReport godoc
// @Summary      DailyReport
// @Description  DailyReport
// @Tags         ParkingLot
// @Accept       json
// @Produce      json
// @Param        date      query  string  true  "date=YYYY-MM-DD"
// @Success      200  {object}  dto.DailyReport
// @Failure      400  {string} string  "Invalid request body"
// @Failure      404  {object}  dto.Error
// @Failure      500  {string}  string  "Error"
// @Router       /v1/reports/daily [get]
func (r *repos) getDailyReports(ctx *gin.Context) {
	var reportDate dto.DailyReportDate
	if err := ctx.ShouldBindQuery(&reportDate); err != nil {
		slog.Error("unable to parse query url", "cause", err)
		ctx.JSON(http.StatusBadRequest, dto.Builder().SetMessage("Invalid query value"))
		return
	}

	reportCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dr, err := services.ParkingLot(r.ds.ParkingLot).DailyReport(reportCtx, reportDate.Date)
	if err != nil {
		slog.Error("unable to get daily report", "cause", err)
		ctx.JSON(http.StatusInternalServerError, dto.Builder().SetMessage("Internal server error"))
		return
	}

	ctx.JSON(http.StatusOK, dto.ToDailyReport(dr))
}
