package web

import (
	"context"
	"log/slog"
	"net/http"

	"parkwise/internal/api/dto"
	"parkwise/internal/domain/services"

	"github.com/gin-gonic/gin"
)

// Park godoc
// @Summary      Park
// @Description  Park
// @Tags         Vehicle
// @Accept       json
// @Produce      json
// @Param        request body dto.Park  true  "Park params"
// @Success      200  {string}  "parked successfully"
// @Failure      400  {string} string  "Invalid request body"
// @Failure      404  {object}  dto.Error
// @Failure      500  {string}  string  "Error"
// @Router       /v1/park [post]
func (r *repos) park(ctx *gin.Context) {
	parkDto := dto.Park{}
	if err := ctx.ShouldBindJSON(&parkDto); err != nil {
		slog.Error("unable to parse park from request body", "cause", err)
		ctx.JSON(http.StatusBadRequest, dto.Builder().SetMessage("Invalid request body"))
		return
	}

	model := parkDto.Model()

	parkCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := services.Vehicle(r.ds.Vehicle).Park(parkCtx, model)
	if err != nil {
		slog.Error("unable to park", "cause", err)
		ctx.JSON(http.StatusInternalServerError, dto.Builder().SetMessage("Internal server error"))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "parked successfully"})
}

// UnPark godoc
// @Summary      UnPark
// @Description  UnPark
// @Tags         Vehicle
// @Accept       json
// @Produce      json
// @Param        request body dto.UnPark  true  "UnPark params"
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  dto.Charge
// @Failure      400  {string} string  "Invalid request body"
// @Failure      404  {object}  dto.Error
// @Failure      500  {string}  string  "Error"
// @Router       /v1/unpark [post]
func (r *repos) unpark(ctx *gin.Context) {
	unparkDto := dto.UnPark{}
	if err := ctx.ShouldBindJSON(&unparkDto); err != nil {
		slog.Error("unable to parse park from request body", "cause", err)
		ctx.JSON(http.StatusBadRequest, dto.Builder().SetMessage("Invalid request body"))
		return
	}

	model := unparkDto.Model()

	unparkCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	charge, err := services.Vehicle(r.ds.Vehicle).UnPark(unparkCtx, model)
	if err != nil {
		slog.Error("unable to unpark", "cause", err)
		ctx.JSON(http.StatusInternalServerError, dto.Builder().SetMessage("Internal server error"))
		return
	}

	ctx.JSON(http.StatusOK, dto.ToCharge(charge))
}
