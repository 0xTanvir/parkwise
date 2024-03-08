package pg

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"parkwise/internal/domain/bo"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type parkingLotStore struct {
	dbPool *pgxpool.Pool
}

var parkingLotFields = []string{
	"lot_id",
	"name",
	"capacity",
	"created_at",
	"updated_at",
}

func (s *parkingLotStore) CreateParkingLot(ctx context.Context, parkingLot *bo.ParkingLot) error {
	insertMap := buildParkingLotInsertMap(*parkingLot)
	if len(insertMap) < 1 {
		slog.Debug("empty core insert for parking lot")
		return fmt.Errorf("empty core insert for parking lot")
	}

	start := 1
	arguments := make([]interface{}, 0, len(insertMap))
	fields := []string{}
	placeholders := []string{}

	for field, v := range insertMap {
		fields = append(fields, field)
		arguments = append(arguments, v)
		placeholders = append(placeholders, "$"+strconv.Itoa(start))
		start++
	}

	sqlQuery := fmt.Sprintf("INSERT INTO parking_lots(%s)VALUES (%s) RETURNING lot_id", strings.Join(fields, ","), strings.Join(placeholders, ","))

	conn, err := s.dbPool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	var id sql.NullInt64
	if err := conn.QueryRow(ctx, sqlQuery, arguments...).Scan(&id); err != nil {
		return err
	}

	parkingLot.LotID = id.Int64
	return nil
}

func buildParkingLotInsertMap(i bo.ParkingLot) map[string]interface{} {
	insertedFields := make(map[string]interface{})

	for _, value := range parkingLotFields {
		switch value {
		case "name":
			if i.Name != "" {
				insertedFields[value] = i.Name
			}
		case "capacity":
			if i.Capacity > 0 {
				insertedFields[value] = i.Capacity
			}

		}
	}

	return insertedFields
}

func (s *parkingLotStore) Maintenance(ctx context.Context, maintenance bo.Maintenance) error {
	return WrapInTx(ctx, s.dbPool, func(tx pgx.Tx) error {
		sqlQuery := `UPDATE parking_slots SET status = CASE WHEN $2 THEN 'maintenance' ELSE 'available' END WHERE slot_id = $1;`
		if _, err := tx.Exec(ctx, sqlQuery, maintenance.SlotID, maintenance.Maintenance); err != nil {
			slog.Error("failed to toggle maintanance", slog.Int64("parking slot id", int64(maintenance.SlotID)), "cause", err)
			return err
		}

		return nil
	})
}

func (s *parkingLotStore) ParkingLotStatus(ctx context.Context, parkingLotID int64) (bo.ParkingLotStatus, error) {
	pls := bo.ParkingLotStatus{}

	conn, err := s.dbPool.Acquire(ctx)
	if err != nil {
		return pls, err
	}
	defer conn.Release()

	dbQuery := `SELECT ps.slot_id, ps.status, COALESCE(pss.vehicle_number, '') as vehicle_number
	FROM parking_slots ps
	LEFT JOIN parking_sessions pss ON ps.slot_id = pss.slot_id AND pss.end_time IS NULL
	WHERE ps.lot_id = $1
	ORDER BY ps.slot_id;`

	rows, err := conn.Query(ctx, dbQuery, parkingLotID)
	if err != nil {
		slog.Error("failed to list parking lots", "cause", err)
		return pls, err
	}
	defer rows.Close()

	var pss []bo.ParkingSlotStatus
	for rows.Next() {
		var (
			slotID        sql.NullInt64
			vehicleNumber sql.NullString
			status        sql.NullString
		)
		if err := rows.Scan(&slotID, &vehicleNumber, &status); err != nil {
			slog.Error("failed to scan parking lot row", "cause", err)
			return pls, err
		}
		pss = append(pss, bo.ParkingSlotStatus{
			SlotID:        slotID.Int64,
			VehicleNumber: vehicleNumber.String,
			Status:        status.String,
		})
	}

	if err = rows.Err(); err != nil {
		slog.Error("failed during rows iteration", "cause", err)
		return pls, err
	}

	pls.SlotsStatus = pss

	return pls, nil
}

func (s *parkingLotStore) DailyReport(ctx context.Context, dateDay string) (bo.DailyReport, error) {
	var (
		date                  sql.NullTime
		totalVehicles         sql.NullInt64
		totalParkingTimeHours sql.NullFloat64
		totalFeesCollected    sql.NullInt64
	)
	conn, err := s.dbPool.Acquire(ctx)
	if err != nil {
		return bo.DailyReport{}, err
	}
	defer conn.Release()

	dbQuery := `SELECT
	DATE(start_time) AS date,
	COUNT(DISTINCT vehicle_number) AS totalVehicles,
	SUM(EXTRACT(EPOCH FROM (COALESCE(end_time, CURRENT_TIMESTAMP) - start_time))/3600) AS totalParkingTimeHours,
	SUM(fee) AS totalFeesCollected
	FROM
		public.parking_sessions
	WHERE
		DATE(start_time) = $1
	GROUP BY
		DATE(start_time);
  `

	row := conn.QueryRow(ctx, dbQuery, dateDay)

	if err = row.Scan(&date, &totalVehicles, &totalParkingTimeHours, &totalFeesCollected); err != nil {
		if err == pgx.ErrNoRows {
			slog.Error("no data found for ", slog.String("date", dateDay))
			return bo.DailyReport{}, bo.ErrNoDataFound
		}
	}

	return bo.DailyReport{
		Date:                   date.Time,
		TotalVehicles:          totalVehicles.Int64,
		TotatlParkingTimeHours: totalParkingTimeHours.Float64,
		TotalFeesCollected:     totalFeesCollected.Int64,
	}, nil
}
