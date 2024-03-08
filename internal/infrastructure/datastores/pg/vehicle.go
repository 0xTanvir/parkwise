package pg

import (
	"context"
	"database/sql"
	"log/slog"
	"math"
	"time"

	"parkwise/internal/domain/bo"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type vehicleStore struct {
	dbPool *pgxpool.Pool
}

func (s *vehicleStore) Park(ctx context.Context, park bo.Park) error {
	return WrapInTx(ctx, s.dbPool, func(tx pgx.Tx) error {
		bookSlotQuery := `WITH available_slot AS 
		(SELECT slot_id FROM parking_slots WHERE status = 'available' AND lot_id = $1 
		ORDER BY slot_number LIMIT 1) 
		UPDATE parking_slots SET status = 'occupied'
		WHERE slot_id IN (SELECT slot_id FROM available_slot) RETURNING slot_id;`

		var slotID sql.NullInt64
		if err := tx.QueryRow(ctx, bookSlotQuery, park.LotID).Scan(&slotID); err != nil {
			slog.Error("failed to book slot", slog.Int64("slotID", park.LotID), "cause", err)
			return err
		}

		startSessionQuery := `INSERT INTO parking_sessions(
			slot_id, vehicle_number, start_time)
			VALUES ($1, $2, $3);`

		if _, err := tx.Exec(ctx, startSessionQuery, slotID.Int64, park.VehicleNumber, time.Now()); err != nil {
			return err
		}

		return nil
	})
}

func (s *vehicleStore) UnPark(ctx context.Context, unpark bo.UnPark) (bo.Charge, error) {
	var (
		vNumber sql.NullString
		hours   sql.NullFloat64
	)

	if err := WrapInTx(ctx, s.dbPool, func(tx pgx.Tx) error {
		updateSlotQuery := `UPDATE parking_slots SET status = 'available' WHERE slot_id = $1;`
		if _, err := tx.Exec(ctx, updateSlotQuery, unpark.SlotID); err != nil {
			slog.Error("failed to update slot", slog.Int64("slotID", unpark.SlotID), "cause", err)
			return err
		}

		endSessionQuery := `UPDATE parking_sessions SET end_time = $2 WHERE slot_id = $1
		RETURNING vehicle_number, EXTRACT(EPOCH FROM (NOW() - start_time)) / 3600 AS hours_parked;`

		if err := tx.QueryRow(ctx, endSessionQuery, unpark.SlotID, time.Now()).Scan(&vNumber, &hours); err != nil {
			slog.Error("failed to update session", slog.Int64("slotID", unpark.SlotID), "cause", err)
			return err
		}

		updateFee := `UPDATE parking_sessions SET fee = $2, end_time = $3 WHERE slot_id = $1;`
		if _, err := tx.Exec(ctx, updateFee, unpark.SlotID, int64(10 * math.Ceil(hours.Float64)), time.Now()); err != nil {
			slog.Error("failed to update fee", slog.Int64("slotID", unpark.SlotID), "cause", err)
			return err
		}

		return nil
	}); err != nil {
		return bo.Charge{}, err
	}

	return bo.Charge{
		VehicleNumber: vNumber.String,
		Bill:          int64(10 * math.Ceil(hours.Float64)),
	}, nil
}
