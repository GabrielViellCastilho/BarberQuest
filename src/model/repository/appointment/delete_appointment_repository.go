package appointment

import (
	"context"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/logger"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func (ar *appointmentRepository) DeleteAppointmentByIdAndUserID(ctx context.Context, id int, userId int) *rest_err.RestErr {
	logger.Info("Init DeleteAppointmentByIdAndUserID repository", zap.String("journey", "DeleteAppointmentByIdAndUserID"))

	tx, err := ar.databaseConection.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		logger.Error("Error starting transaction", err)
		return rest_err.NewInternalServerError("Error starting transaction")
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback(ctx)
			panic(p)
		} else if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	query := `DELETE FROM appointments WHERE id = $1 AND user_id = $2;`
	result, err := tx.Exec(ctx, query, id, userId)
	if err != nil {
		logger.Error("Error executing delete", err)
		return rest_err.NewInternalServerError("Database error during appointment_domain deletion")
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected > 1 {
		logger.Error("Multiple rows affected during delete", pgx.ErrTooManyRows, zap.Int64("rowsAffected", rowsAffected))
		_ = tx.Rollback(ctx)
		return rest_err.NewInternalServerError("Multiple rows affected. Rolling back transaction.")
	} else if rowsAffected == 0 {
		logger.WarnWithoutError("No appointment_domain found with given ID", zap.Int("id", id))
		return rest_err.NewNotFoundError("Appointment not found")
	}

	if err := tx.Commit(ctx); err != nil {
		logger.Error("Error committing transaction", err)
		return rest_err.NewInternalServerError("Error committing transaction")
	}

	logger.Info("Successful DeleteAppointmentByIdAndUserID repository", zap.Int("id", id))
	return nil
}
func (ar *appointmentRepository) DeleteAppointmentByIdAndBarberID(ctx context.Context, id int, barberId int) *rest_err.RestErr {
	logger.Info("Init DeleteAppointmentByIdAndUserID repository", zap.String("journey", "DeleteAppointmentByIdAndUserID"))

	tx, err := ar.databaseConection.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		logger.Error("Error starting transaction", err)
		return rest_err.NewInternalServerError("Error starting transaction")
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback(ctx)
			panic(p)
		} else if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	query := `DELETE FROM appointments WHERE id = $1 AND barber_id = $2;`
	result, err := tx.Exec(ctx, query, id, barberId)
	if err != nil {
		logger.Error("Error executing delete", err)
		return rest_err.NewInternalServerError("Database error during appointment_domain deletion")
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected > 1 {
		logger.Error("Multiple rows affected during delete", pgx.ErrTooManyRows, zap.Int64("rowsAffected", rowsAffected))
		_ = tx.Rollback(ctx)
		return rest_err.NewInternalServerError("Multiple rows affected. Rolling back transaction.")
	} else if rowsAffected == 0 {
		logger.WarnWithoutError("No appointment_domain found with given ID", zap.Int("id", id))
		return rest_err.NewNotFoundError("Appointment not found")
	}

	if err := tx.Commit(ctx); err != nil {
		logger.Error("Error committing transaction", err)
		return rest_err.NewInternalServerError("Error committing transaction")
	}

	logger.Info("Successful DeleteAppointmentByIdAndUserID repository", zap.Int("id", id))
	return nil
}
